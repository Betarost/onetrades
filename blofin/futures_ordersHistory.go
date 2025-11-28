package blofin

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
	// page/cursor Blofin обычно делает через before/after, можно при необходимости добавить позднее
	orderID *string
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

func (s *futures_ordersHistory) OrderID(orderID string) *futures_ordersHistory {
	s.orderID = &orderID
	return s
}

func (s *futures_ordersHistory) Do(ctx context.Context, opts ...utils.RequestOption) (res []entity.Futures_OrdersHistory, err error) {

	r := &utils.Request{
		Method:   http.MethodGet,
		Endpoint: "/api/v1/trade/orders-history",
		SecType:  utils.SecTypeSigned,
	}

	m := utils.Params{"state": "filled"}

	if s.symbol != nil {
		m["instId"] = *s.symbol
	}

	if s.startTime != nil {
		m["startTime"] = *s.startTime
	}
	if s.endTime != nil {
		m["endTime"] = *s.endTime
	}
	if s.limit != nil && *s.limit > 0 {
		m["limit"] = *s.limit
	}
	if s.orderID != nil {
		// в некоторых биржах это before/after, но для простоты держим orderId — при тесте будет видно
		m["orderId"] = *s.orderID
	}

	r.SetParams(m)

	data, _, err := s.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}

	var answ struct {
		Result []futures_ordersHistory_Response `json:"data"`
	}

	if err = json.Unmarshal(data, &answ); err != nil {
		return res, err
	}

	return s.convert.convertOrdersHistory(answ.Result), nil
}

type futures_ordersHistory_Response struct {
	InstId        string `json:"instId"`
	OrderId       string `json:"orderId"`
	ClientOrderId string `json:"clientOrderId"`
	Side          string `json:"side"`
	PositionSide  string `json:"positionSide"`
	Size          string `json:"size"`
	FilledSize    string `json:"filledSize"`
	Price         string `json:"price"`
	AveragePrice  string `json:"averagePrice"` // было AvgPrice + tag "avgPrice"
	Fee           string `json:"fee"`
	Pnl           string `json:"pnl"` // было RealisedPnl + tag "realisedPnl"
	Leverage      string `json:"leverage"`
	OrderType     string `json:"orderType"`
	State         string `json:"state"`
	MarginMode    string `json:"marginMode"`
	CreateTime    string `json:"createTime"`
	UpdateTime    string `json:"updateTime"`
}
