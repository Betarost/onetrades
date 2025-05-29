package optionokx

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Betarost/onetrades/entity"
	"github.com/Betarost/onetrades/utils"
)

// ==============GetContractsInfo=================
type GetContractsInfo struct {
	c *Client
	// symbol *string
}

func (s *GetContractsInfo) Do(ctx context.Context, opts ...utils.RequestOption) (res []entity.ContractInfo_Option, err error) {
	r := &utils.Request{
		Method:     http.MethodGet,
		BaseURL:    s.c.BaseURL,
		Endpoint:   "/api/v5/public/instruments",
		TimeOffset: s.c.TimeOffset,
		SecType:    utils.SecTypeNone,
	}

	m := utils.Params{
		"instType": "OPTION",
		"uly":      "BTC-USD",
	}

	r.SetParams(m)

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}
	var answ struct {
		Result []ContractsInfo `json:"data"`
	}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}
	return ConvertContractsInfo(answ.Result), nil
}

type ContractsInfo struct {
	InstId   string `json:"instId"`
	InstType string `json:"instType"`
	CtVal    string `json:"ctVal"`
	CtMult   string `json:"ctMult"`
	CtValCcy string `json:"ctValCcy"`
	TickSz   string `json:"tickSz"`
	LotSz    string `json:"lotSz"`
	MinSz    string `json:"minSz"`
	Lever    string `json:"lever"`
	Stk      string `json:"stk"`
	State    string `json:"state"`
	RuleType string `json:"ruleType"`
	OptType  string `json:"optType"`
	ExpTime  string `json:"expTime"`
	ListTime string `json:"listTime"`
}

// ==============GetMarketData=================
type GetMarketData struct {
	c       *Client
	expTime *string
}

func (s *GetMarketData) ExpTime(expTime string) *GetMarketData {
	s.expTime = &expTime
	return s
}

func (s *GetMarketData) Do(ctx context.Context, opts ...utils.RequestOption) (res []entity.MarketData_Option, err error) {
	r := &utils.Request{
		Method:     http.MethodGet,
		BaseURL:    s.c.BaseURL,
		Endpoint:   "/api/v5/public/opt-summary",
		TimeOffset: s.c.TimeOffset,
		SecType:    utils.SecTypeNone,
	}

	m := utils.Params{
		"uly": "BTC-USD",
	}

	if s.expTime != nil {
		m["expTime"] = *s.expTime
	}

	r.SetParams(m)

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}

	var answ struct {
		Result []MarketData `json:"data"`
	}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}

	if len(answ.Result) == 0 {
		return res, errors.New("Zero Answer")
	}

	return ConvertMarketData(answ.Result), nil
}

type MarketData struct {
	InstId   string `json:"instId"`
	InstType string `json:"instType"`
	Delta    string `json:"delta"`
	Gamma    string `json:"gamma"`
	Vega     string `json:"vega"`
	Theta    string `json:"theta"`
	DeltaBS  string `json:"deltaBS"`
	GammaBS  string `json:"gammaBS"`
	VegaBS   string `json:"vegaBS"`
	ThetaBS  string `json:"thetaBS"`
	Lever    string `json:"lever"`
	MarkVol  string `json:"markVol"`
	BidVol   string `json:"bidVol"`
	AskVol   string `json:"askVol"`
	RealVol  string `json:"realVol"`
	VolLv    string `json:"volLv"`
	FwdPx    string `json:"fwdPx"`
	Ts       string `json:"ts"`
}
