package okx

import (
	"strings"

	"github.com/Betarost/onetrades/entity"
	"github.com/Betarost/onetrades/utils"
)

func ConvertAccountInfo(in accountInfo) (out entity.AccountInformation) {

	out.UID = in.UID
	out.Label = in.Label
	out.IP = in.Ip
	out.PermSpot = true

	if strings.Contains(in.Perm, "read") {
		out.CanRead = true
	}

	if strings.Contains(in.Perm, "trade") {
		out.CanTrade = true
	}

	if in.PosMode == "long_short_mode" {
		out.HedgeMode = true
	}

	if in.AcctLv != "1" {
		out.PermFutures = true
	}
	return out
}

func ConvertTradingAccountBalance(in []tradingBalance) (out []entity.TradingAccountBalance) {
	if len(in) == 0 {
		return out
	}
	for _, item := range in {
		r := entity.TradingAccountBalance{
			TotalEquity:      item.TotalEq,
			AvailableEquity:  item.AvailEq,
			NotionalUsd:      item.NotionalUsd,
			UnrealizedProfit: item.Upl,
			UpdateTime:       utils.StringToInt64(item.UTime),
		}
		for _, item := range item.Details {
			r.Assets = append(r.Assets, entity.TradingAccountBalanceDetails{
				Asset:            item.Ccy,
				Balance:          item.CashBal,
				EquityBalance:    item.Eq,
				AvailableBalance: item.AvailBal,
				AvailableEquity:  item.AvailEq,
				UnrealizedProfit: item.Upl,
			})
		}

		out = append(out, r)
	}
	return out
}

func ConvertFundingAccountBalance(in []fundingBalance) (out []entity.FundingAccountBalance) {
	if len(in) == 0 {
		return out
	}

	for _, item := range in {
		out = append(out, entity.FundingAccountBalance{
			Asset:            item.Ccy,
			Balance:          item.Bal,
			AvailableBalance: item.AvailBal,
			FrozenBalance:    item.FrozenBal,
		})
	}

	return out
}

func ConvertOrderList(answ []orderList) (res []entity.OrdersPendingList) {
	for _, item := range answ {
		positionSide := "LONG"
		if item.PosSide == "net" {
			if strings.ToUpper(item.Side) == "SELL" {
				positionSide = "SHORT"
			}
		} else {
			positionSide = strings.ToUpper(item.PosSide)
		}

		res = append(res, entity.OrdersPendingList{
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

func ConvertPlaceOrder(in []placeOrder_Response) (out []entity.PlaceOrder) {
	if len(in) == 0 {
		return out
	}
	for _, item := range in {
		out = append(out, entity.PlaceOrder{
			OrderID:       item.OrdId,
			ClientOrderID: item.ClOrdId,
			Ts:            utils.StringToInt64(item.Ts),
		})
	}
	return out
}
