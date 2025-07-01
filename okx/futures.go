package okx

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Betarost/onetrades/entity"
	"github.com/Betarost/onetrades/utils"
)

// ==============GetPositions=================
type getPositions struct {
	callAPI func(ctx context.Context, r *utils.Request, opts ...utils.RequestOption) (data []byte, header *http.Header, err error)
}

func (s *getPositions) Do(ctx context.Context, opts ...utils.RequestOption) (res []entity.Futures_Positions, err error) {
	r := &utils.Request{
		Method:   http.MethodGet,
		Endpoint: "/api/v5/account/positions",
		SecType:  utils.SecTypeSigned,
	}

	data, _, err := s.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}

	var answ struct {
		Result []futures_Position `json:"data"`
	}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}
	return futures_convertPositions(answ.Result), nil
}

type futures_Position struct {
	InstID      string `json:"instId"`
	PosCcy      string `json:"posCcy,omitempty"`
	LiabCcy     string `json:"liabCcy,omitempty"`
	OptVal      string `json:"optVal,omitempty"`
	Ccy         string `json:"ccy"`
	PosID       string `json:"posId"`
	TradeID     string `json:"tradeId"`
	Pos         string `json:"pos"`
	AvailPos    string `json:"availPos,omitempty"`
	AvgPx       string `json:"avgPx"`
	Upl         string `json:"upl"`
	RealizedPnl string `json:"realizedPnl"`
	UplRatio    string `json:"uplRatio"`
	Lever       string `json:"lever"`
	LiqPx       string `json:"liqPx,omitempty"`
	MarkPx      string `json:"markPx,omitempty"`
	Imr         string `json:"imr,omitempty"`
	Margin      string `json:"margin,omitempty"`
	MgnRatio    string `json:"mgnRatio"`
	Mmr         string `json:"mmr"`
	Liab        string `json:"liab,omitempty"`
	Interest    string `json:"interest"`
	NotionalUsd string `json:"notionalUsd"`
	ADL         string `json:"adl"`
	Last        string `json:"last"`
	DeltaBS     string `json:"deltaBS"`
	DeltaPA     string `json:"deltaPA"`
	GammaBS     string `json:"gammaBS"`
	GammaPA     string `json:"gammaPA"`
	ThetaBS     string `json:"thetaBS"`
	ThetaPA     string `json:"thetaPA"`
	VegaBS      string `json:"vegaBS"`
	VegaPA      string `json:"vegaPA"`
	PosSide     string `json:"posSide"`
	MgnMode     string `json:"mgnMode"`
	InstType    string `json:"instType"`
	CTime       string `json:"cTime"`
	UTime       string `json:"uTime"`
}

// ==============getOrderList=================
type futures_getOrderList struct {
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

	return futures_convertOrderList(answ.Result), nil
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
