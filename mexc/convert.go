package mexc

import (
	"github.com/Betarost/onetrades/entity"
	"github.com/Betarost/onetrades/utils"
)

// ===============SPOT=================

type spot_converts struct{}

func (c *spot_converts) convertInstrumentsInfo(in spot_instrumentsInfo) (out []entity.InstrumentsInfo) {
	if len(in.Symbols) == 0 {
		return out
	}
	for _, item := range in.Symbols {
		state := "OTHER"
		// if item.Status == 1 {
		// 	state = "LIVE"
		// } else if item.Status == 0 {
		// 	state = "OFF"
		// } else if item.Status == 5 {
		// 	state = "PRE-OPEN"
		// } else if item.Status == 25 {
		// 	state = "SUSPENDED"
		// }
		rec := entity.InstrumentsInfo{
			Symbol: item.Symbol,
			// StepTickPrice:    utils.FloatToStringAll(item.TickSize),
			// StepContractSize: utils.FloatToStringAll(item.StepSize),
			// MinContractSize:  utils.FloatToStringAll(item.MinQty),
			State: state,
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
		if item.State == 1 {
			state = "LIVE"
		} else if item.State == 0 {
			state = "OFF"
		} else if item.State == 5 {
			state = "PRE-OPEN"
		} else if item.State == 25 {
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
