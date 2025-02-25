package futurebinance

import (
	"github.com/Betarost/onetrades/entity"
	"github.com/Betarost/onetrades/utils"
)

func ConvertAccountBalance(answ []Balance) (res []entity.AccountBalance) {
	for _, item := range answ {
		res = append(res, entity.AccountBalance{
			Asset:              item.Asset,
			Balance:            utils.StringToFloat(item.Balance),
			AvailableBalance:   utils.StringToFloat(item.AvailableBalance),
			CrossWalletBalance: utils.StringToFloat(item.CrossWalletBalance),
			UnrealizedProfit:   utils.StringToFloat(item.CrossUnPnl),
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
			EntryPrice:       utils.StringToFloat(item.EntryPrice),
			MarkPrice:        utils.StringToFloat(item.MarkPrice),
			UnRealizedProfit: utils.StringToFloat(item.UnRealizedProfit),
			Notional:         utils.StringToFloat(item.Notional),
			InitialMargin:    utils.StringToFloat(item.InitialMargin),
			UpdateTime:       item.UpdateTime,
		})
	}
	return res
}
