package futurebinance

import (
	"fmt"
	"strings"

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

func ConvertHistoryOrders(answ []HistoryOrder) (res []entity.OrdersHistory) {
	for _, item := range answ {

		t := "LIMIT"
		if item.Buyer {
			t = "MARKET"
		}
		res = append(res, entity.OrdersHistory{
			Symbol:  item.Symbol,
			OrderID: fmt.Sprintf("%d", item.Id),
			// ClientOrderID: item.ClOrdId,
			Side:         strings.ToUpper(item.Side),
			PositionSide: strings.ToUpper(item.PositionSide),
			// Category:     strings.ToUpper(item.Category),
			Price:     utils.StringToFloat(item.Price),
			FillPrice: utils.StringToFloat(item.Price),
			Size:      utils.StringToFloat(item.Qty),
			FillSize:  utils.StringToFloat(item.Qty),
			Notional:  utils.StringToFloat(item.QuoteQty),
			Type:      t,
			// Status:       strings.ToUpper(item.State), //entity
			Pnl:        utils.StringToFloat(item.RealizedPnl),
			Fee:        0 - utils.StringToFloat(item.Commission),
			CreateTime: item.Time,
			UpdateTime: item.Time,
		})
	}
	return res
}
