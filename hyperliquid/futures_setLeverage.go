package hyperliquid

import (
	"context"
	"strconv"

	"github.com/Betarost/onetrades/entity"
	"github.com/sonirico/go-hyperliquid"
)

type futures_setLeverage struct {
	client   *FuturesClient
	symbol   string
	leverage string
}

// Symbol устанавливает символ
func (r *futures_setLeverage) Symbol(symbol string) *futures_setLeverage {
	r.symbol = symbol
	return r
}

// Leverage устанавливает леверадж
func (r *futures_setLeverage) Leverage(leverage string) *futures_setLeverage {
	r.leverage = leverage
	return r
}

// Do выполняет установку леверража
func (r *futures_setLeverage) Do(ctx context.Context) (*entity.LeverageInfo, error) {
	// Конвертируем строку в int для Hyperliquid API
	leverageInt, err := strconv.Atoi(r.leverage)
	if err != nil {
		return nil, err
	}

	// Создаем запрос на изменение леверража
	leverageRequest := &hyperliquid.UpdateLeverageRequest{
		Coin:     r.symbol,
		IsCross:  false, // Hyperliquid использует изолированную маржу
		Leverage: leverageInt,
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

	return &entity.LeverageInfo{
		Symbol:   r.symbol,
		Leverage: r.leverage,
	}, nil
}
