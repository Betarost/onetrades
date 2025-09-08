package hyperliquid

import (
	"context"
	"strconv"
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
func (r *futures_ordersHistory) Do(_ context.Context) ([]entity.Futures_OrdersHistory, error) {
	// Устанавливаем временные рамки
	endTime := time.Now().UnixMilli()
	if r.endTime != nil {
		endTime = *r.endTime
	}

	startTime := endTime - (24 * 60 * 60 * 1000) // По умолчанию последние 24 часа
	if r.startTime != nil {
		startTime = *r.startTime
	}

	// Получаем историю заполнений (fills) как приближение к истории ордеров
	fills, err := r.client.info.UserFillsByTime(r.client.AccountAddress(), startTime, &endTime)
	if err != nil {
		return nil, err
	}

	orders := make([]entity.Futures_OrdersHistory, 0, len(fills))

	// Конвертируем заполнения в историю ордеров
	for _, fill := range fills {
		// Фильтруем по символу если указан
		if r.symbol != "" && fill.Coin != r.symbol {
			continue
		}

		// Применяем лимит если указан
		if r.limit != nil && len(orders) >= *r.limit {
			break
		}

		// Определяем сторону
		var side string
		if fill.Side == HyperliquidSideBuy {
			side = string(entity.SideTypeBuy)
		} else {
			side = string(entity.SideTypeSell)
		}

		// Создаем запись истории ордера на основе заполнения
		order := entity.Futures_OrdersHistory{
			Symbol:         fill.Coin,
			OrderID:        strconv.FormatInt(fill.Oid, 10),
			Side:           side,
			PositionSize:   fill.Size,
			ExecutedSize:   fill.Size, // Заполнение означает полное исполнение
			Price:          fill.Price,
			ExecutedPrice:  fill.Price,
			RealisedProfit: fill.ClosedPnl,
			Fee:            fill.Fee,
			Type:           "limit",  // Предполагаем лимитный ордер
			Status:         "FILLED", // Заполнение означает исполненный ордер
			CreateTime:     fill.Time,
			UpdateTime:     fill.Time,
		}

		orders = append(orders, order)
	}

	return orders, nil
}
