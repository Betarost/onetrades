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
			Symbol:        item.Symbol,
			Base:          item.BaseAsset,
			Quote:         item.QuoteAsset,
			SizePrecision: utils.Int64ToString(item.BaseAssetPrecision),
			State:         item.Status,
		}
		for _, i := range item.Filters {
			m := i.(map[string]interface{})
			switch m["filterType"] {
			case "PRICE_FILTER":
				rec.PricePrecision = utils.GetPrecisionFromStr(utils.FloatToStringAll(utils.StringToFloat(m["tickSize"].(string))))
			case "LOT_SIZE":
				rec.MinQty = utils.FloatToStringAll(utils.StringToFloat(m["minQty"].(string)))
				// rec.MinQty = m["minQty"].(string)
				// rec.StepContractSize = m["stepSize"].(string)
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

func (c *futures_converts) convertInstrumentsInfo(in futures_instrumentsInfo) (out []entity.InstrumentsInfo) {
	if len(in.Symbols) == 0 {
		return out
	}
	for _, item := range in.Symbols {
		if item.Status == "TRADING" {
			item.Status = "LIVE"
		}
		rec := entity.InstrumentsInfo{
			Symbol: item.Symbol,
			Base:   item.BaseAsset,
			Quote:  item.QuoteAsset,
			State:  item.Status,
		}
		for _, i := range item.Filters {
			m := i.(map[string]interface{})
			switch m["filterType"] {
			case "PRICE_FILTER":
				rec.StepTickPrice = m["tickSize"].(string)
			case "LOT_SIZE":
				rec.MinContractSize = m["minQty"].(string)
				rec.StepContractSize = m["stepSize"].(string)
			}
		}
		out = append(out, rec)
	}
	return
}
