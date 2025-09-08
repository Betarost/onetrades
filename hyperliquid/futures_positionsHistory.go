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
func (r *futures_positionsHistory) Do(ctx context.Context) ([]*entity.Position, error) {
	// Определяем временные рамки
	endTime := time.Now().UnixMilli()
	if r.endTime != nil {
		endTime = *r.endTime
	}

	startTime := endTime - (7 * 24 * 60 * 60 * 1000) // По умолчанию последние 7 дней
	if r.startTime != nil {
		startTime = *r.startTime
	}

	// Получаем историю заполнений для восстановления истории позиций
	fills, err := r.client.info.UserFillsByTime(ctx, r.client.exchange.AccountAddress(), startTime, &endTime)
	if err != nil {
		return nil, err
	}

	// Группируем заполнения по символам и восстанавливаем историю позиций
	positionHistory := make(map[string]*entity.Position)
	positions := make([]*entity.Position, 0, len(fills))

	for _, fill := range fills {
		// Фильтруем по символу если указан
		if r.symbol != "" && fill.Coin != r.symbol {
			continue
		}

		// Применяем лимит если указан
		if r.limit != nil && len(positions) >= *r.limit {
			break
		}

		// Восстанавливаем позицию из заполнения
		startPos, _ := strconv.ParseFloat(fill.StartPosition, 64)
		fillSize, _ := strconv.ParseFloat(fill.Size, 64)

		// Определяем направление заполнения
		var finalPos float64
		if fill.Side == "B" { // Buy
			finalPos = startPos + fillSize
		} else { // Sell
			finalPos = startPos - fillSize
		}

		// Определяем сторону позиции
		var side entity.PositionSideType
		if finalPos > 0 {
			side = entity.PositionSideTypeLong
		} else if finalPos < 0 {
			side = entity.PositionSideTypeShort
			finalPos = -finalPos // Делаем размер положительным
		} else {
			side = entity.PositionSideTypeBoth
		}

		// Создаем запись истории позиции
		position := &entity.Position{
			Symbol:        fill.Coin,
			Size:          strconv.FormatFloat(finalPos, 'f', -1, 64),
			Side:          side,
			EntryPrice:    fill.Price,
			MarkPrice:     fill.Price, // Используем цену заполнения как приближение
			UnrealizedPnL: fill.ClosedPnl,
			UpdateTime:    fill.Time,
		}

		// Сохраняем позицию
		key := fill.Coin + "_" + strconv.FormatInt(fill.Time, 10)
		if _, exists := positionHistory[key]; !exists {
			positionHistory[key] = position
			positions = append(positions, position)
		}
	}

	return positions, nil
}
