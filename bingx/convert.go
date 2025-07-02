package bingx

import (
	"fmt"
	"strings"
	"time"

	"github.com/Betarost/onetrades/entity"
	"github.com/Betarost/onetrades/utils"
)

// ===============ACCOUNT=================
type account_converts struct{}

func (c *account_converts) convertAccountInfo(in accountInfo) (out entity.AccountInformation) {
	out.Label = in.Note
	out.IP = strings.Join(in.IpAddresses, ",")
	for _, item := range in.Permissions {
		switch item {
		case 1:
			out.CanTrade = true
			out.PermSpot = true
		case 2:
			out.CanRead = true
		case 3:
			out.PermFutures = true
		}
	}
	return out
}

func (c *account_converts) convertFundingAccountBalance(in fundingBalance) (out []entity.FundingAccountBalance) {
	if len(in.Assets) == 0 {
		return out
	}

	for _, item := range in.Assets {
		out = append(out, entity.FundingAccountBalance{
			Asset:            item.Asset,
			Balance:          utils.FloatToStringAll(item.Free),
			AvailableBalance: utils.FloatToStringAll(item.Free),
			FrozenBalance:    utils.FloatToStringAll(item.Locked),
		})
	}

	return out
}

func (c *account_converts) convertTradingAccountBalance(in []tradingBalance) (out []entity.TradingAccountBalance) {
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

// ===============SPOT=================

type spot_converts struct{}

func (c *spot_converts) convertInstrumentsInfo(in spot_instrumentsInfo) (out []entity.Spot_InstrumentsInfo) {
	if len(in.Symbols) == 0 {
		return out
	}
	for _, item := range in.Symbols {
		state := "OTHER"
		base := ""
		quote := ""

		sp := strings.Split(item.Symbol, "-")
		if len(sp) == 2 {
			base = sp[0]
			quote = sp[1]
		}

		priceP := utils.GetPrecisionFromStr(utils.FloatToStringAll(item.TickSize))
		sizeP := utils.GetPrecisionFromStr(utils.FloatToStringAll(item.StepSize))
		if item.Status == 1 {
			state = "LIVE"
		} else if item.Status == 0 {
			state = "OFF"
		} else if item.Status == 5 {
			state = "PRE-OPEN"
		} else if item.Status == 25 {
			state = "SUSPENDED"
		}
		rec := entity.Spot_InstrumentsInfo{
			Symbol:         item.Symbol,
			Base:           base,
			Quote:          quote,
			MinQty:         utils.FloatToStringAll(item.MinQty),
			MinNotional:    utils.FloatToStringAll(item.MinNotional),
			PricePrecision: priceP,
			SizePrecision:  sizeP,
			State:          state,
		}
		out = append(out, rec)
	}
	return
}

func (c *spot_converts) convertBalance(in spot_Balance) (out []entity.AssetsBalance) {
	for _, item := range in.Balances {
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
		OrderID:       fmt.Sprintf("%d", in.OrderId),
		ClientOrderID: in.ClientOrderID,
		Ts:            time.Now().UTC().UnixMilli(),
	})
	return out
}

func (c *spot_converts) convertOrderList(in spot_orderList) (out []entity.Spot_OrdersList) {
	if len(in.Orders) == 0 {
		return out
	}
	for _, item := range in.Orders {
		out = append(out, entity.Spot_OrdersList{
			Symbol:        item.Symbol,
			OrderID:       fmt.Sprintf("%d", item.OrderId),
			ClientOrderID: item.ClientOrderId,
			Side:          item.Side,
			Size:          item.OrigQty,
			Price:         item.Price,
			ExecutedSize:  item.ExecutedQty,
			Type:          strings.ToUpper(item.Type),
			Status:        item.Status,
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
		state := "OTHER"
		switch item.Status {
		case 1:
			state = "LIVE"
		case 0:
			state = "OFF"
		case 5:
			state = "PRE-OPEN"
		case 25:
			state = "SUSPENDED"
		}
		out = append(out, entity.Futures_InstrumentsInfo{
			Symbol:         item.Symbol,
			Base:           item.Asset,
			Quote:          item.Currency,
			MinQty:         utils.FloatToStringAll(item.TradeMinQuantity),
			MinNotional:    utils.FloatToStringAll(item.TradeMinUSDT),
			PricePrecision: fmt.Sprintf("%d", item.PricePrecision),
			SizePrecision:  fmt.Sprintf("%d", item.QuantityPrecision),
			State:          state,
		})
	}
	return
}

func (c *futures_converts) convertBalance(in []futures_Balance) (out []entity.FuturesBalance) {
	if len(in) == 0 {
		return out
	}

	for _, item := range in {
		out = append(out, entity.FuturesBalance{
			Asset:     item.Asset,
			Balance:   item.Balance,
			Equity:    item.Equity,
			Available: item.AvailableMargin,
			// AvailableMargin:  item.AvailableMargin,
			UnrealizedProfit: item.UnrealizedProfit,
		})
	}
	return out
}

func (c *futures_converts) convertLeverage(in futures_leverage) (out entity.Futures_Leverage) {

	out.Symbol = in.Symbol
	if in.LongLeverage != 0 {
		out.Leverage = fmt.Sprintf("%d", in.LongLeverage)
	}

	if in.Leverage != 0 {
		out.Leverage = fmt.Sprintf("%d", in.Leverage)
	}
	return out
}

func (c *futures_converts) convertPlaceOrder(in futures_placeOrder_Response) (out []entity.PlaceOrder) {
	out = append(out, entity.PlaceOrder{
		OrderID:       in.Order.OrderID,
		ClientOrderID: in.Order.ClientOrderId,
		Ts:            time.Now().UTC().UnixMilli(),
	})
	return out
}
