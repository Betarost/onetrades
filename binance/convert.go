package binance

import (
	"strings"
	"time"

	"github.com/Betarost/onetrades/entity"
	"github.com/Betarost/onetrades/utils"
)

// ===============ACCOUNT=================
type account_converts struct{}

func (c *account_converts) convertAccountInfo(in accountInfo) (out entity.AccountInformation) {

	out.UID = utils.Int64ToString(in.UID)
	out.CanRead = true
	out.CanTrade = in.CanTrade
	out.CanTransfer = in.CanWithdraw
	// out.Label = in.Label
	// out.IP = in.Ip
	out.PermSpot = true
	out.PermFutures = true

	return out
}

// ===============SPOT=================

type spot_converts struct{}

func (c *spot_converts) convertBalance(in spot_Balance) (out []entity.AssetsBalance) {
	if len(in.Balances) == 0 {
		return out
	}

	for _, item := range in.Balances {
		if utils.StringToFloat(item.Free) == 0 && utils.StringToFloat(item.Locked) == 0 {
			continue
		}
		out = append(out, entity.AssetsBalance{
			Asset:   item.Asset,
			Balance: utils.FloatToStringAll(utils.StringToFloat(item.Free) + utils.StringToFloat(item.Locked)),
			Locked:  item.Locked,
		})
	}
	return out
}

func (c *spot_converts) convertInstrumentsInfo(in spot_instrumentsInfo) (out []entity.Spot_InstrumentsInfo) {
	if len(in.Symbols) == 0 {
		return out
	}
	for _, item := range in.Symbols {
		if item.Status == "TRADING" {
			item.Status = "LIVE"
		}
		rec := entity.Spot_InstrumentsInfo{
			Symbol: item.Symbol,
			Base:   item.BaseAsset,
			Quote:  item.QuoteAsset,
			// SizePrecision: utils.Int64ToString(item.BaseAssetPrecision),
			State: item.Status,
		}
		for _, i := range item.Filters {
			m := i.(map[string]interface{})
			switch m["filterType"] {
			case "PRICE_FILTER":
				rec.PricePrecision = utils.GetPrecisionFromStr(utils.FloatToStringAll(utils.StringToFloat(m["tickSize"].(string))))
			case "LOT_SIZE":
				rec.SizePrecision = utils.GetPrecisionFromStr(utils.FloatToStringAll(utils.StringToFloat(m["minQty"].(string))))
				rec.MinQty = utils.FloatToStringAll(utils.StringToFloat(m["minQty"].(string)))
				// rec.MinQty = m["minQty"].(string)
				// rec.StepContractSize = m["stepSize"].(string)
			case "NOTIONAL":
				rec.MinNotional = utils.FloatToStringAll(utils.StringToFloat(m["minNotional"].(string)))
			}
		}
		out = append(out, rec)
	}
	return
}

func (c *spot_converts) convertPlaceOrder(in placeOrder_Response) (out []entity.PlaceOrder) {

	out = append(out, entity.PlaceOrder{
		OrderID:       utils.Int64ToString(in.OrderId),
		ClientOrderID: in.ClientOrderId,
		Ts:            time.Now().UTC().UnixMilli(),
	})
	return out
}

func (c *spot_converts) convertOrderList(answ []spot_orderList) (res []entity.Spot_OrdersList) {
	for _, item := range answ {
		res = append(res, entity.Spot_OrdersList{
			Symbol:        item.Symbol,
			OrderID:       utils.Int64ToString(item.OrderId),
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

func (c *spot_converts) convertOrdersHistory(in []spot_ordersHistory_Response) (out []entity.Spot_OrdersHistory) {
	if len(in) == 0 {
		return out
	}

	for _, item := range in {
		out = append(out, entity.Spot_OrdersHistory{
			Symbol:        item.Symbol,
			OrderID:       utils.Int64ToString(item.OrderId),
			ClientOrderID: item.ClientOrderId,
			Side:          strings.ToUpper(item.Side),
			Size:          item.OrigQty,
			Price:         item.Price,
			ExecutedSize:  item.ExecutedQty,
			// ExecutedPrice: item.AvgPx,
			// Fee:           item.Fee,
			Type:       strings.ToUpper(item.Type),
			Status:     strings.ToUpper(item.Status),
			CreateTime: item.Time,
			UpdateTime: item.UpdateTime,
		})
	}
	return out
}

// ===============FUTURES=================

type futures_converts struct{}

func (c *futures_converts) convertBalance(in futures_Balance) (out []entity.FuturesBalance) {
	if len(in.Assets) == 0 {
		return out
	}

	for _, item := range in.Assets {

		if utils.StringToFloat(item.WalletBalance) == 0 && utils.StringToFloat(item.AvailableBalance) == 0 {
			continue
		}

		out = append(out, entity.FuturesBalance{
			Asset:   item.Asset,
			Balance: item.WalletBalance,
			// Equity:           item.AvailEq,
			Available:        item.AvailableBalance,
			UnrealizedProfit: item.UnrealizedProfit,
		})
	}
	return out
}

func (c *futures_converts) convertInstrumentsInfo(in futures_instrumentsInfo) (out []entity.Futures_InstrumentsInfo) {
	if len(in.Symbols) == 0 {
		return out
	}
	for _, item := range in.Symbols {
		if item.Status == "TRADING" {
			item.Status = "LIVE"
		}
		rec := entity.Futures_InstrumentsInfo{
			Symbol: item.Symbol,
			Base:   item.BaseAsset,
			Quote:  item.QuoteAsset,
			// SizePrecision: utils.Int64ToString(item.BaseAssetPrecision),
			State: item.Status,
		}
		for _, i := range item.Filters {
			m := i.(map[string]interface{})
			switch m["filterType"] {
			case "PRICE_FILTER":
				rec.PricePrecision = utils.GetPrecisionFromStr(utils.FloatToStringAll(utils.StringToFloat(m["tickSize"].(string))))
			case "LOT_SIZE":
				rec.MinQty = utils.FloatToStringAll(utils.StringToFloat(m["minQty"].(string)))
				rec.SizePrecision = utils.GetPrecisionFromStr(utils.FloatToStringAll(utils.StringToFloat(m["minQty"].(string))))
			case "MIN_NOTIONAL":
				rec.MinNotional = utils.FloatToStringAll(utils.StringToFloat(m["notional"].(string)))
			}
		}
		out = append(out, rec)
	}
	return
}

func (c *futures_converts) convertLeverage(in futures_leverage) (out entity.Futures_Leverage) {

	out.Symbol = in.Symbol
	out.Leverage = utils.Int64ToString(in.Leverage)
	return out
}

func (c *futures_converts) convertPlaceOrder(in futures_placeOrder_Response) (out []entity.PlaceOrder) {

	out = append(out, entity.PlaceOrder{
		OrderID:       utils.Int64ToString(in.OrderId),
		ClientOrderID: in.ClientOrderId,
		Ts:            time.Now().UTC().UnixMilli(),
	})

	return out
}

func (c *futures_converts) convertPositions(answ []futures_Position) (res []entity.Futures_Positions) {
	for _, item := range answ {
		positionSide := "LONG"
		hedgeMode := false
		marginMode := "cross"

		if item.PositionSide == "BOTH" {
			if utils.StringToFloat(item.PositionAmt) < 0 {
				positionSide = "SHORT"
			}
		} else {
			positionSide = strings.ToUpper(item.PositionSide)
		}

		if item.PositionSide != "BOTH" {
			hedgeMode = true
		}

		if utils.StringToFloat(item.IsolatedMargin) != 0 {
			marginMode = "isolated"
		}
		res = append(res, entity.Futures_Positions{
			Symbol:       item.Symbol,
			PositionSide: positionSide,
			PositionSize: item.PositionAmt,
			// Leverage:         item.Lever,
			// PositionID:       item.PosID,
			EntryPrice:       item.EntryPrice,
			MarkPrice:        item.MarkPrice,
			UnRealizedProfit: item.UnRealizedProfit,
			// RealizedProfit:   item.RealizedPnl,
			Notional:   item.Notional,
			HedgeMode:  hedgeMode,
			MarginMode: marginMode,
			// CreateTime:       utils.StringToInt64(item.CTime),
			UpdateTime: item.UpdateTime,
		})
	}
	return res
}

func (c *futures_converts) convertOrderList(answ []futures_orderList) (res []entity.Futures_OrdersList) {
	for _, item := range answ {
		positionSide := "LONG"
		if item.PositionSide == "BOTH" {
			if strings.ToUpper(item.Side) == "SELL" {
				positionSide = "SHORT"
			}
		} else {
			positionSide = strings.ToUpper(item.PositionSide)
		}

		res = append(res, entity.Futures_OrdersList{
			Symbol:        item.Symbol,
			OrderID:       utils.Int64ToString(item.OrderId),
			ClientOrderID: item.ClientOrderId,
			PositionSide:  positionSide,
			Side:          strings.ToUpper(item.Side),
			PositionSize:  item.OrigQty,
			ExecutedSize:  item.ExecutedQty,
			Price:         item.Price,
			Type:          strings.ToUpper(item.Type),
			// MarginMode:    item.TdMode,
			// Leverage: item.Lever,
			Status:     strings.ToUpper(item.Status),
			CreateTime: item.Time,
			UpdateTime: item.UpdateTime,
		})
	}
	return res
}

func (c *futures_converts) convertOrdersHistory(in []futures_ordersHistory_Response) (out []entity.Futures_OrdersHistory) {

	if len(in) == 0 {
		return out
	}

	for _, item := range in {

		hedgeMode := false
		positionSide := "LONG"
		if item.PositionSide == "BOTH" {
			if strings.ToUpper(item.Side) == "SELL" {
				positionSide = "SHORT"
			}
		} else {
			positionSide = strings.ToUpper(item.PositionSide)
			hedgeMode = true
		}

		out = append(out, entity.Futures_OrdersHistory{
			Symbol:        item.Symbol,
			OrderID:       utils.Int64ToString(item.OrderId),
			ClientOrderID: item.ClientOrderId,
			Side:          strings.ToUpper(item.Side),
			PositionSide:  positionSide,
			PositionSize:  item.OrigQty,
			ExecutedSize:  item.ExecutedQty,
			Price:         item.Price,
			ExecutedPrice: item.AvgPrice,
			// RealisedProfit: item.Pnl,
			// Fee:            item.Fee,
			// Leverage:       item.Lever,
			HedgeMode: hedgeMode,
			// MarginMode:     strings.ToLower(item.TdMode),
			Type:       strings.ToUpper(item.Type),
			Status:     strings.ToUpper(item.Status),
			CreateTime: item.Time,
			UpdateTime: item.UpdateTime,
		})
	}
	return out

	// if len(in) == 0 {
	// 	return out
	// }

	// for _, item := range in {
	// 	marginMode := "isolated"
	// 	hedgeMode := false

	// 	if !item.OnlyOnePosition {
	// 		hedgeMode = true
	// 	}

	// 	// if !item.Isolated {
	// 	// 	marginMode = "cross"
	// 	// }

	// 	marginMode = ""

	// 	out = append(out, entity.Futures_OrdersHistory{
	// 		Symbol:         item.Symbol,
	// 		OrderID:        utils.Int64ToString(item.OrderId),
	// 		ClientOrderID:  item.ClientOrderId,
	// 		PositionID:     utils.Int64ToString(item.PositionID),
	// 		Side:           strings.ToUpper(item.Side),
	// 		PositionSide:   strings.ToUpper(item.PositionSide),
	// 		PositionSize:   item.OrigQty,
	// 		ExecutedSize:   item.ExecutedQty,
	// 		Price:          item.Price,
	// 		ExecutedPrice:  item.AvgPrice,
	// 		RealisedProfit: item.Profit,
	// 		Fee:            item.Commission,
	// 		Type:           strings.ToUpper(item.Type),
	// 		Leverage:       strings.Replace(item.Leverage, "X", "", 1),
	// 		Status:         strings.ToUpper(item.Status),
	// 		HedgeMode:      hedgeMode,
	// 		MarginMode:     marginMode,
	// 		CreateTime:     item.Time,
	// 		UpdateTime:     item.UpdateTime,
	// 	})
	// }
	// return out
}

func (c *futures_converts) convertPositionsHistory(in []futures_PositionsHistory_Response) (out []entity.Futures_PositionsHistory) {
	if len(in) == 0 {
		return out
	}

	for _, item := range in {
		// mMode := "cross"
		// if item.MgnMode != "cross" {
		// 	mMode = "isolated"
		// }
		out = append(out, entity.Futures_PositionsHistory{
			Symbol: item.Symbol,
			// PositionID:          item.PosId,
			PositionSide:        strings.ToUpper(item.PositionSide),
			PositionAmt:         item.Qty,
			ExecutedPositionAmt: item.Qty,
			AvgPrice:            item.Price,
			// ExecutedAvgPrice:    item.CloseAvgPx,
			RealisedProfit: item.RealizedPnl,
			Fee:            item.Commission,
			// Funding:             item.FundingFee,
			// MarginMode:          mMode,
			// CreateTime:          utils.StringToInt64(item.CTime),
			UpdateTime: item.Time,
		})
	}
	return out
}
