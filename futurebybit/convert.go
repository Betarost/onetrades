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

func ConvertPositions(answ []Position) (res []entity.Position) {
	for _, item := range answ {
		positionSide := "LONG"
		if item.PositionIdx == 2 {
			positionSide = "SHORT"
		} else if item.PositionIdx == 0 {
			if item.Side == "Sell" {
				positionSide = "SHORT"
			}
		}

		res = append(res, entity.Position{
			Symbol:           item.Symbol,
			PositionSide:     positionSide,
			PositionAmt:      utils.StringToFloat(item.Size),
			EntryPrice:       utils.StringToFloat(item.AvgPrice),
			MarkPrice:        utils.StringToFloat(item.MarkPrice),
			UnRealizedProfit: utils.StringToFloat(item.UnrealisedPnl),
			Notional:         utils.StringToFloat(item.PositionValue),
			UpdateTime:       utils.StringToInt64(item.UpdatedTime),
		})
	}
	return res
}
