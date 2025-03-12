package futurebingx

import (
	"fmt"

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

func ConvertOrdersHistory(answ []OrdersHistory) (res []entity.OrdersHistory) {
	for _, item := range answ {
		res = append(res, entity.OrdersHistory{
			Symbol:       item.Symbol,
			OrderID:      fmt.Sprintf("%d", item.OrderID),
			Side:         item.Side,
			PositionSide: item.PositionSide,
			Price:        utils.StringToFloat(item.Price),
			// OrigQty:      utils.StringToFloat(item.OrigQty),
			// AvgPrice: utils.StringToFloat(item.AvgPrice),
			Type: item.Type,
			// Status:       item.Status,
			// Time:       item.Time,
			UpdateTime: item.UpdateTime,
		})
	}
	return res
}
