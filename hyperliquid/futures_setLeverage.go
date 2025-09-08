package hyperliquid

import (
	"context"
	"strconv"

	"github.com/Betarost/onetrades/entity"
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

// Do выполняет запрос установки леверража
func (r *futures_setLeverage) Do(_ context.Context) (*entity.Futures_Leverage, error) {
	// Конвертируем строку в число
	leverageInt, err := strconv.Atoi(r.leverage)
	if err != nil {
		return nil, err
	}

	// Выполняем запрос (используем isolated margin по умолчанию)
	_, err = r.client.exchange.UpdateLeverage(leverageInt, r.symbol, false)
	if err != nil {
		return nil, err
	}

	return &entity.Futures_Leverage{
		Symbol:   r.symbol,
		Leverage: r.leverage,
	}, nil
}
