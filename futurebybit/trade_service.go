package futurebybit

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Betarost/onetrades/entity"
	"github.com/Betarost/onetrades/utils"
)

// ==============TradeHistoryOrder=================

type TradeHistoryOrder struct {
	c      *Client
	begin  *string
	end    *string
	cursor *string
	limit  *int
}

func (s *TradeHistoryOrder) Begin(begin string) *TradeHistoryOrder {
	s.begin = &begin
	return s
}

func (s *TradeHistoryOrder) End(end string) *TradeHistoryOrder {
	s.end = &end
	return s
}

func (s *TradeHistoryOrder) Cursor(cursor string) *TradeHistoryOrder {
	s.cursor = &cursor
	return s
}

func (s *TradeHistoryOrder) Limit(limit int) *TradeHistoryOrder {
	s.limit = &limit
	return s
}

func (s *TradeHistoryOrder) Do(ctx context.Context, opts ...utils.RequestOption) (res []entity.OrdersHistory, err error) {
	r := &utils.Request{
		Method:     http.MethodGet,
		BaseURL:    s.c.BaseURL,
		Endpoint:   "/v5/position/closed-pnl",
		TimeOffset: s.c.TimeOffset,
		SecType:    utils.SecTypeSigned,
	}

	m := utils.Params{
		"category": "linear",
	}

	if s.begin != nil {
		m["startTime"] = *s.begin
	}

	if s.end != nil {
		m["endTime"] = *s.end
	}

	if s.cursor != nil {
		m["cursor"] = *s.cursor
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
			List []HistoryOrder `json:"list"`
		} `json:"result"`
	}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}

	return ConvertHistoryOrders(answ.Result.List), nil
}

type HistoryOrder struct {
	Symbol        string `json:"symbol"`
	OrderId       string `json:"orderId"`
	OrderType     string `json:"orderType"`
	Leverage      string `json:"leverage"`
	Side          string `json:"side"`
	ClosedPnl     string `json:"closedPnl"`
	AvgEntryPrice string `json:"avgEntryPrice"`
	Qty           string `json:"qty"`
	CumEntryValue string `json:"cumEntryValue"`
	OrderPrice    string `json:"orderPrice"`
	ClosedSize    string `json:"closedSize"`
	AvgExitPrice  string `json:"avgExitPrice"`
	ExecType      string `json:"execType"`
	FillCount     string `json:"fillCount"`
	CumExitValue  string `json:"cumExitValue"`
	CreatedTime   string `json:"createdTime"`
	UpdatedTime   string `json:"updatedTime"`
}
