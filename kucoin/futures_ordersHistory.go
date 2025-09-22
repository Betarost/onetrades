package kucoin

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Betarost/onetrades/entity"
	"github.com/Betarost/onetrades/utils"
)

type futures_ordersHistory struct {
	callAPI func(ctx context.Context, r *utils.Request, opts ...utils.RequestOption) (data []byte, header *http.Header, err error)
	convert futures_converts

	symbol    *string
	startTime *int64
	endTime   *int64
	limit     *int64
	page      *int64
}

func (s *futures_ordersHistory) Symbol(symbol string) *futures_ordersHistory {
	s.symbol = &symbol
	return s
}

func (s *futures_ordersHistory) StartTime(startTime int64) *futures_ordersHistory {
	s.startTime = &startTime
	return s
}

func (s *futures_ordersHistory) EndTime(endTime int64) *futures_ordersHistory {
	s.endTime = &endTime
	return s
}

func (s *futures_ordersHistory) Limit(limit int64) *futures_ordersHistory {
	s.limit = &limit
	return s
}

func (s *futures_ordersHistory) Page(page int64) *futures_ordersHistory {
	s.page = &page
	return s
}

func (s *futures_ordersHistory) Do(ctx context.Context, opts ...utils.RequestOption) (res []entity.Futures_OrdersHistory, err error) {

	r := &utils.Request{
		Method:   http.MethodGet,
		Endpoint: "/api/v1/orders",
		SecType:  utils.SecTypeSigned,
	}

	m := utils.Params{
		"status": "done",
	}

	if s.symbol != nil {
		m["symbol"] = *s.symbol
	}
	if s.limit != nil && *s.limit > 0 {
		m["pageSize"] = *s.limit
	}
	if s.page != nil && *s.page > 0 {
		m["currentPage"] = *s.page
	}
	if s.startTime != nil {
		m["startAt"] = *s.startTime
	}
	if s.endTime != nil {
		m["endAt"] = *s.endTime
	}

	r.SetParams(m)

	data, _, err := s.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}

	var answ struct {
		Result struct {
			Items []futures_ordersHistory_Response `json:"items"`
		} `json:"data"`
	}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}

	return s.convert.convertOrdersHistory(answ.Result.Items), nil
	//===========
}

type futures_ordersHistory_Response struct {
	Symbol       string  `json:"symbol"`
	ID           string  `json:"id"`
	Type         string  `json:"type"`
	Side         string  `json:"side"`
	PositionSide string  `json:"positionSide"`
	Price        string  `json:"price"`
	AvgDealPrice string  `json:"avgDealPrice"`
	Size         float64 `json:"size"`
	Value        string  `json:"value"`
	Leverage     string  `json:"leverage"`
	ClientOid    string  `json:"clientOid"`
	CreatedAt    int64   `json:"createdAt"`
	UpdatedAt    int64   `json:"updatedAt"`
	MarginMode   string  `json:"marginMode"`
	OpType       string  `json:"opType"`
	FilledSize   float64 `json:"filledSize"`
	FilledValue  string  `json:"filledValue"`
	Status       string  `json:"status"`
}
