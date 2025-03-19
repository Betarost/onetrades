package futurebybit

import (
	"strings"

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
			RealizedProfit:   utils.StringToFloat(item.CurRealisedPnl),
			Notional:         utils.StringToFloat(item.PositionValue),
			UpdateTime:       utils.StringToInt64(item.UpdatedTime),
		})
	}
	return res
}

func ConvertHistoryOrders(answ []HistoryOrder) (res []entity.OrdersHistory) {
	for _, item := range answ {
		posSide := "LONG"
		if strings.ToUpper(item.Side) == "BUY" {
			posSide = "SHORT"
		}
		res = append(res, entity.OrdersHistory{
			Symbol:  item.Symbol,
			OrderID: item.OrderId,
			// ClientOrderID: item.ClOrdId,
			Side:         strings.ToUpper(item.Side),
			PositionSide: posSide,
			Category:     strings.ToUpper(item.ExecType),
			Price:        utils.StringToFloat(item.AvgEntryPrice),
			FillPrice:    utils.StringToFloat(item.AvgExitPrice),
			Size:         utils.StringToFloat(item.Qty),
			Notional:     utils.StringToFloat(item.CumEntryValue),
			FillSize:     utils.StringToFloat(item.ClosedSize),
			Type:         strings.ToUpper(item.OrderType),
			// Status:        strings.ToUpper(item.State), //entity
			Pnl: utils.StringToFloat(item.ClosedPnl),
			// Fee:           utils.StringToFloat(item.Fee),
			CreateTime: utils.StringToInt64(item.CreatedTime),
			UpdateTime: utils.StringToInt64(item.UpdatedTime),
		})
	}
	return res
}
