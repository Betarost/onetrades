package futurebingx

import (
	"github.com/Betarost/onetrades/entity"
	"github.com/Betarost/onetrades/utils"
)

func ConvertAccountBalance(answ []Balance) (res []entity.AccountBalance) {
	for _, item := range answ {
		res = append(res, entity.AccountBalance{
			Asset:            item.Asset,
			Balance:          utils.StringToFloat(item.Balance),
			AvailableBalance: utils.StringToFloat(item.AvailableMargin),
			UnrealizedProfit: utils.StringToFloat(item.UnrealizedProfit),
		})
	}
	return res
}
