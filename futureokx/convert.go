package futureokx

import (
	"strings"

	"github.com/Betarost/onetrades/entity"
	"github.com/Betarost/onetrades/utils"
)

func ConvertTickers(answ []Ticker) (res []entity.Ticker) {
	if len(answ) == 0 {
		return res
	}
	for _, item := range answ {
		res = append(res, entity.Ticker{
			Symbol:        item.InstId,
			Open24hPrice:  utils.StringToFloat(item.Open24h),
			LastPrice:     utils.StringToFloat(item.Last),
			Volume24hCoin: utils.StringToFloat(item.VolCcy24h),
			Volume24hUSDT: utils.StringToFloat(item.VolCcy24h) * utils.StringToFloat(item.Last),
			// Change24h:     (utils.StringToFloat(item.Last) - utils.StringToFloat(item.Open24h)) / utils.StringToFloat(item.Open24h) * 100,
			Change24h: (utils.StringToFloat(item.Last) - utils.StringToFloat(item.SodUtc0)) / utils.StringToFloat(item.SodUtc0) * 100,
			Time:      utils.StringToInt64(item.Ts),
		})
	}
	return res
}

func ConvertTransferHistory(answ []TransferHistory) (res []entity.TransferHistory) {
	for _, item := range answ {
		t := "TO"
		if item.Type == "1" {
			t = "FROM"
		}
		res = append(res, entity.TransferHistory{
			Asset:      item.Ccy,
			SubID:      item.SubAcct,
			BillID:     item.BillId,
			Amount:     utils.StringToFloat(item.Amt),
			Type:       t,
			CreateTime: utils.StringToInt64(item.Ts),
		})
	}
	return res
}

func ConvertAccountValuation(answ AccountValuation) (res entity.AccountValuation) {

	res.ClassicBalance = utils.StringToFloat(answ.Details.Classic)
	res.EarnBalance = utils.StringToFloat(answ.Details.Earn)
	res.FundingBalance = utils.StringToFloat(answ.Details.Funding)
	res.TradingBalance = utils.StringToFloat(answ.Details.Trading)
	res.TotalBalance = utils.StringToFloat(answ.TotalBal)
	res.UpdateTime = utils.StringToInt64(answ.Ts)
	return res
}

func ConvertSubAccountFundingBalance(answ SubAccountFundingBalance) (res entity.SubAccountFundingBalance) {
	res.Asset = answ.Ccy
	res.Balance = utils.StringToFloat(answ.Bal)
	res.AvailableBalance = utils.StringToFloat(answ.AvailBal)
	res.FrozenBalance = utils.StringToFloat(answ.FrozenBal)
	return res
}

func ConvertSubAccountBalance(answ SubAccountBalance) (res entity.SubAccountBalance) {
	res.EquityBalance = utils.StringToFloat(answ.TotalEq)
	res.UnrealizedProfit = utils.StringToFloat(answ.Upl)
	return res
}

func ConvertSubAccountInfo(answ []SubAccountsLists) (res []entity.AccountInfo) {
	if len(answ) == 0 {
		return res
	}
	for _, item := range answ {
		canTrade := true
		canTransfer := true
		level := item.Type
		if item.Type == "1" {
			level = "Standard sub-account"
		}
		for _, i := range item.FrozenFunc {
			if i == "trading" {
				canTrade = false
			} else if i == "transfer" {
				canTransfer = false
			}
		}
		res = append(res, entity.AccountInfo{
			UID:         item.UID,
			Name:        item.SubAcct,
			Label:       item.Label,
			Level:       level,
			CanRead:     true,
			CanTrade:    canTrade,
			CanTransfer: canTransfer,
			HedgeMode:   true,
		})
	}
	return res
}

func ConvertAccountInfo(answ AccountInfo) (res entity.AccountInfo) {

	res.UID = answ.UID
	res.MainUID = answ.MainUID
	res.Label = answ.Label
	res.Level = answ.Level

	if answ.UID == answ.MainUID {
		res.IsMain = true
	}

	if strings.Contains(answ.Perm, "read") {
		res.CanRead = true
	}

	if strings.Contains(answ.Perm, "trade") {
		res.CanTrade = true
	}

	if answ.PosMode == "long_short_mode" {
		res.HedgeMode = true
	}

	return res
}

func ConvertAccountBalance(answ []Balance) (res []entity.AccountBalance) {
	if len(answ) == 0 {
		return res
	}
	for _, item := range answ[0].Details {
		res = append(res, entity.AccountBalance{
			Asset: item.Ccy,
			// Balance:          utils.StringToFloat(item.Eq),
			Balance:          utils.StringToFloat(item.CashBal),
			EquityBalance:    utils.StringToFloat(item.Eq),
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

func ConvertHistoryOrders(answ []HistoryOrder) (res []entity.OrdersHistory) {
	for _, item := range answ {
		res = append(res, entity.OrdersHistory{
			Symbol:        item.InstID,
			OrderID:       item.OrdId,
			ClientOrderID: item.ClOrdId,
			Side:          strings.ToUpper(item.Side),
			PositionSide:  strings.ToUpper(item.PosSide),
			Category:      strings.ToUpper(item.Category),
			Price:         utils.StringToFloat(item.Px),
			FillPrice:     utils.StringToFloat(item.AvgPx),
			Size:          utils.StringToFloat(item.Sz),
			FillSize:      utils.StringToFloat(item.AccFillSz),
			Type:          strings.ToUpper(item.OrdType),
			Status:        strings.ToUpper(item.State), //entity
			Pnl:           utils.StringToFloat(item.Pnl),
			Fee:           utils.StringToFloat(item.Fee),
			CreateTime:    utils.StringToInt64(item.CTime),
			UpdateTime:    utils.StringToInt64(item.UTime),
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
			Symbol:             item.InstId,
			ContractSize:       utils.StringToFloat(item.CtVal),
			ContractMultiplier: utils.StringToFloat(item.CtMult),
			StepContractSize:   utils.StringToFloat(item.LotSz),
			MinContractSize:    utils.StringToFloat(item.MinSz),
			StepTickPrice:      utils.StringToFloat(item.TickSz),
			MaxLeverage:        utils.StringToInt(item.Lever),
			State:              strings.ToUpper(item.State),
			Type:               strings.ToUpper(item.RuleType),
		})
	}
	return res
}

func ConvertMarkPrices(answ []MarkPrice) (res []entity.MarkPrice) {
	if len(answ) == 0 {
		return res
	}
	for _, item := range answ {
		res = append(res, entity.MarkPrice{
			Symbol: item.InstId,
			Price:  utils.StringToFloat(item.MarkPx),
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
