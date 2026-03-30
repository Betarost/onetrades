package weex

import (
	"strconv"
	"strings"

	"github.com/Betarost/onetrades/entity"
	"github.com/Betarost/onetrades/utils"
)

// ===============ACCOUNT=================

type account_converts struct{}

// ===============SPOT=================

type spot_converts struct{}

func (c *spot_converts) convertInstrumentsInfo(in []spot_instrumentsInfo) (out []entity.Spot_InstrumentsInfo) {
	if len(in) == 0 {
		return out
	}

	for _, item := range in {
		tickSize := item.TickSize.String()
		stepSize := item.StepSize.String()
		minTradeAmount := item.MinTradeAmount.String()

		out = append(out, entity.Spot_InstrumentsInfo{
			Symbol:         item.Symbol,
			Base:           item.BaseAsset,
			Quote:          item.QuoteAsset,
			MinQty:         minTradeAmount,
			PricePrecision: utils.GetPrecisionFromStr(tickSize),
			SizePrecision:  utils.GetPrecisionFromStr(stepSize),
			State:          strings.ToUpper(item.Status),
		})
	}

	return out
}

func (c *spot_converts) convertBalance(in []spot_Balance) (out []entity.AssetsBalance) {
	if len(in) == 0 {
		return out
	}

	for _, item := range in {
		out = append(out, entity.AssetsBalance{
			Asset:   item.Asset,
			Balance: item.Free,
			Locked:  item.Locked,
		})
	}

	return out
}

func (c *spot_converts) convertPlaceOrder(in placeOrder_Response) (out []entity.PlaceOrder) {
	out = append(out, entity.PlaceOrder{
		OrderID:       strconv.FormatInt(in.OrderId, 10),
		ClientOrderID: in.ClientOrderId,
		Ts:            in.TransactTime,
	})
	return out
}

func (c *spot_converts) convertOrderList(answ []spot_orderList) (res []entity.Spot_OrdersList) {
	for _, item := range answ {
		res = append(res, entity.Spot_OrdersList{
			Symbol:        item.Symbol,
			OrderID:       strconv.FormatInt(item.OrderId, 10),
			ClientOrderID: item.ClientOrderId,
			Size:          item.OrigQty,
			ExecutedSize:  item.ExecutedQty,
			Side:          strings.ToUpper(item.Side),
			Price:         item.Price,
			Type:          strings.ToUpper(item.Type),
			Status:        strings.ToUpper(item.Status),
			CreateTime:    item.Time,
			UpdateTime:    item.UpdateTime,
		})
	}
	return res
}

func (c *spot_converts) convertCancelOrder(in cancelOrder_Response) (out []entity.PlaceOrder) {
	out = append(out, entity.PlaceOrder{
		OrderID:       in.OrderID,
		ClientOrderID: in.ClientOrderID,
	})
	return out
}

func (c *spot_converts) convertOrdersHistory(in []spot_ordersHistory_Response) (out []entity.Spot_OrdersHistory) {
	if len(in) == 0 {
		return out
	}

	for _, item := range in {
		if strings.ToUpper(item.Status) != "FILLED" {
			continue
		}

		executedPrice := item.Price

		executedQty := utils.StringToFloat(item.ExecutedQty)
		cumQuote := utils.StringToFloat(item.CummulativeQuoteQty)
		if executedQty > 0 && cumQuote > 0 {
			executedPrice = strconv.FormatFloat(cumQuote/executedQty, 'f', -1, 64)
		}

		out = append(out, entity.Spot_OrdersHistory{
			Symbol:        item.Symbol,
			OrderID:       strconv.FormatInt(item.OrderId, 10),
			ClientOrderID: item.ClientOrderId,
			Side:          strings.ToUpper(item.Side),
			Size:          item.OrigQty,
			ExecutedSize:  item.ExecutedQty,
			Price:         item.Price,
			ExecutedPrice: executedPrice,
			Type:          strings.ToUpper(item.Type),
			Status:        strings.ToUpper(item.Status),
			CreateTime:    item.Time,
			UpdateTime:    item.UpdateTime,
		})
	}

	return out
}

// ===============FUTURES=================

type futures_converts struct{}

func (c *futures_converts) convertInstrumentsInfo(in []futures_instrumentsInfo) (out []entity.Futures_InstrumentsInfo) {
	if len(in) == 0 {
		return out
	}

	for _, item := range in {

		out = append(out, entity.Futures_InstrumentsInfo{
			Symbol:         item.Symbol,
			Base:           item.BaseAsset,
			Quote:          item.QuoteAsset,
			MinQty:         item.MinOrderSize.String(),
			PricePrecision: strconv.Itoa(item.PricePrecision),
			SizePrecision:  utils.GetPrecisionFromStr(item.MinOrderSize.String()),
			// SizePrecision:  strconv.Itoa(item.QuantityPrecision),
			MaxLeverage:    strconv.Itoa(item.MaxLeverage),
			State:          strings.ToUpper("LIVE"),
			IsSizeContract: false,
			// Multiplier:     "1",
			// ContractSize:   item.ContractVal.String(),
		})
	}

	return out
}

func (c *futures_converts) convertBalance(in []futures_Balance) (out []entity.FuturesBalance) {
	if len(in) == 0 {
		return out
	}

	for _, item := range in {
		out = append(out, entity.FuturesBalance{
			Asset:            item.Asset,
			Balance:          item.Balance,
			Equity:           item.Balance,
			Available:        item.AvailableBalance,
			UnrealizedProfit: item.UnrealizePnl,
		})
	}
	return out
}

func (c *futures_converts) convertPlaceOrder(in futures_placeOrder_Response) (out []entity.PlaceOrder) {
	out = append(out, entity.PlaceOrder{
		OrderID:       in.OrderId,
		ClientOrderID: in.ClientOrderId,
		Ts:            utils.CurrentTimestamp(),
	})
	return out
}

func (c *futures_converts) convertPositions(answ []futures_Position) (res []entity.Futures_Positions) {
	for _, item := range answ {
		hedgeMode := false
		if item.SeparatedMode == "SEPARATED" {
			hedgeMode = true
		}

		res = append(res, entity.Futures_Positions{
			Symbol:           item.Symbol,
			PositionSide:     strings.ToUpper(item.Side),
			PositionSize:     item.Size,
			Leverage:         item.Leverage,
			PositionID:       utils.Int64ToString(item.ID),
			EntryPrice:       "",
			MarkPrice:        "",
			UnRealizedProfit: item.UnrealizePnl,
			RealizedProfit:   "",
			Notional:         item.OpenValue,
			HedgeMode:        hedgeMode,
			MarginMode:       strings.ToUpper(item.MarginType),
			CreateTime:       item.CreatedTime,
			UpdateTime:       item.UpdatedTime,
		})
	}
	return res
}

func (c *futures_converts) convertOrderList(answ []futures_orderList) (res []entity.Futures_OrdersList) {
	for _, item := range answ {
		res = append(res, entity.Futures_OrdersList{
			Symbol:        item.Symbol,
			OrderID:       strconv.FormatInt(item.OrderId, 10),
			ClientOrderID: item.ClientOrderId,
			PositionSide:  strings.ToUpper(item.PositionSide),
			Side:          strings.ToUpper(item.Side),
			PositionSize:  item.OrigQty,
			ExecutedSize:  item.ExecutedQty,
			Price:         item.Price,
			Type:          strings.ToUpper(item.Type),
			Status:        strings.ToUpper(item.Status),
			CreateTime:    item.Time,
			UpdateTime:    item.UpdateTime,
		})
	}
	return res
}

func (c *futures_converts) convertCancelOrder(in futures_cancelOrder_Response) (out []entity.PlaceOrder) {
	out = append(out, entity.PlaceOrder{
		OrderID:       in.OrderId,
		ClientOrderID: in.OrigClientOrderId,
		Ts:            utils.CurrentTimestamp(),
	})
	return out
}

func (c *futures_converts) convertOrdersHistory(in []futures_ordersHistory_Response) (out []entity.Futures_OrdersHistory) {
	if len(in) == 0 {
		return out
	}

	for _, item := range in {
		if strings.ToUpper(item.Status) != "FILLED" {
			continue
		}

		hedgeMode := false
		if item.PositionSide != "" && strings.ToUpper(item.PositionSide) != "BOTH" {
			hedgeMode = true
		}

		out = append(out, entity.Futures_OrdersHistory{
			Symbol:         item.Symbol,
			OrderID:        utils.Int64ToString(item.OrderId),
			ClientOrderID:  item.ClientOrderId,
			Side:           strings.ToUpper(item.Side),
			PositionSide:   strings.ToUpper(item.PositionSide),
			PositionSize:   item.OrigQty,
			ExecutedSize:   item.ExecutedQty,
			Price:          item.Price,
			ExecutedPrice:  item.AvgPrice,
			RealisedProfit: "",
			Fee:            "",
			FeeAsset:       "",
			Leverage:       "",
			HedgeMode:      hedgeMode,
			MarginMode:     "",
			Type:           strings.ToUpper(item.Type),
			Status:         strings.ToUpper(item.Status),
			CreateTime:     item.Time,
			UpdateTime:     item.UpdateTime,
		})
	}
	return out
}

func (c *futures_converts) convertLeverage(in []futures_leverage) (out entity.Futures_Leverage) {
	if len(in) == 0 {
		return out
	}

	item := in[0]
	out.Symbol = strings.ToUpper(item.Symbol)

	if strings.ToUpper(item.MarginType) == "ISOLATED" {
		out.MarginMode = string(entity.MarginModeTypeIsolated)
		out.LongLeverage = item.IsolatedLongLeverage
		out.ShortLeverage = item.IsolatedShortLeverage

		if item.IsolatedLongLeverage == item.IsolatedShortLeverage {
			out.Leverage = item.IsolatedLongLeverage
		} else if utils.StringToFloat(item.IsolatedLongLeverage) < utils.StringToFloat(item.IsolatedShortLeverage) {
			out.Leverage = item.IsolatedLongLeverage
		} else {
			out.Leverage = item.IsolatedShortLeverage
		}
	} else {
		out.MarginMode = string(entity.MarginModeTypeCross)
		out.Leverage = item.CrossLeverage
		out.LongLeverage = item.CrossLeverage
		out.ShortLeverage = item.CrossLeverage
	}

	return out
}
