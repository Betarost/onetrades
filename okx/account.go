package okx

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Betarost/onetrades/entity"
	"github.com/Betarost/onetrades/utils"
)

// ===================GetAccountInfo==================
type getAccountInfo struct {
	callAPI func(ctx context.Context, r *utils.Request, opts ...utils.RequestOption) (data []byte, header *http.Header, err error)
}

func (s *getAccountInfo) Do(ctx context.Context, opts ...utils.RequestOption) (res entity.AccountInformation, err error) {
	r := &utils.Request{
		Method:   http.MethodGet,
		Endpoint: "/api/v5/account/config",
		SecType:  utils.SecTypeSigned,
	}

	data, _, err := s.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}

	var answ struct {
		Result []accountInfo `json:"data"`
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

type accountInfo struct {
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

//===================getTradingAccountBalance==================

type getTradingAccountBalance struct {
	callAPI func(ctx context.Context, r *utils.Request, opts ...utils.RequestOption) (data []byte, header *http.Header, err error)
}

func (s *getTradingAccountBalance) Do(ctx context.Context, opts ...utils.RequestOption) (res []entity.TradingAccountBalance, err error) {
	r := &utils.Request{
		Method:   http.MethodGet,
		Endpoint: "/api/v5/account/balance",
		SecType:  utils.SecTypeSigned,
	}
	data, _, err := s.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}

	var answ struct {
		Result []tradingBalance `json:"data"`
	}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}
	return ConvertTradingAccountBalance(answ.Result), nil
}

type tradingBalance struct {
	TotalEq            string                  `json:"totalEq"`
	IsoEq              string                  `json:"isoEq"`
	AdjEq              string                  `json:"adjEq,omitempty"`
	AvailEq            string                  `json:"availEq,omitempty"`
	OrdFroz            string                  `json:"ordFroz,omitempty"`
	Imr                string                  `json:"imr,omitempty"`
	Mmr                string                  `json:"mmr,omitempty"`
	MgnRatio           string                  `json:"mgnRatio,omitempty"`
	NotionalUsd        string                  `json:"notionalUsd,omitempty"`
	NotionalUsdForSwap string                  `json:"notionalUsdForSwap,omitempty"`
	Upl                string                  `json:"upl,omitempty"`
	Details            []tradingBalanceDetails `json:"details,omitempty"`
	UTime              string                  `json:"uTime"`
}

type tradingBalanceDetails struct {
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

//===================getTradigetFundingAccountBalancengAccountBalance==================

type getFundingAccountBalance struct {
	callAPI func(ctx context.Context, r *utils.Request, opts ...utils.RequestOption) (data []byte, header *http.Header, err error)
}

func (s *getFundingAccountBalance) Do(ctx context.Context, opts ...utils.RequestOption) (res []entity.FundingAccountBalance, err error) {
	r := &utils.Request{
		Method:   http.MethodGet,
		Endpoint: "/api/v5/asset/balances",
		SecType:  utils.SecTypeSigned,
	}
	data, _, err := s.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}

	var answ struct {
		Result []fundingBalance `json:"data"`
	}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}
	return ConvertFundingAccountBalance(answ.Result), nil
}

type fundingBalance struct {
	Ccy       string `json:"ccy"`
	AvailBal  string `json:"availBal"`
	Bal       string `json:"bal"`
	FrozenBal string `json:"frozenBal"`
}
