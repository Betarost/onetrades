package futuremexc

import (
	"github.com/Betarost/onetrades/entity"
)

func ConvertAccountBalance(answ []Balance) (res []entity.AccountBalance) {
	for _, item := range answ {
		if item.Unrealized < 0 {
			item.Equity -= item.Unrealized
		}
		res = append(res, entity.AccountBalance{
			Asset:            item.Currency,
			Balance:          item.Equity,
			AvailableBalance: item.AvailableBalance,
			UnrealizedProfit: item.Unrealized,
		})
	}
	return res
}
