package futurebingx

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Betarost/onetrades/entity"
	"github.com/Betarost/onetrades/utils"
)

// ==============GetOrdersHistory=================
type GetOrdersHistory struct {
	c         *Client
	symbol    *string
	currency  *string
	orderId   *int64
	startTime *int64
	endTime   *int64
	limit     *int
}

func (s *GetOrdersHistory) Symbol(symbol string) *GetOrdersHistory {
	s.symbol = &symbol
	return s
}

func (s *GetOrdersHistory) Currency(currency string) *GetOrdersHistory {
	s.currency = &currency
	return s
}

func (s *GetOrdersHistory) OrderId(orderId int64) *GetOrdersHistory {
	s.orderId = &orderId
	return s
}

func (s *GetOrdersHistory) StartTime(startTime int64) *GetOrdersHistory {
	s.startTime = &startTime
	return s
}

func (s *GetOrdersHistory) EndTime(endTime int64) *GetOrdersHistory {
	s.endTime = &endTime
	return s
}

func (s *GetOrdersHistory) Limit(limit int) *GetOrdersHistory {
	s.limit = &limit
	return s
}

func (s *GetOrdersHistory) Do(ctx context.Context, opts ...utils.RequestOption) (res []entity.OrdersHistory, err error) {
	r := &utils.Request{
		Method:     http.MethodGet,
		BaseURL:    s.c.BaseURL,
		Endpoint:   "/openApi/swap/v2/trade/allOrders",
		TimeOffset: s.c.TimeOffset,
		SecType:    utils.SecTypeSigned,
	}

	m := utils.Params{}

	if s.symbol != nil {
		m["symbol"] = *s.symbol
	}

	if s.currency != nil {
		m["currency"] = *s.currency
	}

	if s.orderId != nil {
		m["orderId"] = *s.orderId
	}

	if s.startTime != nil {
		m["startTime"] = *s.startTime
	}

	if s.endTime != nil {
		m["endTime"] = *s.endTime
	}

	if s.limit != nil {
		m["limit"] = *s.limit
	}

	r.SetParams(m)

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}

	var answ struct {
		Result struct {
			Orders []OrdersHistory `json:"orders"`
		} `json:"data"`
	}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}
	return ConvertOrdersHistory(answ.Result.Orders), nil
}

type OrdersHistory struct {
	Symbol                 string  `json:"symbol" bson:"symbol"`
	OrderID                int64   `json:"orderId" bson:"orderId"`
	Side                   string  `json:"side" bson:"side"`
	PositionSide           string  `json:"positionSide" bson:"positionSide"`
	Type                   string  `json:"type" bson:"type"`
	OrigQty                string  `json:"origQty" bson:"origQty"`
	Price                  string  `json:"price" bson:"price"`
	ExecutedQty            string  `json:"executedQty" bson:"executedQty"`
	AvgPrice               string  `json:"avgPrice" bson:"avgPrice"`
	CumQuote               string  `json:"cumQuote" bson:"cumQuote"`
	StopPrice              string  `json:"stopPrice" bson:"stopPrice"`
	Profit                 string  `json:"profit" bson:"profit"`
	Commission             string  `json:"commission" bson:"commission"`
	Status                 string  `json:"status" bson:"status"`
	Time                   int64   `json:"time" bson:"time"`
	UpdateTime             int64   `json:"updateTime" bson:"updateTime"`
	ClientOrderID          string  `json:"clientOrderId" bson:"clientOrderId"`
	Leverage               string  `json:"leverage" bson:"leverage"`
	AdvanceAttr            int     `json:"advanceAttr" bson:"advanceAttr"`
	PositionID             int64   `json:"positionID" bson:"positionID"`
	TakeProfitEntrustPrice float64 `json:"takeProfitEntrustPrice" bson:"takeProfitEntrustPrice"`
	StopLossEntrustPrice   float64 `json:"stopLossEntrustPrice" bson:"stopLossEntrustPrice"`
	OrderType              string  `json:"orderType" bson:"orderType"`
	WorkingType            string  `json:"workingType" bson:"workingType"`
	StopGuaranteed         string  `json:"stopGuaranteed" bson:"stopGuaranteed"`
	TriggerOrderID         int64   `json:"triggerOrderId" bson:"triggerOrderId"`
}
