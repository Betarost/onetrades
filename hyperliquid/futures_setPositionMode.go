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

// Do выполняет установку режима позиций
func (r *futures_setPositionMode) Do(ctx context.Context) (*entity.PositionMode, error) {
	// Hyperliquid поддерживает только односторонний режим позиций
	if r.dualSidePosition {
		return nil, fmt.Errorf("Hyperliquid supports only one-way position mode")
	}

	return &entity.PositionMode{
		DualSidePosition: false,
	}, nil
}
