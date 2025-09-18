package kucoin

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
	out.Label = in.Remark

	if strings.Contains(in.Permission, "General") {
		out.CanRead = true
	}

	if strings.Contains(in.Permission, "Spot") {
		out.PermSpot = true
		out.CanTrade = true
	}

	if strings.Contains(in.Permission, "Futures") {
		out.PermFutures = true
		out.CanTrade = true
	}

	if strings.Contains(in.Permission, "Withdrawal") {
		out.CanTransfer = true
	}

	return out
}

// ===============SPOT=================

type spot_converts struct{}

func (c *spot_converts) convertBalance(in []spot_Balance) (out []entity.AssetsBalance) {
	if len(in) == 0 {
		return out
	}

	for _, i := range in {
		if i.Type == "main" || (i.Balance == "" || i.Balance == "0") {
			continue
		}
		out = append(out, entity.AssetsBalance{
			Asset:   i.Currency,
			Balance: i.Balance,
			Locked:  i.Holds,
		})
	}
	return out
}

func (c *spot_converts) convertInstrumentsInfo(in []spot_instrumentsInfo) (out []entity.Spot_InstrumentsInfo) {
	if len(in) == 0 {
		return out
	}
	for _, item := range in {

		priceP := utils.GetPrecisionFromStr(item.PriceIncrement)
		sizeP := utils.GetPrecisionFromStr(item.BaseIncrement)
		state := "OFF"
		if item.EnableTrading {
			state = "LIVE"
		}
		out = append(out, entity.Spot_InstrumentsInfo{
			Symbol:         item.Symbol,
			Base:           item.BaseCurrency,
			Quote:          item.QuoteCurrency,
			MinQty:         item.BaseMinSize,
			MinNotional:    item.QuoteMinSize,
			PricePrecision: priceP,
			SizePrecision:  sizeP,
			State:          strings.ToUpper(state),
		})
	}
	return out
}

func (c *spot_converts) convertPlaceOrder(in placeOrder_Response) (out []entity.PlaceOrder) {

	oID := in.OrderId
	if oID == "" {
		oID = in.CancelledOrderIds
	}
	out = append(out, entity.PlaceOrder{
		OrderID: oID,
		Ts:      time.Now().UTC().UnixMilli(),
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
			OrderID:       item.ID,
			ClientOrderID: item.ClientOid,
			Size:          item.Size,
			ExecutedSize:  item.DealSize,
			Side:          strings.ToUpper(item.Side),
			Price:         item.Price,
			Type:          strings.ToUpper(item.Type),
			Status:        strings.ToUpper(item.TradeType),
			CreateTime:    item.CreatedAt,
			UpdateTime:    time.Now().UTC().UnixMilli(),
		})
	}
	return out
}

func (c *spot_converts) convertOrdersHistory(in []spot_ordersHistory_Response) (out []entity.Spot_OrdersHistory) {
	if len(in) == 0 {
		return out
	}

	for _, item := range in {
		out = append(out, entity.Spot_OrdersHistory{
			Symbol:  item.Symbol,
			OrderID: item.OrderId,
			// ClientOrderID: item.ClOrdId,
			Side:          strings.ToUpper(item.Side),
			Size:          item.Size,
			Price:         item.Price,
			ExecutedSize:  item.Size,
			ExecutedPrice: item.Price,
			Fee:           item.Fee,
			Type:          strings.ToUpper(item.Type),
			// Status:        strings.ToUpper(item.TradeType),
			Status:     "FILLED",
			CreateTime: item.CreatedAt,
			UpdateTime: item.CreatedAt,
		})
	}
	return out
}

func (c *spot_converts) convertListenKey(in spot_listenKey) (out entity.Spot_ListenKey) {
	out.ListenKey = in.Token
	return out
}

// ===============FUTURES=================

type futures_converts struct{}

func (c *futures_converts) convertBalance(in []futures_Balance) (out []entity.FuturesBalance) {
	if len(in) == 0 {
		return out
	}

	for _, i := range in {
		for _, item := range i.Details {
			out = append(out, entity.FuturesBalance{
				Asset:            item.Ccy,
				Balance:          item.CashBal,
				Equity:           item.Eq,
				Available:        item.AvailBal,
				UnrealizedProfit: item.Upl,
			})
		}
	}
	return out
}
