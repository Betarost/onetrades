package binance

import (
	"github.com/Betarost/onetrades/entity"
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

// ===============SPOT=================

type spot_converts struct{}

func (c *spot_converts) convertInstrumentsInfo(in spot_instrumentsInfo) (out []entity.InstrumentsInfo) {
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
