package futuremexc

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Betarost/onetrades/entity"
	"github.com/Betarost/onetrades/utils"
)

type GetContractInfo struct {
	c      *Client
	symbol *string
}

func (s *GetContractInfo) Symbol(symbol string) *GetContractInfo {
	s.symbol = &symbol
	return s
}

func (s *GetContractInfo) Do(ctx context.Context, opts ...utils.RequestOption) (res entity.ContractInfo, err error) {
	r := &utils.Request{
		Method:     http.MethodGet,
		BaseURL:    s.c.BaseURL,
		Endpoint:   "/api/v1/contract/detail",
		TimeOffset: s.c.TimeOffset,
		SecType:    utils.SecTypeNone,
	}

	m := utils.Params{}

	if s.symbol != nil {
		m["symbol"] = *s.symbol
	}

	r.SetParams(m)

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}
	var answ struct {
		Result ContractInfo `json:"data"`
	}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}
	return ConvertContractInfo(answ.Result), nil
}

type ContractInfo struct {
	Symbol       string  `json:"symbol"`
	ContractSize float64 `json:"contractSize"`
}
