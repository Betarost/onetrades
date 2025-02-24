package futurebybit

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Betarost/onetrades/entity"
	"github.com/Betarost/onetrades/utils"
)

// ==============GetAccountBalance=================
type GetAccountBalance struct {
	c *Client
}

func (s *GetAccountBalance) Do(ctx context.Context, opts ...utils.RequestOption) (res []entity.AccountBalance, err error) {
	r := &utils.Request{
		Method:     http.MethodGet,
		BaseURL:    s.c.BaseURL,
		Endpoint:   "/v5/account/wallet-balance",
		TimeOffset: s.c.TimeOffset,
		SecType:    utils.SecTypeSigned,
	}

	r.SetParam("accountType", "UNIFIED")

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}

	var answ struct {
		Result struct {
			List []Balance `json:"list"`
		} `json:"result"`
	}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}
	return ConvertAccountBalance(answ.Result.List), nil
}

type Balance struct {
	TotalWalletBalance    string `json:"totalWalletBalance"`
	TotalAvailableBalance string `json:"totalAvailableBalance"`
	TotalPerpUPL          string `json:"totalPerpUPL"`
}
