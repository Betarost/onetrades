package optionokx

import (
	"strings"

	"github.com/Betarost/onetrades/entity"
	"github.com/Betarost/onetrades/utils"
)

func ConvertContractsInfo(answ []ContractsInfo) (res []entity.ContractInfo_Option) {
	if len(answ) == 0 {
		return res
	}
	for _, item := range answ {
		res = append(res, entity.ContractInfo_Option{
			Symbol:             item.InstId,
			ContractSize:       utils.StringToFloat(item.CtVal),
			ContractMultiplier: utils.StringToFloat(item.CtMult),
			StepContractSize:   utils.StringToFloat(item.LotSz),
			MinContractSize:    utils.StringToFloat(item.MinSz),
			StepTickPrice:      utils.StringToFloat(item.TickSz),
			Strike:             utils.StringToFloat(item.Stk),
			ListTime:           utils.StringToInt64(item.ListTime),
			ExpTime:            utils.StringToInt64(item.ExpTime),
			State:              strings.ToUpper(item.State),
			Type:               strings.ToUpper(item.OptType),
		})
	}
	return res
}

func ConvertMarketData(answ []MarketData) (res []entity.MarketData_Option) {
	if len(answ) == 0 {
		return res
	}
	for _, item := range answ {
		sp := strings.Split(item.InstId, "-")
		tCP := ""
		str := 0.0
		if len(sp) == 5 {
			str = utils.StringToFloat(sp[3])
			tCP = sp[4]
		}
		res = append(res, entity.MarketData_Option{
			Symbol:  item.InstId,
			Delta:   utils.StringToFloat(item.Delta),
			Gamma:   utils.StringToFloat(item.Gamma),
			Vega:    utils.StringToFloat(item.Vega),
			Theta:   utils.StringToFloat(item.Theta),
			DeltaBS: utils.StringToFloat(item.DeltaBS),
			GammaBS: utils.StringToFloat(item.GammaBS),
			VegaBS:  utils.StringToFloat(item.VegaBS),
			ThetaBS: utils.StringToFloat(item.ThetaBS),

			MarkVol:    utils.StringToFloat(item.MarkVol),
			BidVol:     utils.StringToFloat(item.BidVol),
			AskVol:     utils.StringToFloat(item.AskVol),
			RealVol:    utils.StringToFloat(item.RealVol),
			VolLv:      utils.StringToFloat(item.VolLv),
			FwdPx:      utils.StringToFloat(item.FwdPx),
			Strike:     str,
			Type:       tCP,
			Leverage:   utils.StringToInt(item.Lever),
			UpdateTime: utils.StringToInt64(item.Ts),
		})
	}
	return res
}

func ConvertOrderBook(answ []OrderBook) (res entity.OrderBook_Option) {
	if len(answ) == 0 {
		return res
	}

	book := answ[0]
	res.UpdateTime = utils.StringToInt64(book.Ts)
	for _, item := range book.Asks {
		if len(item) != 4 {
			continue
		}
		res.Asks = append(res.Asks, entity.AsksBids_Option{
			Price:     utils.StringToFloat(item[0]),
			Qty:       utils.StringToFloat(item[1]),
			NumOrders: utils.StringToFloat(item[3]),
		})
	}

	for _, item := range book.Bids {
		if len(item) != 4 {
			continue
		}
		res.Bids = append(res.Asks, entity.AsksBids_Option{
			Price:     utils.StringToFloat(item[0]),
			Qty:       utils.StringToFloat(item[1]),
			NumOrders: utils.StringToFloat(item[3]),
		})
	}

	return res
}

func ConvertPositions(answ []Position) (res []entity.Position) {
	for _, item := range answ {
		if item.InstType != "OPTION" {
			continue
		}
		positionSide := "LONG"
		if item.PosSide == "net" {
			if utils.StringToFloat(item.Pos) < 0 {
				positionSide = "SHORT"
			}
		} else {
			positionSide = strings.ToUpper(item.PosSide)
		}

		res = append(res, entity.Position{
			Symbol:           item.InstID,
			PositionSide:     positionSide,
			PositionAmt:      utils.StringToFloat(item.Pos),
			EntryPrice:       utils.StringToFloat(item.AvgPx),
			MarkPrice:        utils.StringToFloat(item.MarkPx),
			InitialMargin:    utils.StringToFloat(item.Imr),
			UnRealizedProfit: utils.StringToFloat(item.Upl),
			RealizedProfit:   utils.StringToFloat(item.RealizedPnl),
			Notional:         utils.StringToFloat(item.NotionalUsd),
			UpdateTime:       utils.StringToInt64(item.UTime),

			MarginRatio:      item.MgnRatio,
			AutoDeleveraging: item.ADL,
		})
	}
	return res
}
