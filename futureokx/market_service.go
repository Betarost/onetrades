package futureokx

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Betarost/onetrades/entity"
	"github.com/Betarost/onetrades/utils"
)

// ==============GetTickers=================
type GetTickers struct {
	c *Client
}

func (s *GetTickers) Do(ctx context.Context, opts ...utils.RequestOption) (res []entity.Ticker, err error) {
	r := &utils.Request{
		Method:     http.MethodGet,
		BaseURL:    s.c.BaseURL,
		Endpoint:   "/api/v5/market/tickers",
		TimeOffset: s.c.TimeOffset,
		SecType:    utils.SecTypeNone,
	}

	m := utils.Params{
		"instType": "SWAP",
	}

	r.SetParams(m)

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}

	var answ struct {
		Result []Ticker `json:"data"`
	}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}

	if len(answ.Result) == 0 {
		return res, errors.New("Zero Answer")
	}

	return ConvertTickers(answ.Result), nil
}

type Ticker struct {
	InstId    string `json:"instId"`
	InstType  string `json:"instType"`
	Last      string `json:"last"`
	Open24h   string `json:"open24h"`
	VolCcy24h string `json:"volCcy24h"`
	SodUtc0   string `json:"sodUtc0"`
	SodUtc8   string `json:"sodUtc8"`
	Ts        string `json:"ts"`
}

// ==============GetKline=================
type GetKline struct {
	c         *Client
	symbol    *string
	timeFrame *entity.TimeFrameType
	limit     *int
}

func (s *GetKline) Symbol(symbol string) *GetKline {
	s.symbol = &symbol
	return s
}

func (s *GetKline) TimeFrame(timeFrame entity.TimeFrameType) *GetKline {
	s.timeFrame = &timeFrame
	return s
}

func (s *GetKline) Limit(limit int) *GetKline {
	s.limit = &limit
	return s
}

func (s *GetKline) Do(ctx context.Context, opts ...utils.RequestOption) (res []entity.Kline, err error) {
	r := &utils.Request{
		Method:     http.MethodGet,
		BaseURL:    s.c.BaseURL,
		Endpoint:   "/api/v5/market/mark-price-candles",
		TimeOffset: s.c.TimeOffset,
		SecType:    utils.SecTypeNone,
	}

	m := utils.Params{}

	if s.symbol != nil {
		m["instId"] = *s.symbol
	}

	if s.timeFrame != nil {
		m["bar"] = *s.timeFrame
	}

	if s.limit != nil {
		m["limit"] = *s.limit
	}

	r.SetParams(m)

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}

	var answ struct {
		Result [][]string `json:"data"`
	}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}

	if len(answ.Result) == 0 {
		return res, errors.New("Zero Answer")
	}

	return ConvertKline(answ.Result), nil
}

// ==============GetMarkPrices=================
type GetMarkPrices struct {
	c *Client
}

func (s *GetMarkPrices) Do(ctx context.Context, opts ...utils.RequestOption) (res []entity.MarkPrice, err error) {
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

	r.SetParams(m)

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}

	var answ struct {
		Result []MarkPrice `json:"data"`
	}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}

	if len(answ.Result) == 0 {
		return res, errors.New("Zero Answer")
	}

	return ConvertMarkPrices(answ.Result), nil
}

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
	InstId   string `json:"instId"`
	CtVal    string `json:"ctVal"`
	CtMult   string `json:"ctMult"`
	CtValCcy string `json:"ctValCcy"`
	TickSz   string `json:"tickSz"`
	LotSz    string `json:"lotSz"`
	MinSz    string `json:"minSz"`
	Lever    string `json:"lever"`
	State    string `json:"state"`
	RuleType string `json:"ruleType"`
}
