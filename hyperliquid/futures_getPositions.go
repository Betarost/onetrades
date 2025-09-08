package hyperliquid

import (
	"context"

	"github.com/Betarost/onetrades/entity"
)

type futures_getPositions struct {
	client *FuturesClient
	symbol string
}

// Symbol устанавливает символ для фильтрации позиций
func (r *futures_getPositions) Symbol(symbol string) *futures_getPositions {
	r.symbol = symbol
	return r
}

// Do выполняет запрос получения позиций
func (r *futures_getPositions) Do(ctx context.Context) ([]*entity.Position, error) {
	// Получаем состояние пользователя
	userState, err := r.client.info.UserState(ctx, r.client.exchange.AccountAddress())
	if err != nil {
		return nil, err
	}

	positions := make([]*entity.Position, 0, len(userState.AssetPositions))

	// Проходим по всем позициям
	for _, assetPosition := range userState.AssetPositions {
		if assetPosition.Position == nil {
			continue
		}

		// Фильтруем по символу если указан
		if r.symbol != "" && assetPosition.Position.Coin != r.symbol {
			continue
		}

		// Конвертируем позицию
		position := convertPosition(assetPosition.Position)
		if position != nil {
			positions = append(positions, position)
		}
	}

	return positions, nil
}
