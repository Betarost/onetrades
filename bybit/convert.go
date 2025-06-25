package bybit

import (
	"fmt"
	"strings"
	"time"

	"github.com/Betarost/onetrades/entity"
)

// ===============SPOT=================

type spot_converts struct{}

func (c *spot_converts) convertInstrumentsInfo(in spot_instrumentsInfo) (out []entity.InstrumentsInfo) {
	// if len(in.Symbols) == 0 {
	// 	return out
	// }
	// for _, item := range in.Symbols {
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
	// 		Symbol:           item.Symbol,
	// 		StepTickPrice:    utils.FloatToStringAll(item.TickSize),
	// 		StepContractSize: utils.FloatToStringAll(item.StepSize),
	// 		MinContractSize:  utils.FloatToStringAll(item.MinQty),
	// 		State:            state,
	// 	}
	// 	out = append(out, rec)
	// }
	return
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

func convertAccountInfo(in accountInfo) (out entity.AccountInformation) {

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
func convertPlaceOrder(in spot_placeOrder_Response) (out []entity.PlaceOrder) {
	// if len(in) == 0 {
	// 	return out
	// }
	// for _, item := range in {
	// 	out = append(out, entity.PlaceOrder{
	// 		OrderID:       item.OrdId,
	// 		ClientOrderID: item.ClOrdId,
	// 		Ts:            utils.StringToInt64(item.Ts),
	// 	})
	// }
	return out
}
