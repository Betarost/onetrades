package futurebitget

import (
	"github.com/Betarost/onetrades/entity"
	"github.com/Betarost/onetrades/utils"
)

func ConvertAccountBalance(answ []Balance) (res []entity.AccountBalance) {
	for _, item := range answ {
		balance := utils.StringToFloat(item.AccountEquity)
		if utils.StringToFloat(item.UnrealizedPL) < 0 {
			balance -= utils.StringToFloat(item.UnrealizedPL)
		}
		res = append(res, entity.AccountBalance{
			Asset:            item.MarginCoin,
			Balance:          balance,
			AvailableBalance: utils.StringToFloat(item.Available),
			UnrealizedProfit: utils.StringToFloat(item.UnrealizedPL),
		})
	}
	return res
}
