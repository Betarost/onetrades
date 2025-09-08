package hyperliquid

import (
	"context"

	"github.com/Betarost/onetrades/entity"
)

type futures_getMarginMode struct {
	client *FuturesClient
	symbol string
}

// Symbol устанавливает символ
func (r *futures_getMarginMode) Symbol(symbol string) *futures_getMarginMode {
	r.symbol = symbol
	return r
}

// Do выполняет запрос получения режима маржи
func (r *futures_getMarginMode) Do(_ context.Context) (*entity.Futures_MarginMode, error) {
	// Используем UserActiveAssetData для получения текущего режима маржи для символа
	assetData, err := r.client.info.UserActiveAssetData(r.client.AccountAddress(), r.symbol)
	if err != nil {
		return nil, err
	}

	// Определяем тип маржи из ответа
	var marginModeStr string
	if assetData.Leverage.Type == "cross" {
		marginModeStr = "cross"
	} else {
		marginModeStr = "isolated"
	}

	return &entity.Futures_MarginMode{
		MarginMode: marginModeStr,
	}, nil
}
