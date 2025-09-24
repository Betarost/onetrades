package kucoin

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Betarost/onetrades/entity"
	"github.com/Betarost/onetrades/utils"
)

// ==============getOrderList=================
type futures_getOrderList struct {
	convert futures_converts

	callAPI   func(ctx context.Context, r *utils.Request, opts ...utils.RequestOption) (data []byte, header *http.Header, err error)
	symbol    *string
	instType  *string
	orderType *entity.OrderType
	limit     *int
}

func (s *futures_getOrderList) Symbol(symbol string) *futures_getOrderList {
	s.symbol = &symbol
	return s
}

func (s *futures_getOrderList) InstType(instType string) *futures_getOrderList {
	s.instType = &instType
	return s
}

func (s *futures_getOrderList) OrderType(orderType entity.OrderType) *futures_getOrderList {
	s.orderType = &orderType
	return s
}

func (s *futures_getOrderList) Limit(limit int) *futures_getOrderList {
	s.limit = &limit
	return s
}

func (s *futures_getOrderList) Do(ctx context.Context, opts ...utils.RequestOption) (res []entity.Futures_OrdersList, err error) {
	r := &utils.Request{
		Method:   http.MethodGet,
		Endpoint: "/api/v1/orders",
		SecType:  utils.SecTypeSigned,
	}

	m := utils.Params{
		"status": "active",
	}

	if s.symbol != nil {
		m["symbol"] = *s.symbol
	}

	if s.limit != nil {
		m["pageSize"] = *s.limit
	}

	r.SetParams(m)

	data, _, err := s.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}
	var answ struct {
		Result struct {
			Items []futures_orderList `json:"items"`
		} `json:"data"`
	}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}

	return s.convert.convertOrderList(answ.Result.Items), nil
}

type futures_orderList struct {
	Symbol       string  `json:"symbol"`
	ID           string  `json:"id"`
	ClientOid    string  `json:"clientOid"`
	OpType       string  `json:"opType"`
	Type         string  `json:"type"`
	Side         string  `json:"side"`
	PositionSide string  `json:"positionSide"`
	Price        string  `json:"price"`
	Size         float64 `json:"size"`
	FilledSize   float64 `json:"filledSize"`
	TradeType    string  `json:"tradeType"`
	MarginMode   string  `json:"marginMode"`
	Leverage     string  `json:"leverage"`
	Status       string  `json:"status"`
	CreatedAt    int64   `json:"createdAt"`
	UpdatedAt    int64   `json:"updatedAt"`
}
