package bybit

import (
	"fmt"
	"strings"
	"time"

	"github.com/Betarost/onetrades/entity"
)

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

func convertInstrumentsInfo(in []instrumentsInfo, instType string) (out []entity.InstrumentsInfo) {
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
			InstType:         instType,
			Base:             item.BaseCoin,
			Quote:            item.QuoteCoin,
		})
	}
	return out
}
