package hyperliquid

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Betarost/onetrades/entity"
	"github.com/sonirico/go-hyperliquid"
)

type futures_cancelOrder struct {
	client        *FuturesClient
	symbol        string
	orderId       string
	clientOrderId string
}

// Symbol устанавливает торговый символ
func (r *futures_cancelOrder) Symbol(symbol string) *futures_cancelOrder {
	r.symbol = symbol
	return r
}

// OrderId устанавливает ID ордера для отмены
func (r *futures_cancelOrder) OrderId(orderId string) *futures_cancelOrder {
	r.orderId = orderId
	return r
}

// OrigClientOrderId устанавливает клиентский ID ордера для отмены
func (r *futures_cancelOrder) OrigClientOrderId(clientOrderId string) *futures_cancelOrder {
	r.clientOrderId = clientOrderId
	return r
}

// Do выполняет отмену ордера
func (r *futures_cancelOrder) Do(ctx context.Context) (*entity.Order, error) {
	if r.symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}

	var cancelRequest *hyperliquid.CancelRequest

	if r.clientOrderId != "" {
		// Отмена по клиентскому ID
		cancelRequest = &hyperliquid.CancelRequest{
			Coin:  r.symbol,
			Cloid: &r.clientOrderId,
		}
	} else if r.orderId != "" {
		// Отмена по ID ордера
		orderIdInt, err := strconv.ParseUint(r.orderId, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid order ID: %v", err)
		}
		orderIdInt32 := int32(orderIdInt)
		cancelRequest = &hyperliquid.CancelRequest{
			Coin: r.symbol,
			Oid:  &orderIdInt32,
		}
	} else {
		return nil, fmt.Errorf("either orderId or clientOrderId is required")
	}

	// Выполняем отмену
	response, err := r.client.exchange.Cancel(ctx, cancelRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to cancel order: %v", err)
	}

	// Проверяем результат
	if response.Type != "cancel" {
		return nil, fmt.Errorf("unexpected response type: %s", response.Type)
	}

	if len(response.Data.Statuses) == 0 {
		return nil, fmt.Errorf("no cancel status returned")
	}

	status := response.Data.Statuses[0]
	if status.Error != "" {
		return nil, fmt.Errorf("cancel failed: %s", status.Error)
	}

	// Создаем результат
	order := &entity.Order{
		Symbol:        r.symbol,
		OrderId:       r.orderId,
		ClientOrderId: r.clientOrderId,
		Status:        entity.OrderStatusTypeCanceled,
		Side:          entity.SideTypeBuy, // Значение по умолчанию, реальное значение нужно получать отдельно
		Type:          entity.OrderTypeLimit,
		TimeInForce:   entity.TimeInForceTypeGTC,
	}

	return order, nil
}
