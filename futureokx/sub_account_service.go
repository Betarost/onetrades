package futureokx

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/Betarost/onetrades/entity"
	"github.com/Betarost/onetrades/utils"
)

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

	log.Println("=79bd79=", string(data))

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
