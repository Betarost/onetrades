package bingx

import (
	"github.com/Betarost/onetrades/entity"
	"github.com/Betarost/onetrades/utils"
)

// ===============ACCOUNT=================
type account_converts struct{}

func (c *account_converts) convertAccountInfo(in accountInfo) (out entity.AccountInformation) {

	out.UID = in.UID
	// out.Label = in.Label
	// out.IP = in.Ip
	out.PermSpot = true

	// if strings.Contains(in.Perm, "read") {
	// 	out.CanRead = true
	// }

	// if strings.Contains(in.Perm, "trade") {
	// 	out.CanTrade = true
	// }

	// if in.PosMode == "long_short_mode" {
	// 	out.HedgeMode = true
	// }

	// if in.AcctLv != "1" {
	// 	out.PermFutures = true
	// }
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

func (c *spot_converts) convertInstrumentsInfo(in spot_instrumentsInfo) (out []entity.InstrumentsInfo) {
	if len(in.Symbols) == 0 {
		return out
	}
	for _, item := range in.Symbols {
		state := "OTHER"
		if item.Status == 1 {
			state = "LIVE"
		} else if item.Status == 0 {
			state = "OFF"
		} else if item.Status == 5 {
			state = "PRE-OPEN"
		} else if item.Status == 25 {
			state = "SUSPENDED"
		}
		rec := entity.InstrumentsInfo{
			Symbol:           item.Symbol,
			StepTickPrice:    utils.FloatToStringAll(item.TickSize),
			StepContractSize: utils.FloatToStringAll(item.StepSize),
			MinContractSize:  utils.FloatToStringAll(item.MinQty),
			State:            state,
		}
		out = append(out, rec)
	}
	return
}

// ===============FUTURES=================

type futures_converts struct{}

func (c *futures_converts) convertInstrumentsInfo(in []futures_instrumentsInfo) (out []entity.InstrumentsInfo) {
	if len(in) == 0 {
		return out
	}
	for _, item := range in {
		state := "OTHER"
		if item.Status == 1 {
			state = "LIVE"
		} else if item.Status == 0 {
			state = "OFF"
		} else if item.Status == 5 {
			state = "PRE-OPEN"
		} else if item.Status == 25 {
			state = "SUSPENDED"
		}
		rec := entity.InstrumentsInfo{
			Symbol:          item.Symbol,
			MinContractSize: utils.FloatToStringAll(item.TradeMinQuantity),
			State:           state,
		}
		out = append(out, rec)
	}
	return
}
