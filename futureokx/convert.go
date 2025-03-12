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

			MarginRatio:      item.MgnRatio,
			AutoDeleveraging: item.ADL,
		})
	}
	return res
}

func ConvertHistoryPositions(answ []HistoryPosition) (res []entity.HistoryPosition) {
	for _, item := range answ {

		status := entity.PositionStatusTypeCloseAll
		switch item.Type {
		case "1":
			status = entity.PositionStatusTypeClosePartially
		case "3":
			status = entity.PositionStatusTypeLiquidation
		case "4":
			status = entity.PositionStatusTypeLiquidationPartially
		case "5":
			status = entity.PositionStatusTypeAdl
		}

		res = append(res, entity.HistoryPosition{
			PositionID:       item.PosID,
			Symbol:           item.InstID,
			Status:           status, //entity
			PositionSide:     strings.ToUpper(item.Direction),
			AvgOpenPrice:     utils.StringToFloat(item.OpenAvgPx),
			AvgClosePrice:    utils.StringToFloat(item.CloseAvgPx),
			PositionOpenAmt:  utils.StringToFloat(item.OpenMaxPos),
			PositionCloseAmt: utils.StringToFloat(item.CloseTotalPos),
			RealizedProfit:   utils.StringToFloat(item.RealizedPnl),
			Pnl:              utils.StringToFloat(item.Pnl),
			PnlRatio:         utils.StringToFloat(item.PnlRatio),
			Fee:              utils.StringToFloat(item.Fee),
			FundingFee:       utils.StringToFloat(item.FundingFee),
			LiqPenalty:       utils.StringToFloat(item.LiqPenalty),
			CreateTime:       utils.StringToInt64(item.CTime),
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
			Symbol:       item.InstId,
			ContractSize: utils.StringToFloat(item.CtVal),
			MaxLeverage:  utils.StringToInt(item.Lever),
		})
	}
	return res
}

func ConvertKline(answ [][]string) (res []entity.Kline) {
	if len(answ) == 0 {
		return res
	}
	for _, item := range answ {
		if len(item) >= 6 {
			complete := true
			if item[5] == "0" {
				complete = false
			}
			res = append(res, entity.Kline{
				Time:         utils.StringToInt64(item[0]),
				OpenPrice:    utils.StringToFloat(item[1]),
				HighestPrice: utils.StringToFloat(item[2]),
				LowestPrice:  utils.StringToFloat(item[3]),
				ClosePrice:   utils.StringToFloat(item[4]),
				Complete:     complete,
			})
		}
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

func ConvertWsMarkPrice(answ PublicMarkPrice) (res entity.WsPublicMarkPriceEvent) {
	return entity.WsPublicMarkPriceEvent{
		Symbol: answ.InstId,
		Price:  utils.StringToFloat(answ.MarkPx),
		Time:   utils.StringToInt64(answ.Ts),
	}
}

func ConvertWsPrivateOrders(answ PrivateOrders) (res entity.WsPrivateOrdersEvent) {
	positionSide := "LONG"
	if answ.PosSide == "net" {
		if strings.ToUpper(answ.Side) == "SELL" {
			positionSide = "SHORT"
		}
	} else {
		positionSide = strings.ToUpper(answ.PosSide)
	}
	return entity.WsPrivateOrdersEvent{
		Symbol:       answ.InstId,
		OrderID:      answ.OrdId,
		PositionSide: positionSide,
		Side:         answ.Side,
		PositionAmt:  utils.StringToFloat(answ.Sz),
		Price:        utils.StringToFloat(answ.Px),
		Notional:     utils.StringToFloat(answ.NotionalUsd),
		Type:         strings.ToUpper(answ.OrdType),
		Status:       answ.State,
		Time:         utils.StringToInt64(answ.CTime),
		UpdateTime:   utils.StringToInt64(answ.UTime),
	}
}

func ConvertWsPrivatePositions(answ PrivatePositions) (res entity.WsPrivatePositionsEvent) {

	positionSide := "LONG"
	if answ.PosSide == "net" {
		if utils.StringToFloat(answ.Pos) < 0 {
			positionSide = "SHORT"
		}
	} else {
		positionSide = strings.ToUpper(answ.PosSide)
	}

	res = entity.WsPrivatePositionsEvent{
		Symbol:           answ.InstID,
		PositionSide:     positionSide,
		PositionAmt:      utils.StringToFloat(answ.Pos),
		EntryPrice:       utils.StringToFloat(answ.AvgPx),
		MarkPrice:        utils.StringToFloat(answ.MarkPx),
		InitialMargin:    utils.StringToFloat(answ.Imr),
		UnRealizedProfit: utils.StringToFloat(answ.Upl),
		RealizedProfit:   utils.StringToFloat(answ.RealizedPnl),
		Notional:         utils.StringToFloat(answ.NotionalUsd),
		UpdateTime:       utils.StringToInt64(answ.UTime),
	}
	return res
}
