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
func (r *futures_getMarginMode) Do(ctx context.Context) (*entity.MarginMode, error) {
	// Используем UserActiveAssetData для получения текущего режима маржи для символа
	assetData, err := r.client.info.UserActiveAssetData(ctx, r.client.exchange.AccountAddress(), r.symbol)
	if err != nil {
		return nil, err
	}

	// Определяем тип маржи из ответа
	var marginType entity.MarginType
	if assetData.Leverage.Type == "cross" {
		marginType = entity.MarginTypeCrossed
	} else {
		marginType = entity.MarginTypeIsolated
	}

	return &entity.MarginMode{
		Symbol:     r.symbol,
		MarginType: marginType,
	}, nil
}
