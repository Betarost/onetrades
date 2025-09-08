package hyperliquid

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Betarost/onetrades/entity"
	hyperliquid "github.com/sonirico/go-hyperliquid"
)

type futures_placeOrder struct {
	client        *FuturesClient
	symbol        string
	side          entity.SideType
	quantity      string
	price         string
	orderType     entity.OrderType
	timeInForce   string
	clientOrderId string
}

// Symbol устанавливает символ
func (r *futures_placeOrder) Symbol(symbol string) *futures_placeOrder {
	r.symbol = symbol
	return r
}

// Side устанавливает сторону заказа
func (r *futures_placeOrder) Side(side entity.SideType) *futures_placeOrder {
	r.side = side
	return r
}

// Quantity устанавливает количество
func (r *futures_placeOrder) Quantity(quantity string) *futures_placeOrder {
	r.quantity = quantity
	return r
}

// Price устанавливает цену
func (r *futures_placeOrder) Price(price string) *futures_placeOrder {
	r.price = price
	return r
}

// Type устанавливает тип заказа
func (r *futures_placeOrder) Type(orderType entity.OrderType) *futures_placeOrder {
	r.orderType = orderType
	return r
}

// TimeInForce устанавливает время действия заказа
func (r *futures_placeOrder) TimeInForce(timeInForce string) *futures_placeOrder {
	r.timeInForce = timeInForce
	return r
}

// NewClientOrderId устанавливает клиентский ID заказа
func (r *futures_placeOrder) NewClientOrderId(clientOrderId string) *futures_placeOrder {
	r.clientOrderId = clientOrderId
	return r
}

// Do выполняет размещение заказа
func (r *futures_placeOrder) Do(_ context.Context) ([]entity.PlaceOrder, error) {
	// Валидация обязательных параметров
	if r.symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}
	if r.quantity == "" {
		return nil, fmt.Errorf("quantity is required")
	}

	// Конвертируем параметры
	isBuy := r.side == entity.SideTypeBuy

	// Парсим и валидируем количество
	size, err := strconv.ParseFloat(r.quantity, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid quantity: %v", err)
	}
	if size <= 0 {
		return nil, fmt.Errorf("quantity must be positive")
	}

	// Парсим и валидируем цену (для лимитных ордеров)
	var price float64
	if r.price != "" {
		price, err = strconv.ParseFloat(r.price, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid price: %v", err)
		}
		if price <= 0 {
			return nil, fmt.Errorf("price must be positive")
		}
	}

	// Определяем тип ордера
	var orderType hyperliquid.OrderType
	if r.orderType == entity.OrderTypeMarket {
		// Для рыночного ордера цена не нужна, используем 0
		if price == 0 {
			price = 1 // Минимальная цена для trigger
		}
		orderType = hyperliquid.OrderType{
			Trigger: &hyperliquid.TriggerOrderType{
				TriggerPx: price,
				IsMarket:  true,
				Tpsl:      hyperliquid.TakeProfit,
			},
		}
	} else {
		// Для лимитного ордера цена обязательна
		if price == 0 {
			return nil, fmt.Errorf("price is required for limit orders")
		}

		// Определяем время действия ордера
		tif := hyperliquid.TifGtc // По умолчанию GTC
		switch r.timeInForce {
		case "IOC":
			tif = hyperliquid.TifIoc
		case "FOK":
			// FOK нет в библиотеке, используем IOC
			tif = hyperliquid.TifIoc
		}

		orderType = hyperliquid.OrderType{
			Limit: &hyperliquid.LimitOrderType{
				Tif: tif,
			},
		}
	}

	// Создаем запрос
	var clientOrderID *string
	if r.clientOrderId != "" {
		clientOrderID = &r.clientOrderId
	}

	req := hyperliquid.CreateOrderRequest{
		Coin:          r.symbol,
		IsBuy:         isBuy,
		Price:         price,
		Size:          size,
		ReduceOnly:    false,
		OrderType:     orderType,
		ClientOrderID: clientOrderID,
	}

	// Выполняем запрос
	status, err := r.client.exchange.Order(req, nil)
	if err != nil {
		return nil, err
	}

	// Проверяем результат
	if status.Error != nil {
		return nil, fmt.Errorf("order failed: %s", *status.Error)
	}

	// Формируем результат
	var orderID string
	if status.Resting != nil {
		orderID = fmt.Sprintf("%d", status.Resting.Oid)
	}

	result := []entity.PlaceOrder{
		{
			OrderID:       orderID,
			ClientOrderID: r.clientOrderId,
		},
	}

	return result, nil
}
