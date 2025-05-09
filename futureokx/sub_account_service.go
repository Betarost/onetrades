package futureokx

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Betarost/onetrades/entity"
	"github.com/Betarost/onetrades/utils"
)

// ==============GetTransferHistory=================

type GetTransferHistory struct {
	c      *Client
	after  *string
	before *string
}

func (s *GetTransferHistory) After(after string) *GetTransferHistory {
	s.after = &after
	return s
}

func (s *GetTransferHistory) Before(before string) *GetTransferHistory {
	s.before = &before
	return s
}

func (s *GetTransferHistory) Do(ctx context.Context, opts ...utils.RequestOption) (res []entity.TransferHistory, err error) {
	r := &utils.Request{
		Method:     http.MethodGet,
		BaseURL:    s.c.BaseURL,
		Endpoint:   "/api/v5/asset/subaccount/bills",
		TimeOffset: s.c.TimeOffset,
		SecType:    utils.SecTypeSigned,
	}

	m := utils.Params{
		"ccy": "USDT",
	}

	if s.after != nil {
		m["after"] = *s.after
	}

	if s.before != nil {
		m["before"] = *s.before
	}

	r.SetParams(m)

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}

	var answ struct {
		Result []TransferHistory `json:"data"`
	}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}

	return ConvertTransferHistory(answ.Result), nil
}

type TransferHistory struct {
	BillId  string `json:"billId"`
	Ccy     string `json:"ccy"`
	Amt     string `json:"amt"`
	Type    string `json:"type"`
	SubAcct string `json:"subAcct"`
	Ts      string `json:"ts"`
}

// ===================FundsTransfer==================
type FundsTransfer struct {
	c      *Client
	subID  *string
	amount *string
	way    *string // 1:master-sub	2:sub-master
	from   *string // 6: Funding account	18: Trading account
	to     *string // 6: Funding account	18: Trading account
	tag    *string
}

func (s *FundsTransfer) Amount(amount string) *FundsTransfer {
	s.amount = &amount
	return s
}

func (s *FundsTransfer) Tag(tag string) *FundsTransfer {
	s.tag = &tag
	return s
}

func (s *FundsTransfer) Way(way string) *FundsTransfer {
	s.way = &way
	return s
}

func (s *FundsTransfer) From(from string) *FundsTransfer {
	s.from = &from
	return s
}

func (s *FundsTransfer) To(to string) *FundsTransfer {
	s.to = &to
	return s
}

func (s *FundsTransfer) SubID(subID string) *FundsTransfer {
	s.subID = &subID
	return s
}

func (s *FundsTransfer) Do(ctx context.Context, opts ...utils.RequestOption) (res bool, err error) {
	r := &utils.Request{
		Method:     http.MethodPost,
		BaseURL:    s.c.BaseURL,
		Endpoint:   "/api/v5/asset/transfer",
		TimeOffset: s.c.TimeOffset,
		SecType:    utils.SecTypeSigned,
	}

	m := utils.Params{
		"ccy": "USDT",
	}

	if s.amount != nil {
		m["amt"] = *s.amount
	}

	if s.way != nil {
		m["type"] = *s.way
	}

	if s.tag != nil {
		m["clientId"] = *s.tag
	}

	if s.from != nil {
		m["from"] = *s.from
	}

	if s.to != nil {
		m["to"] = *s.to
	}

	if s.subID != nil {
		m["subAcct"] = *s.subID
	}

	r.SetFormParams(m)

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return false, err
	}

	var answ struct {
		Result []TransferAnswer `json:"data"`
	}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return false, err
	}

	if len(answ.Result) == 0 {
		return false, errors.New("Zero Answer")
	}

	if answ.Result[0].TransId == "" {
		return false, errors.New("Empty Answer")
	}
	return true, nil
}

type TransferAnswer struct {
	TransId string `json:"transId"`
}

// ===================GetSubAccountFundingBalance==================
type GetSubAccountFundingBalance struct {
	c     *Client
	subID *string
}

func (s *GetSubAccountFundingBalance) SubID(subID string) *GetSubAccountFundingBalance {
	s.subID = &subID
	return s
}

func (s *GetSubAccountFundingBalance) Do(ctx context.Context, opts ...utils.RequestOption) (res entity.SubAccountFundingBalance, err error) {
	r := &utils.Request{
		Method:     http.MethodGet,
		BaseURL:    s.c.BaseURL,
		Endpoint:   "/api/v5/asset/subaccount/balances",
		TimeOffset: s.c.TimeOffset,
		SecType:    utils.SecTypeSigned,
	}

	m := utils.Params{"ccy": "USDT"}
	if s.subID != nil {
		m["subAcct"] = *s.subID
	}
	r.SetParams(m)

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}

	var answ struct {
		Result []SubAccountFundingBalance `json:"data"`
	}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}

	if len(answ.Result) == 0 {
		return res, errors.New("Zero Answer")
	}

	return ConvertSubAccountFundingBalance(answ.Result[0]), nil
}

type SubAccountFundingBalance struct {
	Ccy       string `json:"ccy"`
	Bal       string `json:"bal"`
	FrozenBal string `json:"frozenBal"`
	AvailBal  string `json:"availBal"`
}

// ===================GetSubAccountBalance==================
type GetSubAccountBalance struct {
	c     *Client
	subID *string
}

func (s *GetSubAccountBalance) SubID(subID string) *GetSubAccountBalance {
	s.subID = &subID
	return s
}

func (s *GetSubAccountBalance) Do(ctx context.Context, opts ...utils.RequestOption) (res entity.SubAccountBalance, err error) {
	r := &utils.Request{
		Method:     http.MethodGet,
		BaseURL:    s.c.BaseURL,
		Endpoint:   "/api/v5/account/subaccount/balances",
		TimeOffset: s.c.TimeOffset,
		SecType:    utils.SecTypeSigned,
	}

	m := utils.Params{}
	if s.subID != nil {
		m["subAcct"] = *s.subID
	}
	r.SetParams(m)

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}

	var answ struct {
		Result []SubAccountBalance `json:"data"`
	}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}

	if len(answ.Result) == 0 {
		return res, errors.New("Zero Answer")
	}

	return ConvertSubAccountBalance(answ.Result[0]), nil
}

type SubAccountBalance struct {
	TotalEq string `json:"totalEq"`
	Upl     string `json:"upl"`
	UTime   string `json:"uTime"`
}

// ===================GetSubAccountsLists==================
type GetSubAccountsLists struct {
	c *Client
}

func (s *GetSubAccountsLists) Do(ctx context.Context, opts ...utils.RequestOption) (res []entity.AccountInfo, err error) {
	r := &utils.Request{
		Method:     http.MethodGet,
		BaseURL:    s.c.BaseURL,
		Endpoint:   "/api/v5/users/subaccount/list",
		TimeOffset: s.c.TimeOffset,
		SecType:    utils.SecTypeSigned,
	}

	m := utils.Params{"enable": "true"}
	r.SetParams(m)

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}

	var answ struct {
		Result []SubAccountsLists `json:"data"`
	}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}

	return ConvertSubAccountInfo(answ.Result), nil
}

type SubAccountsLists struct {
	UID            string   `json:"uid"`
	Type           string   `json:"type"`
	Enable         bool     `json:"enable"`
	GAuth          bool     `json:"gAuth"`
	CanTransOut    bool     `json:"canTransOut"`
	IfDma          bool     `json:"ifDma"`
	SubAcct        string   `json:"subAcct"`
	Label          string   `json:"label"`
	Mobile         string   `json:"mobile"`
	FrozenFunc     []string `json:"frozenFunc"`
	Ts             string   `json:"ts"`
	SubAcctLv      string   `json:"subAcctLv"`
	FirstLvSubAcct string   `json:"firstLvSubAcct"`
}
