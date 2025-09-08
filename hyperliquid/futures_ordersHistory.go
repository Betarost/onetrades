package hyperliquid

import (
	"context"
	"time"

	"github.com/Betarost/onetrades/entity"
)

type futures_ordersHistory struct {
	client    *FuturesClient
	symbol    string
	startTime *int64
	endTime   *int64
	limit     *int
}

// Symbol устанавливает символ для фильтрации
func (r *futures_ordersHistory) Symbol(symbol string) *futures_ordersHistory {
	r.symbol = symbol
	return r
}

// StartTime устанавливает начальное время
func (r *futures_ordersHistory) StartTime(startTime int64) *futures_ordersHistory {
	r.startTime = &startTime
	return r
}

// EndTime устанавливает конечное время
func (r *futures_ordersHistory) EndTime(endTime int64) *futures_ordersHistory {
	r.endTime = &endTime
	return r
}

// Limit устанавливает лимит записей
func (r *futures_ordersHistory) Limit(limit int) *futures_ordersHistory {
	r.limit = &limit
	return r
}

// Do выполняет запрос получения истории ордеров
func (r *futures_ordersHistory) Do(ctx context.Context) ([]*entity.Order, error) {
	// Определяем временные рамки
	endTime := time.Now().UnixMilli()
	if r.endTime != nil {
		endTime = *r.endTime
	}

	startTime := endTime - (24 * 60 * 60 * 1000) // По умолчанию последние 24 часа
	if r.startTime != nil {
		startTime = *r.startTime
	}

	// Получаем историю заполнений (fills) как приближение к истории ордеров
	fills, err := r.client.info.UserFills(ctx, r.client.exchange.AccountAddress(), startTime, endTime)
	if err != nil {
		return nil, err
	}

	orders := make([]*entity.Order, 0, len(fills))

	// Конвертируем fills в ордера (приблизительно)
	for _, fill := range fills {
		// Фильтруем по символу если указан
		if r.symbol != "" && fill.Coin != r.symbol {
			continue
		}

		// Применяем лимит если указан
		if r.limit != nil && len(orders) >= *r.limit {
			break
		}

		// Конвертируем fill в order
		order := &entity.Order{
			Symbol:              fill.Coin,
			OrderId:             fill.Oid,
			Price:               fill.Px,
			OrigQty:             fill.Sz,
			ExecutedQty:         fill.Sz,
			CummulativeQuoteQty: "0", // Не предоставляется
			Status:              entity.OrderStatusTypeFilled,
			TimeInForce:         entity.TimeInForceTypeGTC,
			Type:                entity.OrderTypeLimit,
			Side:                convertSide(fill.Side),
			Time:                fill.Time,
			UpdateTime:          fill.Time,
		}

		orders = append(orders, order)
	}

	return orders, nil
}
