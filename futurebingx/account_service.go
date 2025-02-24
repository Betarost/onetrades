package futurebingx

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
		Endpoint:   "/openApi/swap/v3/user/balance",
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
	Asset            string `json:"asset"`
	Balance          string `json:"balance"`
	UnrealizedProfit string `json:"unrealizedProfit"`
	AvailableMargin  string `json:"availableMargin"`
	Equity           string `json:"equity"`
}
