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
func (r *futures_getPositions) Do(_ context.Context) ([]*entity.Futures_Positions, error) {
	// Получаем состояние пользователя
	userState, err := r.client.info.UserState(r.client.AccountAddress())
	if err != nil {
		return nil, err
	}

	positions := make([]*entity.Futures_Positions, 0, len(userState.AssetPositions))

	// Проходим по всем позициям
	for _, assetPosition := range userState.AssetPositions {
		// Проверяем что позиция не пустая (размер не равен "0")
		if assetPosition.Position.Szi == "" || assetPosition.Position.Szi == "0" {
			continue
		}

		// Фильтруем по символу если указан
		if r.symbol != "" && assetPosition.Position.Coin != r.symbol {
			continue
		}

		// Конвертируем позицию
		position := convertPosition(&assetPosition.Position)
		if position != nil {
			positions = append(positions, position)
		}
	}

	return positions, nil
}
