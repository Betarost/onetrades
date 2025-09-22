package kucoin

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Betarost/onetrades/entity"
	"github.com/Betarost/onetrades/utils"
)

type futures_setLeverage struct {
	callAPI func(ctx context.Context, r *utils.Request, opts ...utils.RequestOption) (data []byte, header *http.Header, err error)
	convert futures_converts

	symbol   *string
	leverage *string
}

func (s *futures_setLeverage) Symbol(symbol string) *futures_setLeverage {
	s.symbol = &symbol
	return s
}

func (s *futures_setLeverage) Leverage(leverage string) *futures_setLeverage {
	s.leverage = &leverage
	return s
}

func (s *futures_setLeverage) Do(ctx context.Context, opts ...utils.RequestOption) (res entity.Futures_Leverage, err error) {
	r := &utils.Request{
		Method:   http.MethodPost,
		Endpoint: "/api/v2/changeCrossUserLeverage",
		SecType:  utils.SecTypeSigned,
	}

	m := utils.Params{}

	if s.symbol != nil {
		m["symbol"] = *s.symbol
		res.Symbol = *s.symbol
	}

	if s.leverage != nil {
		m["leverage"] = *s.leverage
		res.Leverage = *s.leverage
	}

	r.SetFormParams(m)

	data, _, err := s.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}
	var answ struct {
		Result bool `json:"data"`
	}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}

	if answ.Result {
		return res, nil
	}

	return res, errors.New("Result False")

}
