package bitget

import (
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
