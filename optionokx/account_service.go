package optionokx

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Betarost/onetrades/entity"
	"github.com/Betarost/onetrades/utils"
)

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
