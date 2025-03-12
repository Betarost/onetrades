package futureokx

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/Betarost/onetrades/entity"
	"github.com/Betarost/onetrades/utils"
)

// ==============TradeHistoryOrder=================

type TradeHistoryOrder struct {
	c      *Client
	after  *string
	before *string
	begin  *string
	end    *string
}

func (s *TradeHistoryOrder) After(after string) *TradeHistoryOrder {
	s.after = &after
	return s
}

func (s *TradeHistoryOrder) Before(before string) *TradeHistoryOrder {
	s.before = &before
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
		Endpoint:   "/api/v5/trade/orders-history",
		TimeOffset: s.c.TimeOffset,
		SecType:    utils.SecTypeSigned,
	}

	m := utils.Params{
		"instType": "SWAP",
		"state":    "filled",
	}

	if s.after != nil {
		m["after"] = *s.after
	}

	if s.before != nil {
		m["before"] = *s.before
	}

	if s.begin != nil {
		m["begin"] = *s.begin
	}

	if s.end != nil {
		m["end"] = *s.end
	}

	r.SetParams(m)

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}

	var answ struct {
		Result []HistoryOrder `json:"data"`
	}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}

	return ConvertHistoryOrders(answ.Result), nil
}

type HistoryOrder struct {
	InstType  string `json:"instType"`
	InstID    string `json:"instId"`
	Ccy       string `json:"ccy"`
	OrdId     string `json:"ordId"`
	ClOrdId   string `json:"clOrdId"`
	Tag       string `json:"tag"`
	Px        string `json:"px"`
	Sz        string `json:"sz"`
	OrdType   string `json:"ordType"`
	Side      string `json:"side"`
	PosSide   string `json:"posSide"`
	TdMode    string `json:"tdMode"`
	AccFillSz string `json:"accFillSz"`
	AvgPx     string `json:"avgPx"`
	State     string `json:"state"`
	Lever     string `json:"lever"`
	Fee       string `json:"fee"`
	Pnl       string `json:"pnl"`
	Category  string `json:"category"`
	CTime     string `json:"cTime"`
	UTime     string `json:"uTime"`
}

// ==============TradeCancelOrders=================

type TradeCancelOrders struct {
	c        *Client
	symbol   *string
	orderIDs *[]string
}

func (s *TradeCancelOrders) Symbol(symbol string) *TradeCancelOrders {
	s.symbol = &symbol
	return s
}

func (s *TradeCancelOrders) OrderIDs(orderIDs []string) *TradeCancelOrders {
	s.orderIDs = &orderIDs
	return s
}

func (s *TradeCancelOrders) Do(ctx context.Context, opts ...utils.RequestOption) (res bool, err error) {
	r := &utils.Request{
		Method:     http.MethodPost,
		BaseURL:    s.c.BaseURL,
		Endpoint:   "/api/v5/trade/cancel-batch-orders",
		TimeOffset: s.c.TimeOffset,
		SecType:    utils.SecTypeSigned,
	}

	m := utils.Params{
		"is_batch": "true",
	}

	if s.symbol != nil && s.orderIDs != nil {
		orderIDs := []OrdersIDs{}
		for _, item := range *s.orderIDs {
			orderIDs = append(orderIDs, OrdersIDs{
				InstId: *s.symbol,
				OrdId:  item,
			})
		}
		j, err := json.Marshal(orderIDs)
		if err != nil {
			return false, err
		}
		m["is_batch"] = string(j)
	}

	r.SetFormParams(m)

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return false, err
	}

	var answ struct {
		Result []PlaceOrder `json:"data"`
	}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return false, err
	}

	if len(answ.Result) == 0 {
		return false, errors.New("Zero Answer")
	}

	if answ.Result[0].SCode != "0" {
		return false, errors.New(answ.Result[0].SMsg)
	}
	return true, nil
}

type CancelOrders struct {
	ClOrdId string `json:"clOrdId"`
	OrdId   string `json:"ordId"`
	Tag     string `json:"tag"`
	Ts      string `json:"ts"`
	SCode   string `json:"sCode"`
	SMsg    string `json:"sMsg"`
}

type OrdersIDs struct {
	InstId string `json:"instId"`
	OrdId  string `json:"ordId"`
}

// ==============TradePlaceOrder=================
type TradePlaceOrder struct {
	c            *Client
	symbol       *string
	positionSide *entity.PositionSideType
	side         *entity.SideType
	orderType    *entity.OrderType
	size         *string
	price        *string
}

func (s *TradePlaceOrder) Size(size string) *TradePlaceOrder {
	s.size = &size
	return s
}

func (s *TradePlaceOrder) Price(price string) *TradePlaceOrder {
	s.price = &price
	return s
}

func (s *TradePlaceOrder) Symbol(symbol string) *TradePlaceOrder {
	s.symbol = &symbol
	return s
}

func (s *TradePlaceOrder) PositionSide(positionSide entity.PositionSideType) *TradePlaceOrder {
	s.positionSide = &positionSide
	return s
}

func (s *TradePlaceOrder) Side(side entity.SideType) *TradePlaceOrder {
	s.side = &side
	return s
}

func (s *TradePlaceOrder) OrderType(orderType entity.OrderType) *TradePlaceOrder {
	s.orderType = &orderType
	return s
}

func (s *TradePlaceOrder) Do(ctx context.Context, opts ...utils.RequestOption) (res bool, err error) {
	r := &utils.Request{
		Method:     http.MethodPost,
		BaseURL:    s.c.BaseURL,
		Endpoint:   "/api/v5/trade/order",
		TimeOffset: s.c.TimeOffset,
		SecType:    utils.SecTypeSigned,
	}

	m := utils.Params{
		"tdMode": "cross",
	}

	if s.size != nil {
		m["sz"] = *s.size
	}

	if s.price != nil {
		m["px"] = *s.price
	}

	if s.symbol != nil {
		m["instId"] = *s.symbol
	}

	if s.positionSide != nil {
		m["posSide"] = strings.ToLower(string(*s.positionSide))
	}

	if s.side != nil {
		m["side"] = strings.ToLower(string(*s.side))
	}

	if s.orderType != nil {
		m["ordType"] = strings.ToLower(string(*s.orderType))
	}

	r.SetFormParams(m)

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return false, err
	}

	var answ struct {
		Result []PlaceOrder `json:"data"`
	}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return false, err
	}

	if len(answ.Result) == 0 {
		return false, errors.New("Zero Answer")
	}

	if answ.Result[0].SCode != "0" {
		return false, errors.New(answ.Result[0].SMsg)
	}
	return true, nil
}

type PlaceOrder struct {
	ClOrdId string `json:"clOrdId"`
	OrdId   string `json:"ordId"`
	Tag     string `json:"tag"`
	Ts      string `json:"ts"`
	SCode   string `json:"sCode"`
	SMsg    string `json:"sMsg"`
}

// ==============GetOrderList=================
type GetOrderList struct {
	c         *Client
	symbol    *string
	orderType *entity.OrderType
	limit     *int
}

func (s *GetOrderList) Symbol(symbol string) *GetOrderList {
	s.symbol = &symbol
	return s
}

func (s *GetOrderList) OrderType(orderType entity.OrderType) *GetOrderList {
	s.orderType = &orderType
	return s
}

func (s *GetOrderList) Limit(limit int) *GetOrderList {
	s.limit = &limit
	return s
}

func (s *GetOrderList) Do(ctx context.Context, opts ...utils.RequestOption) (res []entity.OrderList, err error) {
	r := &utils.Request{
		Method:     http.MethodGet,
		BaseURL:    s.c.BaseURL,
		Endpoint:   "/api/v5/trade/orders-pending",
		TimeOffset: s.c.TimeOffset,
		SecType:    utils.SecTypeSigned,
	}

	m := utils.Params{
		"instType": "SWAP",
	}

	if s.symbol != nil {
		m["instId"] = *s.symbol
	}

	if s.limit != nil {
		m["limit"] = *s.limit
	}

	if s.orderType != nil {
		m["ordType"] = strings.ToLower(string(*s.orderType))
	}

	r.SetParams(m)

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}

	var answ struct {
		Result []OrderList `json:"data"`
	}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}

	return ConvertOrderList(answ.Result), nil
}

type OrderList struct {
	InstId  string `json:"instId"`
	OrdId   string `json:"ordId"`
	Px      string `json:"px"`
	Sz      string `json:"sz"`
	PosSide string `json:"posSide"`
	OrdType string `json:"ordType"`
	Side    string `json:"side"`
	State   string `json:"state"`
	UTime   string `json:"uTime"`
}
