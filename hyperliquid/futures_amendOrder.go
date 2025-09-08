package hyperliquid

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Betarost/onetrades/entity"
	"github.com/sonirico/go-hyperliquid"
)

type futures_amendOrder struct {
	client        *FuturesClient
	symbol        string
	orderId       string
	clientOrderId string
	quantity      string
	price         string
}

// Symbol устанавливает торговый символ
func (r *futures_amendOrder) Symbol(symbol string) *futures_amendOrder {
	r.symbol = symbol
	return r
}

// OrderId устанавливает ID ордера для изменения
func (r *futures_amendOrder) OrderId(orderId string) *futures_amendOrder {
	r.orderId = orderId
	return r
}

// OrigClientOrderId устанавливает клиентский ID ордера для изменения
func (r *futures_amendOrder) OrigClientOrderId(clientOrderId string) *futures_amendOrder {
	r.clientOrderId = clientOrderId
	return r
}

// Quantity устанавливает новое количество
func (r *futures_amendOrder) Quantity(quantity string) *futures_amendOrder {
	r.quantity = quantity
	return r
}

// Price устанавливает новую цену
func (r *futures_amendOrder) Price(price string) *futures_amendOrder {
	r.price = price
	return r
}

// Do выполняет изменение ордера через отмену старого и создание нового
func (r *futures_amendOrder) Do(ctx context.Context) (*entity.Order, error) {
	if r.symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}

	// Получаем информацию о конкретном ордере
	var orderQueryResult *hyperliquid.OrderQueryResult
	var err error

	if r.orderId != "" {
		// Запрашиваем ордер по ID
		orderIdInt, parseErr := strconv.ParseInt(r.orderId, 10, 64)
		if parseErr != nil {
			return nil, fmt.Errorf("invalid order ID: %v", parseErr)
		}
		orderQueryResult, err = r.client.info.QueryOrderByOid(ctx, r.client.exchange.AccountAddress(), orderIdInt)
	} else if r.clientOrderId != "" {
		// Запрашиваем ордер по клиентскому ID
		orderQueryResult, err = r.client.info.QueryOrderByCloid(ctx, r.client.exchange.AccountAddress(), r.clientOrderId)
	} else {
		return nil, fmt.Errorf("either orderId or clientOrderId is required")
	}

	if err != nil {
		return nil, fmt.Errorf("failed to query order: %v", err)
	}

	// Проверяем, что ордер найден и активен
	if orderQueryResult.Status != "open" {
		return nil, fmt.Errorf("order is not open (status: %s)", orderQueryResult.Status)
	}

	originalOrder := &orderQueryResult.Order.Order

	// Шаг 1: Отменяем старый ордер
	cancelRequest := r.client.NewCancelOrder().Symbol(r.symbol)
	if r.orderId != "" {
		cancelRequest = cancelRequest.OrderId(r.orderId)
	} else {
		cancelRequest = cancelRequest.OrigClientOrderId(r.clientOrderId)
	}

	_, err = cancelRequest.Do(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to cancel original order: %v", err)
	}

	// Шаг 2: Создаем новый ордер с обновленными параметрами
	newOrderRequest := r.client.NewPlaceOrder().
		Symbol(r.symbol).
		Side(convertSide(string(originalOrder.Side))).
		Type(convertOrderType(originalOrder.OrderType)).
		TimeInForce(convertTimeInForce(string(originalOrder.Tif)))

	// Используем новые значения если указаны, иначе оригинальные
	if r.quantity != "" {
		newOrderRequest = newOrderRequest.Quantity(r.quantity)
	} else {
		newOrderRequest = newOrderRequest.Quantity(originalOrder.Sz)
	}

	if r.price != "" {
		newOrderRequest = newOrderRequest.Price(r.price)
	} else {
		newOrderRequest = newOrderRequest.Price(originalOrder.LimitPx)
	}

	// Создаем новый клиентский ID если был оригинальный
	if originalOrder.Cloid != nil && *originalOrder.Cloid != "" {
		newClientOrderId := *originalOrder.Cloid + "_amended"
		newOrderRequest = newOrderRequest.NewClientOrderId(newClientOrderId)
	}

	// Размещаем новый ордер
	newOrder, err := newOrderRequest.Do(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to place new order: %v", err)
	}

	return newOrder, nil
}
