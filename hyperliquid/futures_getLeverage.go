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
func (r *futures_getLeverage) Do(ctx context.Context) (*entity.LeverageInfo, error) {
	// Используем UserActiveAssetData для получения текущего леверража для символа
	assetData, err := r.client.info.UserActiveAssetData(ctx, r.client.exchange.AccountAddress(), r.symbol)
	if err != nil {
		return nil, err
	}

	// Получаем леверадж из ответа
	leverage := strconv.Itoa(assetData.Leverage.Value)

	return &entity.LeverageInfo{
		Symbol:   r.symbol,
		Leverage: leverage,
	}, nil
}
