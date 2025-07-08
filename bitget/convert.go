package bitget

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/Betarost/onetrades/entity"
	"github.com/Betarost/onetrades/utils"
)

// ===============ACCOUNT=================
type account_converts struct{}

func (c *account_converts) convertAccountInfo(in accountInfo) (out entity.AccountInformation) {

	out.UID = in.UserId
	out.IP = in.Ips

	for _, item := range in.Authorities {
		switch item {
		case "coor":
			out.PermFutures = true
			out.CanRead = true
		case "stor":
			out.PermSpot = true
			out.CanRead = true
		case "coow":
			out.PermFutures = true
			out.CanTrade = true
			out.CanRead = true
		case "stow":
			out.PermSpot = true
			out.CanTrade = true
			out.CanRead = true
		}
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

		if item.Status == "online" {
			item.Status = "LIVE"
		}

		rec := entity.Spot_InstrumentsInfo{
			Symbol: item.Symbol,
			Base:   item.BaseCoin,
			Quote:  item.QuoteCoin,
			// MinQty:         utils.FloatToStringAll(item.MinQty),
			MinNotional:    item.MinTradeUSDT,
			PricePrecision: item.PricePrecision,
			SizePrecision:  item.QuantityPrecision,
			State:          item.Status,
		}
		out = append(out, rec)
	}
	return
}

func (c *spot_converts) convertBalance(in []spot_Balance) (out []entity.AssetsBalance) {
	if len(in) == 0 {
		return out
	}
	for _, item := range in {
		out = append(out, entity.AssetsBalance{
			Asset:   item.Coin,
			Balance: item.Available,
			Locked:  item.Locked,
		})
	}
	return out
}

func (c *spot_converts) convertPlaceOrder(in placeOrder_Response) (out []entity.PlaceOrder) {

	out = append(out, entity.PlaceOrder{
		OrderID:       in.OrderId,
		ClientOrderID: in.ClientOid,
		Ts:            time.Now().UTC().UnixMilli(),
	})
	return out
}

func (c *spot_converts) convertOrdersHistory(in []spot_ordersHistory_Response) (out []entity.Spot_OrdersHistory) {
	if len(in) == 0 {
		return out
	}

	type feeDetailJson struct {
		NewFees struct {
			T float64 `json:"t"`
		} `json:"newFees"`
	}

	for _, item := range in {
		executedQty := item.BaseVolume
		if strings.ToUpper(item.OrderType) == "MARKET" {
			executedQty = item.QuoteVolume
		}
		fee := "0"
		answ := feeDetailJson{}
		err := json.Unmarshal([]byte(item.FeeDetail), &answ)
		if err == nil {
			fee = utils.FloatToStringAll(answ.NewFees.T)
		} else {
			log.Println("=Err convertOrdersHistory=", err)
		}
		out = append(out, entity.Spot_OrdersHistory{
			Symbol:        item.Symbol,
			OrderID:       item.OrderId,
			ClientOrderID: item.ClientOid,
			Side:          strings.ToUpper(item.Side),
			Size:          item.Size,
			Price:         item.Price,
			ExecutedSize:  executedQty,
			ExecutedPrice: item.PriceAvg,
			Fee:           fee,
			Type:          strings.ToUpper(item.OrderType),
			Status:        strings.ToUpper(item.Status),
			CreateTime:    utils.StringToInt64(item.CTime),
			UpdateTime:    utils.StringToInt64(item.UTime),
		})
	}
	return out
}

func (c *spot_converts) convertOrderList(in []spot_orderList) (out []entity.Spot_OrdersList) {
	if len(in) == 0 {
		return out
	}
	for _, item := range in {
		out = append(out, entity.Spot_OrdersList{
			Symbol:        item.Symbol,
			OrderID:       item.OrderId,
			ClientOrderID: item.ClientOid,
			Side:          strings.ToUpper(item.Side),
			Size:          item.Size,
			Price:         item.PriceAvg,
			ExecutedSize:  item.BaseVolume,
			Type:          strings.ToUpper(item.OrderType),
			Status:        item.Status,
			CreateTime:    utils.StringToInt64(item.CTime),
			UpdateTime:    utils.StringToInt64(item.UTime),
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
		if item.SymbolStatus == "normal" {
			item.SymbolStatus = "LIVE"
		}

		rec := entity.Futures_InstrumentsInfo{
			Symbol:         item.Symbol,
			Base:           item.BaseCoin,
			Quote:          item.QuoteCoin,
			MinQty:         item.MinTradeNum,
			MinNotional:    item.MinTradeUSDT,
			PricePrecision: item.PricePlace,
			SizePrecision:  item.VolumePlace,
			MaxLeverage:    item.MaxLever,
			Multiplier:     item.SizeMultiplier,
			State:          item.SymbolStatus,
		}
		out = append(out, rec)
	}
	return
}

func (c *futures_converts) convertBalance(in []futures_Balance) (out []entity.FuturesBalance) {
	if len(in) == 0 {
		return out
	}

	for _, item := range in {
		out = append(out, entity.FuturesBalance{
			Asset:            item.MarginCoin,
			Balance:          item.AccountEquity,
			Equity:           item.AccountEquity,
			Available:        item.Available,
			UnrealizedProfit: item.UnrealizedPL,
		})

		// if len(i.AssetList) == 0 {
		// 	continue
		// }
		// for _, item := range i.AssetList {
		// out = append(out, entity.FuturesBalance{
		// 	Asset:   item.Coin,
		// 	Balance: item.Balance,
		// 	// Equity:           item.AccountEquity,
		// 	Available:        item.Available,
		// 	UnrealizedProfit: i.UnrealizedPL,
		// })
		// }
	}
	return out
}

func (c *futures_converts) convertLeverage(in futures_leverage) (out entity.Futures_Leverage) {

	out.Symbol = in.Symbol
	out.Leverage = fmt.Sprintf("%d", in.CrossedMarginLeverage)
	return out
}

func (c *futures_converts) convertPositionsHistory(in []futures_PositionsHistory_Response) (out []entity.Futures_PositionsHistory) {
	if len(in) == 0 {
		return out
	}

	for _, item := range in {
		mMode := "cross"
		if item.MarginMode == "isolated" {
			mMode = "isolated"
		}

		fee := utils.StringToFloat(item.OpenFee) + utils.StringToFloat(item.CloseFee)

		out = append(out, entity.Futures_PositionsHistory{
			Symbol:              item.Symbol,
			PositionID:          item.PositionId,
			PositionSide:        strings.ToUpper(item.HoldSide),
			PositionAmt:         item.OpenTotalPos,
			ExecutedPositionAmt: item.CloseTotalPos,
			AvgPrice:            item.OpenAvgPrice,
			ExecutedAvgPrice:    item.CloseAvgPrice,
			RealisedProfit:      item.Pnl,
			Fee:                 utils.FloatToStringAll(fee),
			Funding:             item.TotalFunding,
			MarginMode:          mMode,
			CreateTime:          utils.StringToInt64(item.CTime),
			UpdateTime:          utils.StringToInt64(item.UTime),
		})
	}
	return out
}
