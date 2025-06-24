package gateio

import (
	"strings"

	"github.com/Betarost/onetrades/entity"
)

// ===============ACCOUNT=================
type account_converts struct{}

func (c *account_converts) convertAccountInfo(in accountInfo) (out entity.AccountInformation) {

	out.UID = in.User_id
	out.IP = strings.Join(in.Ip_whitelist, ",")
	// out.PermSpot = true

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

func (c *spot_converts) convertInstrumentsInfo(in []spot_instrumentsInfo) (out []entity.InstrumentsInfo) {
	if len(in) == 0 {
		return out
	}
	for _, item := range in {
		if item.Trade_status == "tradable" {
			item.Trade_status = "LIVE"
		}
		rec := entity.InstrumentsInfo{
			Symbol: item.ID,
			Base:   item.Base,
			State:  item.Trade_status,
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
		if !item.In_delisting {
			state = "LIVE"
		}
		rec := entity.InstrumentsInfo{
			Symbol: item.Name,
			// MinContractSize: utils.FloatToStringAll(item.TradeMinQuantity),
			State: state,
		}
		out = append(out, rec)
	}
	return
}
