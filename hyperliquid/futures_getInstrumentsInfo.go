package hyperliquid

import (
	"context"
	"math"
	"strconv"

	"github.com/Betarost/onetrades/entity"
)

type futures_getInstrumentsInfo struct {
	client *FuturesClient
	symbol string
}

// Symbol устанавливает символ для фильтрации инструментов
func (r *futures_getInstrumentsInfo) Symbol(symbol string) *futures_getInstrumentsInfo {
	r.symbol = symbol
	return r
}

// Do выполняет запрос получения информации об инструментах
func (r *futures_getInstrumentsInfo) Do(_ context.Context) ([]entity.Futures_InstrumentsInfo, error) {
	// Получаем метаданные о всех инструментах
	meta, err := r.client.info.Meta()
	if err != nil {
		return nil, err
	}

	if meta == nil || meta.Universe == nil {
		return []entity.Futures_InstrumentsInfo{}, nil
	}

	instruments := make([]entity.Futures_InstrumentsInfo, 0, len(meta.Universe))

	// Конвертируем каждый инструмент
	for _, assetMeta := range meta.Universe {
		if assetMeta.Name == "" {
			continue
		}

		// Фильтруем по символу если указан
		if r.symbol != "" && assetMeta.Name != r.symbol {
			continue
		}

		instrument := convertInstrumentInfo(&assetMeta)

		if assetMeta.SzDecimals > 0 {
			minQtyFloat := 1.0 / math.Pow10(assetMeta.SzDecimals)
			instrument.MinQty = strconv.FormatFloat(minQtyFloat, 'f', assetMeta.SzDecimals, 64)
		} else {
			instrument.MinQty = "1"
		}

		instruments = append(instruments, *instrument)
	}

	return instruments, nil
}
