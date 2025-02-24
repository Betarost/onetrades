package futurebitget

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
		Endpoint:   "/api/v2/mix/account/account",
		TimeOffset: s.c.TimeOffset,
		SecType:    utils.SecTypeSigned,
	}

	r.SetParam("productType", "USDT-FUTURES")

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
	MarginCoin    string `json:"marginCoin"`
	Available     string `json:"available"`
	AccountEquity string `json:"accountEquity"`
	UnrealizedPL  string `json:"unrealizedPL"`
}
