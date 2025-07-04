package okx

import (
	"strconv"
	"strings"
	"time"

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

func (c *spot_converts) convertOrdersHistory(in []spot_ordersHistory_Response) (out []entity.Spot_OrdersHistory) {
	if len(in) == 0 {
		return out
	}

	for _, item := range in {
		out = append(out, entity.Spot_OrdersHistory{
			Symbol:        item.InstId,
			OrderID:       item.OrdId,
			ClientOrderID: item.ClOrdId,
			Side:          strings.ToUpper(item.Side),
			Size:          item.Sz,
			Price:         item.Px,
			ExecutedSize:  item.AccFillSz,
			ExecutedPrice: item.AvgPx,
			Fee:           item.Fee,
			Type:          strings.ToUpper(item.OrdType),
			Status:        strings.ToUpper(item.State),
			CreateTime:    utils.StringToInt64(item.CTime),
			UpdateTime:    utils.StringToInt64(item.UTime),
		})
	}
	return out
}

func (c *spot_converts) convertPlaceOrder(in []placeOrder_Response) (out []entity.PlaceOrder) {
	if len(in) == 0 {
		return out
	}
	for _, item := range in {
		out = append(out, entity.PlaceOrder{
			OrderID:       item.OrdId,
			ClientOrderID: item.ClOrdId,
			Ts:            time.Now().UTC().UnixMilli(),
		})
	}
	return out
}

func (c *spot_converts) convertOrderList(answ []spot_orderList) (res []entity.Spot_OrdersList) {
	for _, item := range answ {
		res = append(res, entity.Spot_OrdersList{
			Symbol:        item.InstId,
			OrderID:       item.OrdId,
			ClientOrderID: item.ClOrdId,
			Size:          item.Sz,
			ExecutedSize:  item.FillSz,
			Side:          strings.ToUpper(item.Side),
			Price:         item.Px,
			Type:          strings.ToUpper(item.OrdType),
			Status:        item.State,
			CreateTime:    utils.StringToInt64(item.CTime),
			UpdateTime:    utils.StringToInt64(item.UTime),
		})
	}
	return res
}

// ===============FUTURES=================

type futures_converts struct{}

func (c *futures_converts) convertInstrumentsInfo(in []futures_instrumentsInfo) (out []entity.Futures_InstrumentsInfo) {
	if len(in) == 0 {
		return out
	}
	for _, item := range in {
		out = append(out, entity.Futures_InstrumentsInfo{
			Symbol:         item.InstId,
			Base:           item.CtValCcy,
			Quote:          item.SettleCcy,
			MinQty:         item.MinSz,
			PricePrecision: utils.GetPrecisionFromStr(item.TickSz),
			SizePrecision:  utils.GetPrecisionFromStr(item.MinSz),
			MaxLeverage:    item.Lever,
			State:          strings.ToUpper(item.State),
			IsSizeContract: true,
			Multiplier:     item.CtMult,
			ContractSize:   item.CtVal,
		})
	}
	return out
}

func (c *futures_converts) convertBalance(in []futures_Balance) (out []entity.FuturesBalance) {
	if len(in) == 0 {
		return out
	}

	for _, i := range in {
		for _, item := range i.Details {
			out = append(out, entity.FuturesBalance{
				Asset:            item.Ccy,
				Balance:          item.CashBal,
				Equity:           item.AvailEq,
				Available:        item.AvailBal,
				UnrealizedProfit: item.Upl,
			})
		}
	}
	return out
}

func (c *futures_converts) convertLeverage(in futures_leverage) (out entity.Futures_Leverage) {

	out.Symbol = in.InstId
	out.Leverage = in.Lever
	return out
}

func (c *futures_converts) convertPositionsHistory(in []futures_PositionsHistory_Response) (out []entity.Futures_PositionsHistory) {
	if len(in) == 0 {
		return out
	}

	for _, item := range in {
		mMode := "cross"
		if item.MgnMode != "cross" {
			mMode = "isolated"
		}
		out = append(out, entity.Futures_PositionsHistory{
			Symbol:              item.InstId,
			PositionId:          item.PosId,
			PositionSide:        strings.ToUpper(item.PosSide),
			PositionAmt:         item.OpenMaxPos,
			ExecutedPositionAmt: item.CloseTotalPos,
			AvgPrice:            item.OpenAvgPx,
			ExecutedAvgPrice:    item.CloseAvgPx,
			RealisedProfit:      item.Pnl,
			Fee:                 item.Fee,
			Funding:             item.FundingFee,
			MarginMode:          mMode,
			CreateTime:          utils.StringToInt64(item.CTime),
			UpdateTime:          utils.StringToInt64(item.UTime),
		})
	}
	return out
}

// =======OLD

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
