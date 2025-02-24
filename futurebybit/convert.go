package futurebybit

import (
	"github.com/Betarost/onetrades/entity"
	"github.com/Betarost/onetrades/utils"
)

func ConvertAccountBalance(answ []Balance) (res []entity.AccountBalance) {
	for _, item := range answ {
		res = append(res, entity.AccountBalance{
			Asset:            "USDT",
			Balance:          utils.StringToFloat(item.TotalWalletBalance),
			AvailableBalance: utils.StringToFloat(item.TotalAvailableBalance),
			UnrealizedProfit: utils.StringToFloat(item.TotalPerpUPL),
		})
	}
	return res
}
