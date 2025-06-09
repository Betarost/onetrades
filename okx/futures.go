package okx

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/Betarost/onetrades/entity"
	"github.com/Betarost/onetrades/utils"
)

// ==============GetInstrumentsInfo=================
type futures_getInstrumentsInfo struct {
	callAPI func(ctx context.Context, r *utils.Request, opts ...utils.RequestOption) (data []byte, header *http.Header, err error)

	symbol *string
}

func (s *futures_getInstrumentsInfo) Symbol(symbol string) *futures_getInstrumentsInfo {
	s.symbol = &symbol
	return s
}

func (s *futures_getInstrumentsInfo) Do(ctx context.Context, opts ...utils.RequestOption) (res []entity.Futures_InstrumentsInfo, err error) {
	r := &utils.Request{
		Method:   http.MethodGet,
		Endpoint: "/api/v5/public/instruments",
		SecType:  utils.SecTypeNone,
	}

	m := utils.Params{
		"instType": "SWAP",
	}

	if s.symbol != nil {
		m["instId"] = *s.symbol
	}

	r.SetParams(m)

	data, _, err := s.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}
	var answ struct {
		Result []futures_instrumentsInfo `json:"data"`
	}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}
	return futures_convertInstrumentsInfo(answ.Result), nil
}

type futures_instrumentsInfo struct {
	InstId   string `json:"instId"`
	InstType string `json:"instType"`
	CtVal    string `json:"ctVal"`
	CtMult   string `json:"ctMult"`
	BaseCcy  string `json:"baseCcy"`
	QuoteCcy string `json:"quoteCcy"`
	CtValCcy string `json:"ctValCcy"`
	TickSz   string `json:"tickSz"`
	LotSz    string `json:"lotSz"`
	MinSz    string `json:"minSz"`
	Lever    string `json:"lever"`
	State    string `json:"state"`
	RuleType string `json:"ruleType"`
}

// ==============SetPositionMode=================
type setPositionMode struct {
	callAPI func(ctx context.Context, r *utils.Request, opts ...utils.RequestOption) (data []byte, header *http.Header, err error)

	mode *entity.PositionModeType
}

func (s *setPositionMode) Mode(mode entity.PositionModeType) *setPositionMode {
	s.mode = &mode
	return s
}

func (s *setPositionMode) Do(ctx context.Context, opts ...utils.RequestOption) (res bool, err error) {
	r := &utils.Request{
		Method:   http.MethodPost,
		Endpoint: "/api/v5/account/set-position-mode",
		SecType:  utils.SecTypeSigned,
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

	data, _, err := s.callAPI(ctx, r, opts...)
	if err != nil {
		return false, err
	}

	var answ struct {
		Result []positionMode `json:"data"`
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

type positionMode struct {
	PosMode string `json:"posMode"`
}

// ==============SetLeverage=================
type setLeverage struct {
	callAPI func(ctx context.Context, r *utils.Request, opts ...utils.RequestOption) (data []byte, header *http.Header, err error)

	symbol       *string
	leverage     *string
	positionSide *entity.PositionSideType
	marginMode   *entity.MarginModeType
}

func (s *setLeverage) Symbol(symbol string) *setLeverage {
	s.symbol = &symbol
	return s
}

func (s *setLeverage) Leverage(leverage string) *setLeverage {
	s.leverage = &leverage
	return s
}

func (s *setLeverage) MarginMode(marginMode entity.MarginModeType) *setLeverage {
	s.marginMode = &marginMode
	return s
}

func (s *setLeverage) PositionSide(positionSide entity.PositionSideType) *setLeverage {
	s.positionSide = &positionSide
	return s
}

func (s *setLeverage) Do(ctx context.Context, opts ...utils.RequestOption) (res entity.Futures_Leverage, err error) {
	r := &utils.Request{
		Method:   http.MethodPost,
		Endpoint: "/api/v5/account/set-leverage",
		SecType:  utils.SecTypeSigned,
	}

	m := utils.Params{}

	if s.symbol != nil {
		m["instId"] = *s.symbol
	}

	if s.leverage != nil {
		m["lever"] = *s.leverage
	}

	if s.positionSide != nil {
		m["posSide"] = strings.ToLower(string(*s.positionSide))
	}

	if s.marginMode != nil {
		if *s.marginMode == entity.MarginModeTypeCross {
			m["mgnMode"] = "cross"
		} else if *s.marginMode == entity.MarginModeTypeIsolated {
			m["mgnMode"] = "isolated"
		}
	}
	r.SetFormParams(m)

	data, _, err := s.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}

	var answ struct {
		Result []setLeverageAnswer `json:"data"`
	}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}

	if len(answ.Result) == 0 {
		return res, errors.New("Zero Answer")
	}
	return futures_convertSetLeverage(answ.Result), nil
}

type setLeverageAnswer struct {
	MgnMode string `json:"mgnMode"`
	Lever   string `json:"lever"`
	InstId  string `json:"instId"`
	PosSide string `json:"posSide"`
}
