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
			// out.CanTransfer = true
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

// ===============FUTURES=================

type futures_converts struct{}

func (c *futures_converts) convertInstrumentsInfo(in []futures_instrumentsInfo) (out []entity.InstrumentsInfo) {
	// if len(in) == 0 {
	// 	return out
	// }
	// for _, item := range in {
	// 	state := "OTHER"
	// 	if item.Status == 1 {
	// 		state = "LIVE"
	// 	} else if item.Status == 0 {
	// 		state = "OFF"
	// 	} else if item.Status == 5 {
	// 		state = "PRE-OPEN"
	// 	} else if item.Status == 25 {
	// 		state = "SUSPENDED"
	// 	}
	// 	rec := entity.InstrumentsInfo{
	// 		Symbol:          item.Symbol,
	// 		MinContractSize: utils.FloatToStringAll(item.TradeMinQuantity),
	// 		State:           state,
	// 	}
	// 	out = append(out, rec)
	// }
	return
}

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

func convertInstrumentsInfo(in []spot_instrumentsInfo, instType string) (out []entity.InstrumentsInfo) {
	if len(in) == 0 {
		return out
	}

	for _, item := range in {
		if item.Status == "Trading" {
			item.Status = "LIVE"
		}
		out = append(out, entity.InstrumentsInfo{
			Symbol:           item.Symbol,
			StepContractSize: item.LotSizeFilter.BasePrecision,
			MinContractSize:  item.LotSizeFilter.MinOrderQty,
			StepTickPrice:    item.PriceFilter.TickSize,
			State:            item.Status,
			// InstType:         instType,
			Base:  item.BaseCoin,
			Quote: item.QuoteCoin,
		})
	}
	return out
}

func futures_convertInstrumentsInfo(in []futures_instrumentsInfo, instType string) (out []entity.InstrumentsInfo) {
	if len(in) == 0 {
		return out
	}

	for _, item := range in {
		if item.Status == "Trading" {
			item.Status = "LIVE"
		}
		out = append(out, entity.InstrumentsInfo{
			Symbol:           item.Symbol,
			StepContractSize: item.LotSizeFilter.BasePrecision,
			MinContractSize:  item.LotSizeFilter.MinOrderQty,
			StepTickPrice:    item.PriceFilter.TickSize,
			State:            item.Status,
			// InstType:         instType,
			Base:  item.BaseCoin,
			Quote: item.QuoteCoin,
		})
	}
	return out
}
