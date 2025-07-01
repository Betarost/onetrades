package bybit

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

	out.UID = fmt.Sprintf("%d", in.UserID)
	out.Label = in.Note
	out.IP = strings.Join(in.Ips, ",")
	out.CanRead = true

	if in.ReadOnly == 0 {
		out.CanTrade = true
	}

	for _, item := range in.Permissions.Spot {
		if item == "SpotTrade" {
			out.PermSpot = true
			break
		}
	}

	for _, item := range in.Permissions.Derivatives {
		if item == "DerivativesTrade" {
			out.PermFutures = true
			break
		}
	}

	for _, item := range in.Permissions.Wallet {
		if item == "AccountTransfer" {
			out.CanTransfer = true
			break
		}
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
		for _, item := range i.Coin {
			out = append(out, entity.AssetsBalance{
				Asset:   item.Coin,
				Balance: item.WalletBalance,
				Locked:  item.Locked,
			})
		}
	}
	return out
}

func (c *spot_converts) convertInstrumentsInfo(in []spot_instrumentsInfo) (out []entity.Spot_InstrumentsInfo) {
	if len(in) == 0 {
		return out
	}
	for _, item := range in {
		if item.Status == "Trading" {
			item.Status = "LIVE"
		}

		sizeP := utils.GetPrecisionFromStr(item.LotSizeFilter.BasePrecision)
		priceP := utils.GetPrecisionFromStr(item.PriceFilter.TickSize)

		rec := entity.Spot_InstrumentsInfo{
			Symbol:         item.Symbol,
			Base:           item.BaseCoin,
			Quote:          item.QuoteCoin,
			MinQty:         item.LotSizeFilter.MinOrderQty,
			MinNotional:    item.LotSizeFilter.MinOrderAmt,
			PricePrecision: priceP,
			SizePrecision:  sizeP,
			State:          item.Status,
		}
		out = append(out, rec)
	}
	return out
}

func (c *spot_converts) convertPlaceOrder(in placeOrder_Response) (out []entity.PlaceOrder) {

	out = append(out, entity.PlaceOrder{
		OrderID:       in.OrderId,
		ClientOrderID: in.OrderLinkId,
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
			ClientOrderID: item.OrderLinkId,
			Side:          strings.ToUpper(item.Side),
			Size:          item.Qty,
			Price:         item.Price,
			ExecutedSize:  item.CumExecQty,
			Type:          strings.ToUpper(item.OrderType),
			Status:        item.OrderStatus,
			CreateTime:    utils.StringToInt64(item.CreatedTime),
			UpdateTime:    utils.StringToInt64(item.UpdatedTime),
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
		if item.Status == "Trading" {
			item.Status = "LIVE"
		}

		rec := entity.Futures_InstrumentsInfo{
			Symbol:         item.Symbol,
			Base:           item.BaseCoin,
			Quote:          item.QuoteCoin,
			MinQty:         item.LotSizeFilter.MinOrderQty,
			MinNotional:    item.LotSizeFilter.MinNotionalValue,
			PricePrecision: item.PriceScale,
			SizePrecision:  utils.GetPrecisionFromStr(item.LotSizeFilter.MinOrderQty),
			MaxLeverage:    item.LeverageFilter.MaxLeverage,
			State:          item.Status,
		}
		out = append(out, rec)
	}
	return
}

func (c *futures_converts) convertBalance(in []futures_Balance) (out []entity.FuturesBalance) {
	if len(in) == 0 {
		return out
	}
	for _, i := range in {
		for _, item := range i.Coin {
			out = append(out, entity.FuturesBalance{
				Asset:   item.Coin,
				Balance: item.WalletBalance,
				Equity:  item.Equity,
				// AvailableMargin:  item.AvailableToWithdraw,
				UnrealizedProfit: item.UnrealisedPnl,
			})
		}
	}
	return out
}

func (c *futures_converts) convertLeverage(in futures_leverage) (out entity.Futures_Leverage) {

	out.Symbol = in.Symbol
	out.Leverage = fmt.Sprintf("%d", in.LongLeverage)
	return out
}

// ======================OLD
func convertTradingAccountBalance(in []tradingBalance) (out []entity.TradingAccountBalance) {
	if len(in) == 0 {
		return out
	}

	for _, item := range in {
		r := entity.TradingAccountBalance{
			TotalEquity:      item.TotalEquity,
			AvailableEquity:  item.TotalAvailableBalance,
			UnrealizedProfit: item.TotalPerpUPL,
			UpdateTime:       time.Now().UnixMilli(),
		}
		for _, i := range item.Coin {
			r.Assets = append(r.Assets, entity.TradingAccountBalanceDetails{
				Asset:            i.Coin,
				Balance:          i.WalletBalance,
				EquityBalance:    i.Equity,
				AvailableBalance: i.AvailableToWithdraw,
				AvailableEquity:  i.AvailableToWithdraw,
				UnrealizedProfit: i.UnrealisedPnl,
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
			Asset:            item.Coin,
			Balance:          item.WalletBalance,
			AvailableBalance: item.TransferBalance,
		})
	}

	return out
}
