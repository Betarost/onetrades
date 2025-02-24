package futuregate

import (
	"github.com/Betarost/onetrades/entity"
	"github.com/Betarost/onetrades/utils"
)

func ConvertAccountBalance(answ Balance) (res []entity.AccountBalance) {

	res = append(res, entity.AccountBalance{
		Asset:            answ.Currency,
		Balance:          utils.StringToFloat(answ.Total),
		AvailableBalance: utils.StringToFloat(answ.Available),
		UnrealizedProfit: utils.StringToFloat(answ.Unrealised_pnl),
	})
	return res
}
