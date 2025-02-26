package futureokx

import (
	"github.com/Betarost/onetrades/entity"
	"github.com/Betarost/onetrades/utils"
)

func ConvertAccountBalance(answ []Balance) (res []entity.AccountBalance) {
	if len(answ) == 0 {
		return res
	}
	for _, item := range answ[0].Details {
		res = append(res, entity.AccountBalance{
			Asset:            item.Ccy,
			Balance:          utils.StringToFloat(item.Eq),
			AvailableBalance: utils.StringToFloat(item.AvailBal),
			UnrealizedProfit: utils.StringToFloat(item.Upl),
		})
	}
	return res
}
