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

func (c *futures_converts) convertInstrumentsInfo(in []futures_instrumentsInfo) (out []entity.InstrumentsInfo) {
	if len(in) == 0 {
		return out
	}
	for _, item := range in {
		state := "OTHER"
		// if item.Status == 1 {
		// 	state = "LIVE"
		// } else if item.Status == 0 {
		// 	state = "OFF"
		// } else if item.Status == 5 {
		// 	state = "PRE-OPEN"
		// } else if item.Status == 25 {
		// 	state = "SUSPENDED"
		// }
		rec := entity.InstrumentsInfo{
			Symbol: item.Symbol,
			// MinContractSize: utils.FloatToStringAll(item.TradeMinQuantity),
			State: state,
		}
		out = append(out, rec)
	}
	return
}
