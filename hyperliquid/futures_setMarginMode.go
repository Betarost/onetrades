package hyperliquid

import (
	"context"

	"github.com/Betarost/onetrades/entity"
)

type futures_setMarginMode struct {
	client     *FuturesClient
	symbol     string
	marginType entity.MarginModeType
}

// Symbol устанавливает символ
func (r *futures_setMarginMode) Symbol(symbol string) *futures_setMarginMode {
	r.symbol = symbol
	return r
}

// MarginType устанавливает тип маржи
func (r *futures_setMarginMode) MarginType(marginType entity.MarginModeType) *futures_setMarginMode {
	r.marginType = marginType
	return r
}

// Do выполняет запрос установки режима маржи
func (r *futures_setMarginMode) Do(_ context.Context) (*entity.Futures_MarginMode, error) {
	// Определяем режим маржи для Hyperliquid
	var isCross bool
	switch r.marginType {
	case entity.MarginModeTypeCross:
		isCross = true
	case entity.MarginModeTypeIsolated:
		isCross = false
	default:
		isCross = false
	}

	// Получаем текущий леверадж из UserActiveAssetData
	assetData, err := r.client.info.UserActiveAssetData(r.client.AccountAddress(), r.symbol)
	if err != nil {
		return nil, err
	}

	// Используем текущий леверадж
	currentLeverage := assetData.Leverage.Value

	// Выполняем запрос на изменение режима маржи
	_, err = r.client.exchange.UpdateLeverage(currentLeverage, r.symbol, isCross)
	if err != nil {
		return nil, err
	}

	// Возвращаем результат
	var marginModeStr string
	if isCross {
		marginModeStr = "cross"
	} else {
		marginModeStr = "isolated"
	}

	return &entity.Futures_MarginMode{
		MarginMode: marginModeStr,
	}, nil
}
