package hyperliquid

import (
	"context"

	"github.com/Betarost/onetrades/entity"
	"github.com/sonirico/go-hyperliquid"
)

type futures_setMarginMode struct {
	client     *FuturesClient
	symbol     string
	marginType entity.MarginType
}

// Symbol устанавливает символ
func (r *futures_setMarginMode) Symbol(symbol string) *futures_setMarginMode {
	r.symbol = symbol
	return r
}

// MarginType устанавливает тип маржи
func (r *futures_setMarginMode) MarginType(marginType entity.MarginType) *futures_setMarginMode {
	r.marginType = marginType
	return r
}

// Do выполняет установку режима маржи
func (r *futures_setMarginMode) Do(ctx context.Context) (*entity.MarginMode, error) {
	// Определяем режим маржи для Hyperliquid
	var isCross bool
	switch r.marginType {
	case entity.MarginTypeCrossed:
		isCross = true
	case entity.MarginTypeIsolated:
		isCross = false
	default:
		isCross = false // По умолчанию Isolated
	}

	// Получаем текущий леверадж из UserActiveAssetData
	assetData, err := r.client.info.UserActiveAssetData(ctx, r.client.exchange.AccountAddress(), r.symbol)
	if err != nil {
		return nil, err
	}

	// Используем текущий леверадж
	currentLeverage := assetData.Leverage.Value

	// Создаем запрос на изменение режима маржи
	leverageRequest := &hyperliquid.UpdateLeverageRequest{
		Coin:     r.symbol,
		IsCross:  isCross,
		Leverage: currentLeverage, // Сохраняем текущий леверадж
	}

	// Выполняем запрос
	response, err := r.client.exchange.UpdateLeverage(ctx, leverageRequest)
	if err != nil {
		return nil, err
	}

	// Проверяем результат
	if response.Type != "leverage" {
		return nil, err
	}

	if len(response.Data.Statuses) > 0 && response.Data.Statuses[0].Error != "" {
		return nil, err
	}

	return &entity.MarginMode{
		Symbol:     r.symbol,
		MarginType: r.marginType,
	}, nil
}
