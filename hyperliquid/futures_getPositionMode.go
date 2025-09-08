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

// Do выполняет запрос получения режима позиций
func (r *futures_getPositionMode) Do(ctx context.Context) (*entity.PositionMode, error) {
	// Hyperliquid использует односторонний режим позиций
	return &entity.PositionMode{
		DualSidePosition: false,
	}, nil
}
