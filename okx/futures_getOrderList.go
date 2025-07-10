package okx

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Betarost/onetrades/entity"
	"github.com/Betarost/onetrades/utils"
)

// ==============getOrderList=================
type futures_getOrderList struct {
	convert futures_converts

	callAPI   func(ctx context.Context, r *utils.Request, opts ...utils.RequestOption) (data []byte, header *http.Header, err error)
	symbol    *string
	orderType *entity.OrderType
	limit     *int
}

func (s *futures_getOrderList) Symbol(symbol string) *futures_getOrderList {
	s.symbol = &symbol
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
		Endpoint: "/api/v5/trade/orders-pending",
		SecType:  utils.SecTypeSigned,
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

	data, _, err := s.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}

	var answ struct {
		Result []futures_orderList `json:"data"`
	}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}

	return s.convert.convertOrderList(answ.Result), nil
}

type futures_orderList struct {
	InstId         string                             `json:"instId"`
	OrdId          string                             `json:"ordId"`
	ClOrdId        string                             `json:"clOrdId"`
	Px             string                             `json:"px"`
	Sz             string                             `json:"sz"`
	AttachAlgoOrds []futures_orderList_attachAlgoOrds `json:"AttachAlgoOrds"`
	PosSide        string                             `json:"posSide"`
	OrdType        string                             `json:"ordType"`
	TdMode         string                             `json:"tdMode"`
	InstType       string                             `json:"instType"`
	Lever          string                             `json:"lever"`
	Side           string                             `json:"side"`
	State          string                             `json:"state"`
	FillSz         string                             `json:"fillSz"`
	IsTpLimit      string                             `json:"isTpLimit"`
	UTime          string                             `json:"uTime"`
	CTime          string                             `json:"cTime"`
}

type futures_orderList_attachAlgoOrds struct {
	AttachAlgoId string `json:"attachAlgoId"`
	SlOrdPx      string `json:"slOrdPx"`
	SlTriggerPx  string `json:"slTriggerPx"`
	TpOrdPx      string `json:"tpOrdPx"`
	TpTriggerPx  string `json:"tpTriggerPx"`
}
