package gateio

import (
	"github.com/Betarost/onetrades/entity"
)

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
