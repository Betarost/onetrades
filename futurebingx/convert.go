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

func ConvertPositions(answ []Position) (res []entity.Position) {
	for _, item := range answ {
		res = append(res, entity.Position{
			Symbol:           item.Symbol,
			PositionSide:     item.PositionSide,
			PositionAmt:      utils.StringToFloat(item.PositionAmt),
			EntryPrice:       utils.StringToFloat(item.AvgPrice),
			MarkPrice:        utils.StringToFloat(item.MarkPrice),
			UnRealizedProfit: utils.StringToFloat(item.UnrealizedProfit),
			Notional:         utils.StringToFloat(item.PositionValue),
			UpdateTime:       item.UpdateTime,
		})
	}
	return res
}
