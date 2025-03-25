package futureokx

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Betarost/onetrades/entity"
	"github.com/Betarost/onetrades/utils"
)

// ===================GetAccountInfo==================
type GetAccountInfo struct {
	c *Client
}

func (s *GetAccountInfo) Do(ctx context.Context, opts ...utils.RequestOption) (res entity.AccountInfo, err error) {
	r := &utils.Request{
		Method:     http.MethodGet,
		BaseURL:    s.c.BaseURL,
		Endpoint:   "/api/v5/account/config",
		TimeOffset: s.c.TimeOffset,
		SecType:    utils.SecTypeSigned,
	}

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}

	var answ struct {
		Result []AccountInfo `json:"data"`
	}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}

	if len(answ.Result) == 0 {
		return res, errors.New("Zero Answer")
	}
	return ConvertAccountInfo(answ.Result[0]), nil
}

type AccountInfo struct {
	UID                 string `json:"uid"`
	MainUID             string `json:"mainUid"`
	AcctLv              string `json:"acctLv"`
	AcctStpMode         string `json:"acctStpMode"`
	PosMode             string `json:"posMode"`
	AutoLoan            bool   `json:"autoLoan"`
	GreeksType          string `json:"greeksType"`
	Level               string `json:"level"`
	LevelTmp            string `json:"levelTmp"`
	CtIsoMode           string `json:"ctIsoMode"`
	MgnIsoMode          string `json:"mgnIsoMode"`
	RoleType            string `json:"roleType"`
	SpotRoleType        string `json:"spotRoleType"`
	OpAuth              string `json:"opAuth"`
	KycLv               string `json:"kycLv"`
	Label               string `json:"label"`
	Ip                  string `json:"ip"`
	Perm                string `json:"perm"`
	LiquidationGear     string `json:"liquidationGear"`
	EnableSpotBorrow    bool   `json:"enableSpotBorrow"`
	SpotBorrowAutoRepay bool   `json:"spotBorrowAutoRepay"`
	Type                string `json:"type"`
}

// ===================SetAccountMode==================
type SetAccountMode struct {
	c    *Client
	mode *string
}

func (s *SetAccountMode) Mode(mode string) *SetAccountMode {
	s.mode = &mode
	return s
}

func (s *SetAccountMode) Do(ctx context.Context, opts ...utils.RequestOption) (res bool, err error) {
	r := &utils.Request{
		Method:     http.MethodPost,
		BaseURL:    s.c.BaseURL,
		Endpoint:   "/api/v5/account/set-account-level",
		TimeOffset: s.c.TimeOffset,
		SecType:    utils.SecTypeSigned,
	}

	m := utils.Params{}

	if s.mode != nil {
		m["acctLv"] = *s.mode
	}
	r.SetFormParams(m)

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return false, err
	}

	var answ struct {
		Result []AccountMode `json:"data"`
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

type AccountMode struct {
	AcctLv string `json:"acctLv"`
}

//===================GetAccountBalance==================

type GetAccountBalance struct {
	c *Client
}

func (s *GetAccountBalance) Do(ctx context.Context, opts ...utils.RequestOption) (res []entity.AccountBalance, err error) {
	r := &utils.Request{
		Method:     http.MethodGet,
		BaseURL:    s.c.BaseURL,
		Endpoint:   "/api/v5/account/balance",
		TimeOffset: s.c.TimeOffset,
		SecType:    utils.SecTypeSigned,
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}

	var answ struct {
		Result []Balance `json:"data"`
	}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}
	return ConvertAccountBalance(answ.Result), nil
}

type Balance struct {
	TotalEq            string           `json:"totalEq"`
	IsoEq              string           `json:"isoEq"`
	AdjEq              string           `json:"adjEq,omitempty"`
	OrdFroz            string           `json:"ordFroz,omitempty"`
	Imr                string           `json:"imr,omitempty"`
	Mmr                string           `json:"mmr,omitempty"`
	MgnRatio           string           `json:"mgnRatio,omitempty"`
	NotionalUsd        string           `json:"notionalUsd,omitempty"`
	NotionalUsdForSwap string           `json:"notionalUsdForSwap,omitempty"`
	Upl                string           `json:"upl,omitempty"`
	Details            []BalanceDetails `json:"details,omitempty"`
	UTime              string           `json:"uTime"`
}

type BalanceDetails struct {
	Ccy           string `json:"ccy"`
	Eq            string `json:"eq"`
	CashBal       string `json:"cashBal"`
	IsoEq         string `json:"isoEq,omitempty"`
	AvailEq       string `json:"availEq,omitempty"`
	DisEq         string `json:"disEq"`
	AvailBal      string `json:"availBal"`
	FrozenBal     string `json:"frozenBal"`
	OrdFrozen     string `json:"ordFrozen"`
	Liab          string `json:"liab,omitempty"`
	Upl           string `json:"upl,omitempty"`
	UplLib        string `json:"uplLib,omitempty"`
	CrossLiab     string `json:"crossLiab,omitempty"`
	IsoLiab       string `json:"isoLiab,omitempty"`
	MgnRatio      string `json:"mgnRatio,omitempty"`
	Interest      string `json:"interest,omitempty"`
	Twap          string `json:"twap,omitempty"`
	MaxLoan       string `json:"maxLoan,omitempty"`
	EqUsd         string `json:"eqUsd"`
	NotionalLever string `json:"notionalLever,omitempty"`
	StgyEq        string `json:"stgyEq"`
	IsoUpl        string `json:"isoUpl,omitempty"`
	UTime         string `json:"uTime"`
}

//===================GetAccountValuation==================

type GetAccountValuation struct {
	c *Client
}

func (s *GetAccountValuation) Do(ctx context.Context, opts ...utils.RequestOption) (res entity.AccountValuation, err error) {
	r := &utils.Request{
		Method:     http.MethodGet,
		BaseURL:    s.c.BaseURL,
		Endpoint:   "/api/v5/asset/asset-valuation",
		TimeOffset: s.c.TimeOffset,
		SecType:    utils.SecTypeSigned,
	}

	m := utils.Params{"ccy": "USDT"}
	r.SetParams(m)

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}

	var answ struct {
		Result []AccountValuation `json:"data"`
	}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}

	if len(answ.Result) == 0 {
		return res, errors.New("Zero Answer")
	}
	return ConvertAccountValuation(answ.Result[0]), nil

}

type AccountValuation struct {
	TotalBal string                  `json:"totalBal"`
	Details  AccountValuationDetails `json:"details,omitempty"`
	Ts       string                  `json:"ts"`
}

type AccountValuationDetails struct {
	Classic string `json:"classic"`
	Earn    string `json:"earn"`
	Funding string `json:"funding"`
	Trading string `json:"trading,omitempty"`
}
