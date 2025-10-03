package okx

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

func (s *futures_ordersHistory) Page(page int64) *futures_ordersHistory {
	s.page = &page
	return s
}

func (s *futures_ordersHistory) OrderID(orderID string) *futures_ordersHistory {
	s.orderID = &orderID
	return s
}

func (s *futures_ordersHistory) Do(ctx context.Context, opts ...utils.RequestOption) (res []entity.Futures_OrdersHistory, err error) {

	r := &utils.Request{
		Method:   http.MethodGet,
		Endpoint: "/api/v5/trade/orders-history",
		SecType:  utils.SecTypeSigned,
	}

	m := utils.Params{
		"instType": "SWAP",
		"state":    "filled",
	}

	if s.symbol != nil {
		m["instId"] = *s.symbol
	}
	if s.limit != nil && *s.limit > 0 {
		m["limit"] = *s.limit
	}
	if s.page != nil && *s.page > 0 {
		m["page"] = *s.page
	}
	if s.startTime != nil {
		m["begin"] = *s.startTime
	}
	if s.endTime != nil {
		m["end"] = *s.endTime
	}

	if s.orderID != nil {
		m["before"] = *s.orderID
	}

	r.SetParams(m)

	data, _, err := s.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}

	var answ struct {
		Result []futures_ordersHistory_Response `json:"data"`
	}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}

	return s.convert.convertOrdersHistory(answ.Result), nil
	//===========
}

type futures_ordersHistory_Response struct {
	InstId    string `json:"instId"`
	OrdId     string `json:"ordId"`
	ClOrdId   string `json:"clOrdId"`
	Side      string `json:"side"`
	PosSide   string `json:"posSide"`
	Sz        string `json:"sz"`
	FillSz    string `json:"fillSz"`
	AccFillSz string `json:"accFillSz"`
	Px        string `json:"px"`
	AvgPx     string `json:"avgPx"`
	Fee       string `json:"fee"`
	Pnl       string `json:"pnl"`
	Lever     string `json:"felevere"`
	OrdType   string `json:"ordType"`
	State     string `json:"state"`
	TdMode    string `json:"tdMode"`
	CTime     string `json:"cTime"`
	UTime     string `json:"uTime"`
	// Symbol          string `json:"symbol"`
	// OrderId         int64  `json:"orderId"`
	// Side            string `json:"side"`
	// PositionSide    string `json:"positionSide"`
	// Type            string `json:"type"`
	// OrigQty         string `json:"origQty"`
	// Price           string `json:"price"`
	// ExecutedQty     string `json:"executedQty"`
	// AvgPrice        string `json:"avgPrice"`
	// CumQuote        string `json:"cumQuote"`
	// Profit          string `json:"profit"`
	// Commission      string `json:"commission"`
	// Status          string `json:"status"`
	// ClientOrderId   string `json:"clientOrderId"`
	// Leverage        string `json:"leverage"`
	// PositionID      int64  `json:"positionID"`
	// OnlyOnePosition bool   `json:"onlyOnePosition"`
	// Time            int64  `json:"time"`
	// UpdateTime      int64  `json:"updateTime"`
}
