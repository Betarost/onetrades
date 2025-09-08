package hyperliquid

import (
	"context"
	"strconv"
	"time"

	"github.com/Betarost/onetrades/entity"
)

type futures_positionsHistory struct {
	client    *FuturesClient
	symbol    string
	startTime *int64
	endTime   *int64
	limit     *int
}

// Symbol устанавливает символ для фильтрации
func (r *futures_positionsHistory) Symbol(symbol string) *futures_positionsHistory {
	r.symbol = symbol
	return r
}

// StartTime устанавливает начальное время
func (r *futures_positionsHistory) StartTime(startTime int64) *futures_positionsHistory {
	r.startTime = &startTime
	return r
}

// EndTime устанавливает конечное время
func (r *futures_positionsHistory) EndTime(endTime int64) *futures_positionsHistory {
	r.endTime = &endTime
	return r
}

// Limit устанавливает лимит записей
func (r *futures_positionsHistory) Limit(limit int) *futures_positionsHistory {
	r.limit = &limit
	return r
}

// Do выполняет запрос получения истории позиций
func (r *futures_positionsHistory) Do(_ context.Context) ([]*entity.Futures_Positions, error) {
	// Устанавливаем временные рамки
	endTime := time.Now().UnixMilli()
	if r.endTime != nil {
		endTime = *r.endTime
	}

	startTime := endTime - (7 * 24 * 60 * 60 * 1000) // По умолчанию 7 дней назад
	if r.startTime != nil {
		startTime = *r.startTime
	}

	// Получаем историю заполнений для восстановления истории позиций
	fills, err := r.client.info.UserFillsByTime(r.client.AccountAddress(), startTime, &endTime)
	if err != nil {
		return nil, err
	}

	// Восстанавливаем историю позиций из заполнений
	positionHistory := make(map[string]*entity.Futures_Positions)
	positions := make([]*entity.Futures_Positions, 0, len(fills))

	for _, fill := range fills {
		// Фильтруем по символу если указан
		if r.symbol != "" && fill.Coin != r.symbol {
			continue
		}

		// Ограничиваем количество если указан лимит
		if r.limit != nil && len(positions) >= *r.limit {
			break
		}

		// Рассчитываем финальную позицию после заполнения
		startPos, err1 := strconv.ParseFloat(fill.StartPosition, 64)
		fillSize, err2 := strconv.ParseFloat(fill.Size, 64)

		// Пропускаем записи с некорректными данными
		if err1 != nil || err2 != nil {
			continue
		}

		var finalPos float64
		if fill.Side == HyperliquidSideBuy { // Buy
			finalPos = startPos + fillSize
		} else { // Sell
			finalPos = startPos - fillSize
		}

		// Определяем сторону позиции
		var positionSide string
		var positionSize string
		if finalPos > 0 {
			positionSide = string(entity.PositionSideTypeLong)
			positionSize = strconv.FormatFloat(finalPos, 'f', -1, 64)
		} else if finalPos < 0 {
			positionSide = string(entity.PositionSideTypeShort)
			positionSize = strconv.FormatFloat(-finalPos, 'f', -1, 64) // Делаем размер положительным
		} else {
			positionSide = string(entity.PositionSideTypeBoth)
			positionSize = "0"
		}

		// Создаем позицию
		position := &entity.Futures_Positions{
			Symbol:           fill.Coin,
			PositionSide:     positionSide,
			PositionSize:     positionSize,
			EntryPrice:       fill.Price,
			MarkPrice:        fill.Price,
			UnRealizedProfit: fill.ClosedPnl,
			RealizedProfit:   fill.ClosedPnl,
			UpdateTime:       fill.Time,
		}

		// Используем уникальный ключ для избежания дублирования
		key := fill.Coin + "_" + strconv.FormatInt(fill.Time, 10)
		if _, exists := positionHistory[key]; !exists {
			positionHistory[key] = position
			positions = append(positions, position)
		}
	}

	return positions, nil
}
