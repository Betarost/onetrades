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

func ConvertContractsInfo(answ []ContractsInfo) (res []entity.ContractInfo) {
	if len(answ) == 0 {
		return res
	}
	for _, item := range answ {
		res = append(res, entity.ContractInfo{
			Symbol:      item.InstId,
			MaxLeverage: utils.StringToInt(item.Lever),
		})
	}
	return res
}

func ConvertOrderList(answ []OrderList) (res []entity.OrderList) {
	for _, item := range answ {
		positionSide := "LONG"
		if item.PosSide == "net" {
			if strings.ToUpper(item.Side) == "SELL" {
				positionSide = "SHORT"
			}
		} else {
			positionSide = strings.ToUpper(item.PosSide)
		}

		res = append(res, entity.OrderList{
			Symbol:       item.InstId,
			OrderID:      item.OrdId,
			PositionSide: positionSide,
			Side:         item.Side,
			PositionAmt:  utils.StringToFloat(item.Sz),
			Price:        utils.StringToFloat(item.Px),
			Notional:     utils.StringToFloat(item.Sz) * utils.StringToFloat(item.Px),
			Type:         strings.ToUpper(item.OrdType),
			Status:       item.State,
			UpdateTime:   utils.StringToInt64(item.UTime),
		})
	}
	return res
}
