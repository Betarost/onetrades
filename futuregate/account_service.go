package futuregate

import (
	"context"
	"encoding/json"
	"log"
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
		Endpoint:   "/api/v4/futures/usdt/accounts",
		TimeOffset: s.c.TimeOffset,
		SecType:    utils.SecTypeSigned,
	}

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}

	log.Println("=3b11ba=", string(data))
	var answ Balance

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}
	return ConvertAccountBalance(answ), nil
}

type Balance struct {
	Currency       string `json:"currency"`
	Total          string `json:"total"`
	Unrealised_pnl string `json:"unrealised_pnl"`
	Available      string `json:"available"`
}
