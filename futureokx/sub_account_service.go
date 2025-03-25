package futureokx

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Betarost/onetrades/entity"
	"github.com/Betarost/onetrades/utils"
)

// ===================GetSubAccountInfo==================
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
