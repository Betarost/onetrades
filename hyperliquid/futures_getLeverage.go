package hyperliquid

import (
	"context"
	"strconv"

	"github.com/Betarost/onetrades/entity"
)

type futures_getLeverage struct {
	client *FuturesClient
	symbol string
}

// Symbol устанавливает символ
func (r *futures_getLeverage) Symbol(symbol string) *futures_getLeverage {
	r.symbol = symbol
	return r
}

// Do выполняет запрос получения леверража
func (r *futures_getLeverage) Do(_ context.Context) (*entity.Futures_Leverage, error) {
	// Используем UserActiveAssetData для получения текущего леверража для символа
	assetData, err := r.client.info.UserActiveAssetData(r.client.AccountAddress(), r.symbol)
	if err != nil {
		return nil, err
	}

	// Получаем леверадж из ответа
	leverage := strconv.Itoa(assetData.Leverage.Value)

	return &entity.Futures_Leverage{
		Symbol:   r.symbol,
		Leverage: leverage,
	}, nil
}
