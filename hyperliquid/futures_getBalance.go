package hyperliquid

import (
	"context"

	"github.com/Betarost/onetrades/entity"
)

type futures_getBalance struct {
	client *FuturesClient
}

// Do выполняет запрос получения баланса
func (r *futures_getBalance) Do(_ context.Context) ([]entity.FuturesBalance, error) {
	// Для фьючерсов получаем маржинальную информацию
	userState, err := r.client.info.UserState(r.client.AccountAddress())
	if err != nil {
		return nil, err
	}

	// Конвертируем маржинальные балансы
	balances := convertFuturesBalances(userState)
	return balances, nil
}
