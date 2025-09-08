package hyperliquid

import (
	"context"

	"github.com/Betarost/onetrades/entity"
)

type futures_getInstrumentsInfo struct {
	client *FuturesClient
}

// Do выполняет запрос получения информации об инструментах
func (r *futures_getInstrumentsInfo) Do(ctx context.Context) ([]*entity.Symbol, error) {
	// Получаем метаинформацию о всех инструментах
	meta, err := r.client.info.Meta(ctx)
	if err != nil {
		return nil, err
	}

	// Конвертируем в общий формат
	symbols := convertInstrumentInfo(meta)
	return symbols, nil
}
