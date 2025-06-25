package okx

import (
	"strconv"
	"strings"

	"github.com/Betarost/onetrades/entity"
	"github.com/Betarost/onetrades/utils"
)

// ===============ACCOUNT=================
type account_converts struct{}

func (c *account_converts) convertAccountInfo(in accountInfo) (out entity.AccountInformation) {

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
		// out.HedgeMode = true
	}

	if in.AcctLv != "1" {
		out.PermFutures = true
	}
	return out
}

// ===============SPOT=================

type spot_converts struct{}

func (c *spot_converts) convertInstrumentsInfo(in []spot_instrumentsInfo) (out []entity.Spot_InstrumentsInfo) {
	if len(in) == 0 {
		return out
	}
	for _, item := range in {

		priceP := utils.GetPrecisionFromStr(item.TickSz)
		sizeP := utils.GetPrecisionFromStr(item.LotSz)

		out = append(out, entity.Spot_InstrumentsInfo{
			Symbol: item.InstId,
			Base:   item.BaseCcy,
			Quote:  item.QuoteCcy,
			MinQty: item.MinSz,
			// MinNotional: item,
			PricePrecision: priceP,
			SizePrecision:  sizeP,
			State:          strings.ToUpper(item.State),
		})
	}
	return out
}

func (c *spot_converts) convertBalance(in []spot_Balance) (out []entity.AssetsBalance) {
	if len(in) == 0 {
		return out
	}

	for _, i := range in {
		for _, item := range i.Details {
			out = append(out, entity.AssetsBalance{
				Asset:   item.Ccy,
				Balance: item.Eq,
				Locked:  item.FrozenBal,
			})
		}
	}
	return out
}

// =======OLD

func convertAccountInfo(in accountInfo) (out entity.AccountInformation) {

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
		// out.HedgeMode = true
	}

	if in.AcctLv != "1" {
		out.PermFutures = true
	}
	return out
}

func convertTradingAccountBalance(in []tradingBalance) (out []entity.TradingAccountBalance) {
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

func convertFundingAccountBalance(in []fundingBalance) (out []entity.FundingAccountBalance) {
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

func convertInstrumentsInfo(in []spot_instrumentsInfo) (out []entity.InstrumentsInfo) {
	if len(in) == 0 {
		return out
	}
	for _, item := range in {
		out = append(out, entity.InstrumentsInfo{
			Symbol:             item.InstId,
			ContractSize:       item.CtVal,
			ContractMultiplier: item.CtMult,
			StepContractSize:   item.LotSz,
			MinContractSize:    item.MinSz,
			StepTickPrice:      item.TickSz,
			MaxLeverage:        item.Lever,
			State:              strings.ToUpper(item.State),
			// InstType:           item.InstType,
			Base:  item.BaseCcy,
			Quote: item.QuoteCcy,
		})
	}
	return out
}

func futures_convertInstrumentsInfo(in []futures_instrumentsInfo) (out []entity.Futures_InstrumentsInfo) {
	if len(in) == 0 {
		return out
	}
	for _, item := range in {
		out = append(out, entity.Futures_InstrumentsInfo{
			Symbol:             item.InstId,
			ContractSize:       item.CtVal,
			ContractMultiplier: item.CtMult,
			StepContractSize:   item.LotSz,
			MinContractSize:    item.MinSz,
			StepTickPrice:      item.TickSz,
			MaxLeverage:        item.Lever,
			State:              strings.ToUpper(item.State),
			// InstType:           item.InstType,
			Base:  item.BaseCcy,
			Quote: item.QuoteCcy,
		})
	}
	return out
}

func futures_convertSetLeverage(in []setLeverageAnswer) (out entity.Futures_Leverage) {
	if len(in) == 0 {
		return out
	}
	for _, item := range in {
		out.Symbol = item.InstId
		out.Leverage = item.Lever
		out.MarginMode = item.MgnMode
		out.PositionSide = item.PosSide
	}
	return out
}

func futures_convertPositions(answ []futures_Position) (res []entity.Futures_Positions) {
	for _, item := range answ {
		positionSide := "LONG"
		if item.PosSide == "net" {
			if utils.StringToFloat(item.Pos) < 0 {
				positionSide = "SHORT"
			}
		} else {
			positionSide = strings.ToUpper(item.PosSide)
		}

		res = append(res, entity.Futures_Positions{
			Symbol:           item.InstID,
			PositionSide:     positionSide,
			PositionID:       item.PosID,
			PositionAmt:      item.Pos,
			EntryPrice:       item.AvgPx,
			MarkPrice:        item.MarkPx,
			InitialMargin:    item.Imr,
			UnRealizedProfit: item.Upl,
			RealizedProfit:   item.RealizedPnl,
			Notional:         item.NotionalUsd,
			MarginRatio:      item.MgnRatio,
			AutoDeleveraging: item.ADL,
			UpdateTime:       utils.StringToInt64(item.UTime),
		})
	}
	return res
}

func convertOrderList(answ []orderList) (res []entity.OrdersPendingList) {
	for _, item := range answ {
		positionSide := "LONG"
		if item.PosSide == "net" {
			if strings.ToUpper(item.Side) == "SELL" {
				positionSide = "SHORT"
			}
		} else {
			positionSide = strings.ToUpper(item.PosSide)
		}

		tp := ""
		sl := ""
		if len(item.AttachAlgoOrds) > 0 {
			if item.AttachAlgoOrds[0].TpOrdPx != "-1" && item.AttachAlgoOrds[0].TpOrdPx != "" {
				tp = item.AttachAlgoOrds[0].TpOrdPx
			} else if item.AttachAlgoOrds[0].TpTriggerPx != "-1" && item.AttachAlgoOrds[0].TpTriggerPx != "" {
				tp = item.AttachAlgoOrds[0].TpTriggerPx
			}

			if item.AttachAlgoOrds[0].SlOrdPx != "-1" && item.AttachAlgoOrds[0].SlOrdPx != "" {
				sl = item.AttachAlgoOrds[0].SlOrdPx
			} else if item.AttachAlgoOrds[0].SlTriggerPx != "-1" && item.AttachAlgoOrds[0].SlTriggerPx != "" {
				sl = item.AttachAlgoOrds[0].SlTriggerPx
			}
		}
		b, _ := strconv.ParseBool(item.IsTpLimit)
		res = append(res, entity.OrdersPendingList{
			Symbol:        item.InstId,
			OrderID:       item.OrdId,
			ClientOrderID: item.ClOrdId,
			PositionSide:  positionSide,
			Side:          item.Side,
			PositionAmt:   item.Sz,
			Price:         item.Px,
			TpPrice:       tp,
			SlPrice:       sl,
			Type:          strings.ToUpper(item.OrdType),
			TradeMode:     item.TdMode,
			InstType:      item.InstType,
			Leverage:      item.Lever,
			Status:        item.State,
			IsTpLimit:     b,
			CreateTime:    utils.StringToInt64(item.CTime),
			UpdateTime:    utils.StringToInt64(item.UTime),
		})
	}
	return res
}

func futures_convertOrderList(answ []futures_orderList) (res []entity.Futures_OrdersList) {
	for _, item := range answ {
		positionSide := "LONG"
		if item.PosSide == "net" {
			if strings.ToUpper(item.Side) == "SELL" {
				positionSide = "SHORT"
			}
		} else {
			positionSide = strings.ToUpper(item.PosSide)
		}

		tp := ""
		sl := ""
		if len(item.AttachAlgoOrds) > 0 {
			if item.AttachAlgoOrds[0].TpOrdPx != "-1" && item.AttachAlgoOrds[0].TpOrdPx != "" {
				tp = item.AttachAlgoOrds[0].TpOrdPx
			} else if item.AttachAlgoOrds[0].TpTriggerPx != "-1" && item.AttachAlgoOrds[0].TpTriggerPx != "" {
				tp = item.AttachAlgoOrds[0].TpTriggerPx
			}

			if item.AttachAlgoOrds[0].SlOrdPx != "-1" && item.AttachAlgoOrds[0].SlOrdPx != "" {
				sl = item.AttachAlgoOrds[0].SlOrdPx
			} else if item.AttachAlgoOrds[0].SlTriggerPx != "-1" && item.AttachAlgoOrds[0].SlTriggerPx != "" {
				sl = item.AttachAlgoOrds[0].SlTriggerPx
			}
		}
		b, _ := strconv.ParseBool(item.IsTpLimit)
		res = append(res, entity.Futures_OrdersList{
			Symbol:        item.InstId,
			OrderID:       item.OrdId,
			ClientOrderID: item.ClOrdId,
			PositionSide:  positionSide,
			Side:          item.Side,
			PositionAmt:   item.Sz,
			Price:         item.Px,
			TpPrice:       tp,
			SlPrice:       sl,
			Type:          strings.ToUpper(item.OrdType),
			TradeMode:     item.TdMode,
			InstType:      item.InstType,
			Leverage:      item.Lever,
			Status:        item.State,
			IsTpLimit:     b,
			CreateTime:    utils.StringToInt64(item.CTime),
			UpdateTime:    utils.StringToInt64(item.UTime),
		})
	}
	return res
}

func convertPlaceOrder(in []placeOrder_Response) (out []entity.PlaceOrder) {
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
