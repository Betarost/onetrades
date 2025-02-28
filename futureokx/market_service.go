package futureokx

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Betarost/onetrades/entity"
	"github.com/Betarost/onetrades/utils"
)

// ==============GetMarkPrice=================
type GetMarkPrice struct {
	c      *Client
	symbol *string
}

func (s *GetMarkPrice) Symbol(symbol string) *GetMarkPrice {
	s.symbol = &symbol
	return s
}

func (s *GetMarkPrice) Do(ctx context.Context, opts ...utils.RequestOption) (res float64, err error) {
	r := &utils.Request{
		Method:     http.MethodGet,
		BaseURL:    s.c.BaseURL,
		Endpoint:   "/api/v5/public/mark-price",
		TimeOffset: s.c.TimeOffset,
		SecType:    utils.SecTypeNone,
	}

	m := utils.Params{
		"instType": "SWAP",
	}

	if s.symbol != nil {
		m["instId"] = *s.symbol
	}

	r.SetParams(m)

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return 0, err
	}

	var answ struct {
		Result []MarkPrice `json:"data"`
	}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return 0, err
	}

	if len(answ.Result) == 0 {
		return 0, errors.New("Zero Answer")
	}

	return utils.StringToFloat(answ.Result[0].MarkPx), nil
}

type MarkPrice struct {
	InstId   string `json:"instId"`
	InstType string `json:"instType"`
	MarkPx   string `json:"markPx"`
	Ts       string `json:"ts"`
}

// ==============GetContractsInfo=================
type GetContractsInfo struct {
	c      *Client
	symbol *string
}

func (s *GetContractsInfo) Symbol(symbol string) *GetContractsInfo {
	s.symbol = &symbol
	return s
}

func (s *GetContractsInfo) Do(ctx context.Context, opts ...utils.RequestOption) (res []entity.ContractInfo, err error) {
	r := &utils.Request{
		Method:     http.MethodGet,
		BaseURL:    s.c.BaseURL,
		Endpoint:   "/api/v5/public/instruments",
		TimeOffset: s.c.TimeOffset,
		SecType:    utils.SecTypeNone,
	}

	m := utils.Params{
		"instType": "SWAP",
	}

	if s.symbol != nil {
		m["instId"] = *s.symbol
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
	InstId string `json:"instId"`
	CtVal  string `json:"ctVal"`
	CtMult string `json:"ctMult"`
	TickSz string `json:"tickSz"`
	LotSz  string `json:"lotSz"`
	Lever  string `json:"lever"`
}
