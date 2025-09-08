package hyperliquid

import (
	"context"

	"github.com/Betarost/onetrades/entity"
)

type futures_getPositionMode struct {
	client *FuturesClient
	symbol string
}

// Symbol устанавливает символ
func (r *futures_getPositionMode) Symbol(symbol string) *futures_getPositionMode {
	r.symbol = symbol
	return r
}

// Do выполняет запрос получения режима позиции
func (r *futures_getPositionMode) Do(_ context.Context) (*entity.Futures_PositionsMode, error) {
	// Hyperliquid поддерживает только односторонний режим позиций
	return &entity.Futures_PositionsMode{
		HedgeMode: false, // Hyperliquid не поддерживает hedge mode
	}, nil
}
