package hyperliquid

import (
	"context"

	"github.com/Betarost/onetrades/entity"
)

type futures_getOrderList struct {
	client *FuturesClient
	symbol string
}

// Symbol устанавливает символ для фильтрации ордеров
func (r *futures_getOrderList) Symbol(symbol string) *futures_getOrderList {
	r.symbol = symbol
	return r
}

// Do выполняет запрос получения списка активных ордеров
func (r *futures_getOrderList) Do(_ context.Context) ([]entity.Futures_OrdersList, error) {
	// Получаем открытые ордера
	openOrders, err := r.client.info.OpenOrders(r.client.AccountAddress())
	if err != nil {
		return nil, err
	}

	orders := make([]entity.Futures_OrdersList, 0, len(openOrders))

	// Проходим по всем открытым ордерам
	for _, order := range openOrders {
		// Фильтруем по символу если указан
		if r.symbol != "" && order.Coin != r.symbol {
			continue
		}

		// Конвертируем ордер
		convertedOrder := convertOrder(&order)
		if convertedOrder != nil {
			orders = append(orders, *convertedOrder)
		}
	}

	return orders, nil
}
