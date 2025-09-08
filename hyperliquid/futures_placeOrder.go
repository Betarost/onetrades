package hyperliquid

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Betarost/onetrades/entity"
	"github.com/sonirico/go-hyperliquid"
)

type futures_placeOrder struct {
	client        *FuturesClient
	symbol        string
	side          entity.SideType
	orderType     entity.OrderType
	timeInForce   entity.TimeInForceType
	quantity      string
	price         string
	stopPrice     string
	clientOrderId string
	reduceOnly    bool
}

// Symbol устанавливает торговый символ
func (r *futures_placeOrder) Symbol(symbol string) *futures_placeOrder {
	r.symbol = symbol
	return r
}

// Side устанавливает сторону ордера (покупка/продажа)
func (r *futures_placeOrder) Side(side entity.SideType) *futures_placeOrder {
	r.side = side
	return r
}

// Type устанавливает тип ордера
func (r *futures_placeOrder) Type(orderType entity.OrderType) *futures_placeOrder {
	r.orderType = orderType
	return r
}

// TimeInForce устанавливает время действия ордера
func (r *futures_placeOrder) TimeInForce(timeInForce entity.TimeInForceType) *futures_placeOrder {
	r.timeInForce = timeInForce
	return r
}

// Quantity устанавливает количество
func (r *futures_placeOrder) Quantity(quantity string) *futures_placeOrder {
	r.quantity = quantity
	return r
}

// Price устанавливает цену для лимитных ордеров
func (r *futures_placeOrder) Price(price string) *futures_placeOrder {
	r.price = price
	return r
}

// StopPrice устанавливает стоп-цену для стоп-ордеров
func (r *futures_placeOrder) StopPrice(stopPrice string) *futures_placeOrder {
	r.stopPrice = stopPrice
	return r
}

// NewClientOrderId устанавливает клиентский ID ордера
func (r *futures_placeOrder) NewClientOrderId(clientOrderId string) *futures_placeOrder {
	r.clientOrderId = clientOrderId
	return r
}

// ReduceOnly устанавливает флаг только уменьшения позиции
func (r *futures_placeOrder) ReduceOnly(reduceOnly bool) *futures_placeOrder {
	r.reduceOnly = reduceOnly
	return r
}

// Do выполняет размещение ордера
func (r *futures_placeOrder) Do(ctx context.Context) (*entity.Order, error) {
	if r.symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}
	if r.quantity == "" {
		return nil, fmt.Errorf("quantity is required")
	}

	// Конвертируем количество в float для проверки и обработки
	qty, err := strconv.ParseFloat(r.quantity, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid quantity: %v", err)
	}

	// Определяем направление для Hyperliquid
	var isBuy bool
	if r.side == entity.SideTypeBuy {
		isBuy = true
	} else {
		isBuy = false
		// Для продажи делаем количество отрицательным
		qty = -qty
	}

	// Создаем параметры ордера
	orderRequest := &hyperliquid.OrderRequest{
		Coin:       r.symbol,
		IsBuy:      isBuy,
		Sz:         qty,
		ReduceOnly: r.reduceOnly,
	}

	// Устанавливаем цену в зависимости от типа ордера
	switch r.orderType {
	case entity.OrderTypeMarket:
		// Для рыночного ордера используем null цену
		orderRequest.LimitPx = nil
		orderRequest.OrderType = &hyperliquid.OrderTypeMarket{}
	case entity.OrderTypeLimit:
		if r.price == "" {
			return nil, fmt.Errorf("price is required for limit orders")
		}
		price, err := strconv.ParseFloat(r.price, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid price: %v", err)
		}
		orderRequest.LimitPx = &price
		orderRequest.OrderType = &hyperliquid.OrderTypeLimit{}
	default:
		return nil, fmt.Errorf("unsupported order type: %v", r.orderType)
	}

	// Устанавливаем клиентский ID если указан
	if r.clientOrderId != "" {
		orderRequest.Cloid = &r.clientOrderId
	}

	// Размещаем ордер
	response, err := r.client.exchange.Order(ctx, orderRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to place order: %v", err)
	}

	// Проверяем успешность операции
	if response.Type != "order" {
		return nil, fmt.Errorf("unexpected response type: %s", response.Type)
	}

	if len(response.Data.Statuses) == 0 {
		return nil, fmt.Errorf("no order status returned")
	}

	status := response.Data.Statuses[0]
	if status.Error != "" {
		return nil, fmt.Errorf("order failed: %s", status.Error)
	}

	// Создаем результат
	order := &entity.Order{
		Symbol:        r.symbol,
		OrderId:       strconv.FormatUint(uint64(status.Resting.Oid), 10),
		ClientOrderId: r.clientOrderId,
		Price:         r.price,
		OrigQty:       r.quantity,
		ExecutedQty:   "0",
		Status:        entity.OrderStatusTypeNew,
		TimeInForce:   r.timeInForce,
		Type:          r.orderType,
		Side:          r.side,
		Time:          status.Resting.Timestamp,
		UpdateTime:    status.Resting.Timestamp,
	}

	return order, nil
}
