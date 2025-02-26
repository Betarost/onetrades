package futureokx

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/Betarost/onetrades/entity"
	"github.com/Betarost/onetrades/utils"
)

// ==============SetLeverage=================
type SetLeverage struct {
	c        *Client
	symbol   *string
	leverage *int
}

func (s *SetLeverage) Symbol(symbol string) *SetLeverage {
	s.symbol = &symbol
	return s
}

func (s *SetLeverage) Leverage(leverage int) *SetLeverage {
	s.leverage = &leverage
	return s
}

func (s *SetLeverage) Do(ctx context.Context, opts ...utils.RequestOption) (res bool, err error) {
	r := &utils.Request{
		Method:     http.MethodPost,
		BaseURL:    s.c.BaseURL,
		Endpoint:   "/api/v5/account/set-leverage",
		TimeOffset: s.c.TimeOffset,
		SecType:    utils.SecTypeSigned,
	}

	m := utils.Params{
		"mgnMode": "cross",
	}

	if s.symbol != nil {
		m["instId"] = *s.symbol
	}

	if s.leverage != nil {
		m["lever"] = fmt.Sprintf("%d", *s.leverage)
	}
	r.SetFormParams(m)

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return false, err
	}

	var answ struct {
		Result []SetLeverageAnswer `json:"data"`
	}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return false, err
	}

	if len(answ.Result) == 0 {
		return false, errors.New("Zero Answer")
	}
	return true, nil
}

type SetLeverageAnswer struct {
	MgnMode string `json:"mgnMode"`
	InstId  string `json:"instId"`
	PosSide string `json:"posSide"`
}

// ==============SetPositionMode=================
type SetPositionMode struct {
	c    *Client
	mode *entity.PositionModeType
}

func (s *SetPositionMode) Mode(mode entity.PositionModeType) *SetPositionMode {
	s.mode = &mode
	return s
}

func (s *SetPositionMode) Do(ctx context.Context, opts ...utils.RequestOption) (res bool, err error) {
	r := &utils.Request{
		Method:     http.MethodPost,
		BaseURL:    s.c.BaseURL,
		Endpoint:   "/api/v5/account/set-position-mode",
		TimeOffset: s.c.TimeOffset,
		SecType:    utils.SecTypeSigned,
	}

	m := utils.Params{}

	if s.mode != nil {
		if *s.mode == entity.PositionModeTypeHedge {
			m["posMode"] = "long_short_mode"
		} else if *s.mode == entity.PositionModeTypeOneWay {
			m["posMode"] = "net_mode"
		}
	}
	r.SetFormParams(m)

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return false, err
	}

	var answ struct {
		Result []PositionMode `json:"data"`
	}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return false, err
	}

	if len(answ.Result) == 0 {
		return false, errors.New("Zero Answer")
	}
	return true, nil
}

type PositionMode struct {
	PosMode string `json:"posMode"`
}

// ==============GetPositions=================
type GetPositions struct {
	c *Client
}

func (s *GetPositions) Do(ctx context.Context, opts ...utils.RequestOption) (res []entity.Position, err error) {
	r := &utils.Request{
		Method:     http.MethodGet,
		BaseURL:    s.c.BaseURL,
		Endpoint:   "/api/v5/account/positions",
		TimeOffset: s.c.TimeOffset,
		SecType:    utils.SecTypeSigned,
	}

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}

	var answ struct {
		Result []Position `json:"data"`
	}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}
	return ConvertPositions(answ.Result), nil
}

type Position struct {
	InstID      string `json:"instId"`
	PosCcy      string `json:"posCcy,omitempty"`
	LiabCcy     string `json:"liabCcy,omitempty"`
	OptVal      string `json:"optVal,omitempty"`
	Ccy         string `json:"ccy"`
	PosID       string `json:"posId"`
	TradeID     string `json:"tradeId"`
	Pos         string `json:"pos"`
	AvailPos    string `json:"availPos,omitempty"`
	AvgPx       string `json:"avgPx"`
	Upl         string `json:"upl"`
	RealizedPnl string `json:"realizedPnl"`
	UplRatio    string `json:"uplRatio"`
	Lever       string `json:"lever"`
	LiqPx       string `json:"liqPx,omitempty"`
	MarkPx      string `json:"markPx,omitempty"`
	Imr         string `json:"imr,omitempty"`
	Margin      string `json:"margin,omitempty"`
	MgnRatio    string `json:"mgnRatio"`
	Mmr         string `json:"mmr"`
	Liab        string `json:"liab,omitempty"`
	Interest    string `json:"interest"`
	NotionalUsd string `json:"notionalUsd"`
	ADL         string `json:"adl"`
	Last        string `json:"last"`
	DeltaBS     string `json:"deltaBS"`
	DeltaPA     string `json:"deltaPA"`
	GammaBS     string `json:"gammaBS"`
	GammaPA     string `json:"gammaPA"`
	ThetaBS     string `json:"thetaBS"`
	ThetaPA     string `json:"thetaPA"`
	VegaBS      string `json:"vegaBS"`
	VegaPA      string `json:"vegaPA"`
	PosSide     string `json:"posSide"`
	MgnMode     string `json:"mgnMode"`
	InstType    string `json:"instType"`
	CTime       string `json:"cTime"`
	UTime       string `json:"uTime"`
}
