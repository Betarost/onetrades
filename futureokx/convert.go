package futureokx

import (
	"strings"

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

func ConvertPositions(answ []Position) (res []entity.Position) {
	for _, item := range answ {
		positionSide := "LONG"
		if item.PosSide == "net" {
			if utils.StringToFloat(item.Pos) < 0 {
				positionSide = "SHORT"
			}
		} else {
			positionSide = strings.ToUpper(item.PosSide)
		}

		res = append(res, entity.Position{
			Symbol:           item.InstID,
			PositionSide:     positionSide,
			PositionAmt:      utils.StringToFloat(item.Pos),
			EntryPrice:       utils.StringToFloat(item.AvgPx),
			MarkPrice:        utils.StringToFloat(item.MarkPx),
			InitialMargin:    utils.StringToFloat(item.Imr),
			UnRealizedProfit: utils.StringToFloat(item.Upl),
			RealizedProfit:   utils.StringToFloat(item.RealizedPnl),
			Notional:         utils.StringToFloat(item.NotionalUsd),
			UpdateTime:       utils.StringToInt64(item.UTime),
		})
	}
	return res
}
