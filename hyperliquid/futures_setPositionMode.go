package hyperliquid

import (
	"context"
	"fmt"

	"github.com/Betarost/onetrades/entity"
)

type futures_setPositionMode struct {
	client           *FuturesClient
	dualSidePosition bool
}

// DualSidePosition устанавливает режим двусторонних позиций
func (r *futures_setPositionMode) DualSidePosition(dualSidePosition bool) *futures_setPositionMode {
	r.dualSidePosition = dualSidePosition
	return r
}

// Do выполняет запрос установки режима позиции
func (r *futures_setPositionMode) Do(_ context.Context) (*entity.Futures_PositionsMode, error) {
	// Hyperliquid поддерживает только односторонний режим позиций
	if r.dualSidePosition {
		return nil, fmt.Errorf("Hyperliquid does not support dual side position mode")
	}

	return &entity.Futures_PositionsMode{
		HedgeMode: false, // Hyperliquid не поддерживает hedge mode
	}, nil
}
