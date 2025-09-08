package hyperliquid

import (
	"context"
	"fmt"
	"strconv"

	"github.com/sonirico/go-hyperliquid"

	"github.com/Betarost/onetrades/entity"
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
func (r *futures_cancelOrder) Do(_ context.Context) ([]entity.PlaceOrder, error) {
	if r.symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}

	var response *hyperliquid.APIResponse[hyperliquid.CancelOrderResponse]
	var err error

	if r.orderId != "" {
		// Отмена по ID ордера
		orderIdInt, parseErr := strconv.ParseInt(r.orderId, 10, 64)
		if parseErr != nil {
			return nil, fmt.Errorf("invalid order ID: %v", parseErr)
		}
		response, err = r.client.exchange.Cancel(r.symbol, orderIdInt)
	} else if r.clientOrderId != "" {
		// Отмена по клиентскому ID
		response, err = r.client.exchange.CancelByCloid(r.symbol, r.clientOrderId)
	} else {
		return nil, fmt.Errorf("either orderId or clientOrderId is required")
	}

	if err != nil {
		return nil, err
	}

	// Проверяем ответ
	if !response.Ok {
		return nil, fmt.Errorf("cancel failed: %s", response.Err)
	}

	// Возвращаем результат отмены
	result := []entity.PlaceOrder{
		{
			OrderID:       r.orderId,
			ClientOrderID: r.clientOrderId,
		},
	}

	return result, nil
}
