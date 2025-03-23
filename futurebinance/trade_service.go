package futurebinance

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
	symbol *string
	begin  *string
	end    *string
	limit  *int
}

func (s *TradeHistoryOrder) Symbol(symbol string) *TradeHistoryOrder {
	s.symbol = &symbol
	return s
}

func (s *TradeHistoryOrder) Limit(limit int) *TradeHistoryOrder {
	s.limit = &limit
	return s
}

func (s *TradeHistoryOrder) Begin(begin string) *TradeHistoryOrder {
	s.begin = &begin
	return s
}

func (s *TradeHistoryOrder) End(end string) *TradeHistoryOrder {
	s.end = &end
	return s
}

func (s *TradeHistoryOrder) Do(ctx context.Context, opts ...utils.RequestOption) (res []entity.OrdersHistory, err error) {
	r := &utils.Request{
		Method:     http.MethodGet,
		BaseURL:    s.c.BaseURL,
		Endpoint:   "/fapi/v1/userTrades",
		TimeOffset: s.c.TimeOffset,
		SecType:    utils.SecTypeSigned,
	}

	m := utils.Params{}

	if s.symbol != nil {
		m["symbol"] = *s.symbol
	}

	if s.limit != nil {
		m["limit"] = *s.limit
	}

	if s.begin != nil {
		m["startTime"] = *s.begin
	}

	if s.end != nil {
		m["endTime"] = *s.end
	}

	r.SetParams(m)

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}

	answ := []HistoryOrder{}
	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}

	return ConvertHistoryOrders(answ), nil
}

type HistoryOrder struct {
	Symbol          string `json:"symbol"`
	Buyer           bool   `json:"buyer"`
	Maker           bool   `json:"maker"`
	Commission      string `json:"commission"`
	CommissionAsset string `json:"commissionAsset"`
	Id              int64  `json:"id"`
	Price           string `json:"price"`
	Qty             string `json:"qty"`
	QuoteQty        string `json:"quoteQty"`
	RealizedPnl     string `json:"realizedPnl"`
	Side            string `json:"side"`
	PositionSide    string `json:"positionSide"`
	Time            int64  `json:"time"`
}

// ==============GetDownloadLinkHistoryOrder=================
type GetDownloadLinkHistoryOrder struct {
	c          *Client
	downloadId *string
}

func (s *GetDownloadLinkHistoryOrder) DownloadId(downloadId string) *GetDownloadLinkHistoryOrder {
	s.downloadId = &downloadId
	return s
}

func (s *GetDownloadLinkHistoryOrder) Do(ctx context.Context, opts ...utils.RequestOption) (res string, err error) {
	r := &utils.Request{
		Method:     http.MethodGet,
		BaseURL:    s.c.BaseURL,
		Endpoint:   "/fapi/v1/order/asyn/id",
		TimeOffset: s.c.TimeOffset,
		SecType:    utils.SecTypeSigned,
	}

	m := utils.Params{}

	if s.downloadId != nil {
		m["downloadId"] = *s.downloadId
	}

	r.SetParams(m)

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}

	answ := DownloadLinkHistoryOrder{}
	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}
	return answ.Url, nil
}

type DownloadLinkHistoryOrder struct {
	DownloadId string `json:"downloadId"`
	Status     string `json:"downstatusloadId"`
	Url        string `json:"url"`
}

// ==============GetDownloadIdHistoryOrder=================
type GetDownloadIdHistoryOrder struct {
	c     *Client
	begin *int64
	end   *int64
}

func (s *GetDownloadIdHistoryOrder) Begin(begin int64) *GetDownloadIdHistoryOrder {
	s.begin = &begin
	return s
}

func (s *GetDownloadIdHistoryOrder) End(end int64) *GetDownloadIdHistoryOrder {
	s.end = &end
	return s
}

func (s *GetDownloadIdHistoryOrder) Do(ctx context.Context, opts ...utils.RequestOption) (res string, err error) {
	r := &utils.Request{
		Method:     http.MethodGet,
		BaseURL:    s.c.BaseURL,
		Endpoint:   "/fapi/v1/order/asyn",
		TimeOffset: s.c.TimeOffset,
		SecType:    utils.SecTypeSigned,
	}

	m := utils.Params{}

	if s.begin != nil {
		m["startTime"] = *s.begin
	}

	if s.end != nil {
		m["endTime"] = *s.end
	}

	r.SetParams(m)

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}

	answ := DownloadIdHistoryOrder{}
	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}
	return answ.DownloadId, nil
}

type DownloadIdHistoryOrder struct {
	DownloadId string `json:"downloadId"`
}
