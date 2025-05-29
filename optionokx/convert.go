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
