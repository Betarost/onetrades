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

func ConvertPositions(answ []Position) (res []entity.Position) {
	for _, item := range answ {
		positionSide := "LONG"
		if item.PositionType == 2 {
			positionSide = "SHORT"
		}

		res = append(res, entity.Position{
			Symbol:        item.Symbol,
			PositionSide:  positionSide,
			PositionAmt:   item.HoldVol,
			EntryPrice:    item.HoldAvgPrice,
			InitialMargin: item.Oim,
			// UnRealizedProfit: utils.StringToFloat(item.UnrealisedPnl),
			Notional:   item.HoldVol * item.HoldAvgPrice,
			UpdateTime: item.UpdateTime,
		})
	}
	return res
}
