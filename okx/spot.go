package okx

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/Betarost/onetrades/entity"
	"github.com/Betarost/onetrades/utils"
)

// ==============AmendOrder=================

type amendOrder struct {
	callAPI func(ctx context.Context, r *utils.Request, opts ...utils.RequestOption) (data []byte, header *http.Header, err error)

	symbol   *string
	orderID  *string
	newSize  *string
	newPrice *string
}

func (s *amendOrder) Symbol(symbol string) *amendOrder {
	s.symbol = &symbol
	return s
}

func (s *amendOrder) OrderID(orderID string) *amendOrder {
	s.orderID = &orderID
	return s
}

func (s *amendOrder) NewSize(newSize string) *amendOrder {
	s.newSize = &newSize
	return s
}

func (s *amendOrder) NewPrice(newPrice string) *amendOrder {
	s.newPrice = &newPrice
	return s
}

func (s *amendOrder) Do(ctx context.Context, opts ...utils.RequestOption) (res []entity.PlaceOrder, err error) {
	r := &utils.Request{
		Method:   http.MethodPost,
		Endpoint: "/api/v5/trade/amend-order",
		SecType:  utils.SecTypeSigned,
	}

	m := utils.Params{}

	if s.symbol != nil {
		m["instId"] = *s.symbol
	}

	if s.orderID != nil {
		m["ordId"] = *s.orderID
	}

	if s.newSize != nil {
		m["newSz"] = *s.newSize
	}

	if s.newPrice != nil {
		m["newPx"] = *s.newPrice
	}

	r.SetFormParams(m)

	data, _, err := s.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}

	var answ struct {
		Result []placeOrder_Response `json:"data"`
	}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}

	if len(answ.Result) == 0 {
		return res, errors.New("Zero Answer")
	}

	if answ.Result[0].SCode != "0" {
		return res, errors.New(answ.Result[0].SMsg)
	}
	return convertPlaceOrder(answ.Result), nil
}

// ==============TradeCancelOrders=================

type cancelOrder struct {
	callAPI func(ctx context.Context, r *utils.Request, opts ...utils.RequestOption) (data []byte, header *http.Header, err error)

	symbol  *string
	orderID *string
}

func (s *cancelOrder) Symbol(symbol string) *cancelOrder {
	s.symbol = &symbol
	return s
}

func (s *cancelOrder) OrderID(orderID string) *cancelOrder {
	s.orderID = &orderID
	return s
}

func (s *cancelOrder) Do(ctx context.Context, opts ...utils.RequestOption) (res []entity.PlaceOrder, err error) {
	r := &utils.Request{
		Method:   http.MethodPost,
		Endpoint: "/api/v5/trade/cancel-order",
		SecType:  utils.SecTypeSigned,
	}

	m := utils.Params{}

	if s.symbol != nil {
		m["instId"] = *s.symbol
	}

	if s.orderID != nil {
		m["ordId"] = *s.orderID
	}

	r.SetFormParams(m)

	data, _, err := s.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}

	var answ struct {
		Result []placeOrder_Response `json:"data"`
	}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}

	if len(answ.Result) == 0 {
		return res, errors.New("Zero Answer")
	}

	if answ.Result[0].SCode != "0" {
		return res, errors.New(answ.Result[0].SMsg)
	}
	return convertPlaceOrder(answ.Result), nil
}

// ==============multiCancelOrders=================

type multiCancelOrders struct {
	callAPI func(ctx context.Context, r *utils.Request, opts ...utils.RequestOption) (data []byte, header *http.Header, err error)

	symbol   *string
	orderIDs *[]string
}

func (s *multiCancelOrders) Symbol(symbol string) *multiCancelOrders {
	s.symbol = &symbol
	return s
}

func (s *multiCancelOrders) OrderIDs(orderIDs []string) *multiCancelOrders {
	s.orderIDs = &orderIDs
	return s
}

func (s *multiCancelOrders) Do(ctx context.Context, opts ...utils.RequestOption) (res []entity.PlaceOrder, err error) {
	r := &utils.Request{
		Method:   http.MethodPost,
		Endpoint: "/api/v5/trade/cancel-batch-orders",
		SecType:  utils.SecTypeSigned,
	}

	m := utils.Params{
		"is_batch": "true",
	}

	if s.symbol != nil && s.orderIDs != nil {
		orderIDs := []ordersIDs{}
		for _, item := range *s.orderIDs {
			orderIDs = append(orderIDs, ordersIDs{
				InstId: *s.symbol,
				OrdId:  item,
			})
		}
		j, err := json.Marshal(orderIDs)
		if err != nil {
			return res, err
		}
		m["is_batch"] = string(j)
	}

	r.SetFormParams(m)

	data, _, err := s.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}

	var answ struct {
		Result []placeOrder_Response `json:"data"`
	}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}

	if len(answ.Result) == 0 {
		return res, errors.New("Zero Answer")
	}

	if answ.Result[0].SCode != "0" {
		return res, errors.New(answ.Result[0].SMsg)
	}
	return convertPlaceOrder(answ.Result), nil
}

type ordersIDs struct {
	InstId string `json:"instId"`
	OrdId  string `json:"ordId"`
}

// ==============placeOrder=================
type placeOrder struct {
	callAPI func(ctx context.Context, r *utils.Request, opts ...utils.RequestOption) (data []byte, header *http.Header, err error)

	symbol        *string
	side          *entity.SideType
	size          *string
	price         *string
	orderType     *entity.OrderType
	clientOrderID *string
	positionSide  *entity.PositionSideType
	tradeMode     *entity.MarginModeType
	tpPrice       *string
	slPrice       *string
}

func (s *placeOrder) TradeMode(tradeMode entity.MarginModeType) *placeOrder {
	s.tradeMode = &tradeMode
	return s
}

func (s *placeOrder) SlPrice(slPrice string) *placeOrder {
	s.slPrice = &slPrice
	return s
}

func (s *placeOrder) TpPrice(tpPrice string) *placeOrder {
	s.tpPrice = &tpPrice
	return s
}

func (s *placeOrder) Symbol(symbol string) *placeOrder {
	s.symbol = &symbol
	return s
}

func (s *placeOrder) Side(side entity.SideType) *placeOrder {
	s.side = &side
	return s
}

func (s *placeOrder) Size(size string) *placeOrder {
	s.size = &size
	return s
}

func (s *placeOrder) Price(price string) *placeOrder {
	s.price = &price
	return s
}

func (s *placeOrder) OrderType(orderType entity.OrderType) *placeOrder {
	s.orderType = &orderType
	return s
}

func (s *placeOrder) ClientOrderID(clientOrderID string) *placeOrder {
	s.clientOrderID = &clientOrderID
	return s
}

func (s *placeOrder) PositionSide(positionSide entity.PositionSideType) *placeOrder {
	s.positionSide = &positionSide
	return s
}

func (s *placeOrder) Do(ctx context.Context, opts ...utils.RequestOption) (res []entity.PlaceOrder, err error) {
	r := &utils.Request{
		Method:   http.MethodPost,
		Endpoint: "/api/v5/trade/order",
		SecType:  utils.SecTypeSigned,
	}

	m := utils.Params{
		"tdMode": "cash",
	}

	if s.tradeMode != nil {
		if *s.tradeMode == entity.MarginModeTypeCross {
			m["tdMode"] = "cross"
		} else if *s.tradeMode == entity.MarginModeTypeIsolated {
			m["tdMode"] = "isolated"
		}
	}

	if s.symbol != nil {
		m["instId"] = *s.symbol
	}

	if s.side != nil {
		m["side"] = strings.ToLower(string(*s.side))
	}

	if s.size != nil {
		m["sz"] = *s.size
	}

	if s.price != nil {
		m["px"] = *s.price
	}

	if s.orderType != nil {
		m["ordType"] = strings.ToLower(string(*s.orderType))
	}

	if s.clientOrderID != nil {
		m["clOrdId"] = *s.clientOrderID
	}

	if s.positionSide != nil {
		m["posSide"] = strings.ToLower(string(*s.positionSide))
	}

	if s.tpPrice != nil || s.slPrice != nil {
		attachAlgoOrds := []orderList_attachAlgoOrds{{}}
		if s.tpPrice != nil {
			attachAlgoOrds[0].TpTriggerPx = *s.tpPrice
			attachAlgoOrds[0].TpOrdPx = "-1"
		}

		if s.slPrice != nil {
			attachAlgoOrds[0].SlTriggerPx = *s.slPrice
			attachAlgoOrds[0].SlOrdPx = "-1"
		}
		j, err := json.Marshal(attachAlgoOrds)
		if err != nil {
			return res, err
		}

		m["attachAlgoOrds"] = string(j)
	}

	r.SetFormParams(m)

	data, _, err := s.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}

	var answ struct {
		Result []placeOrder_Response `json:"data"`
	}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}

	if len(answ.Result) == 0 {
		return res, errors.New("Zero Answer")
	}

	if answ.Result[0].SCode != "0" {
		return res, errors.New(answ.Result[0].SMsg)
	}
	return convertPlaceOrder(answ.Result), nil
}

type placeOrder_Response struct {
	ClOrdId string `json:"clOrdId"`
	OrdId   string `json:"ordId"`
	Tag     string `json:"tag"`
	Ts      string `json:"ts"`
	SCode   string `json:"sCode"`
	SMsg    string `json:"sMsg"`
}

// ==============getOrderList=================
type getOrderList struct {
	callAPI   func(ctx context.Context, r *utils.Request, opts ...utils.RequestOption) (data []byte, header *http.Header, err error)
	symbol    *string
	orderType *entity.OrderType
	limit     *int
}

func (s *getOrderList) Symbol(symbol string) *getOrderList {
	s.symbol = &symbol
	return s
}

func (s *getOrderList) OrderType(orderType entity.OrderType) *getOrderList {
	s.orderType = &orderType
	return s
}

func (s *getOrderList) Limit(limit int) *getOrderList {
	s.limit = &limit
	return s
}

func (s *getOrderList) Do(ctx context.Context, opts ...utils.RequestOption) (res []entity.OrdersPendingList, err error) {
	r := &utils.Request{
		Method:   http.MethodGet,
		Endpoint: "/api/v5/trade/orders-pending",
		SecType:  utils.SecTypeSigned,
	}

	m := utils.Params{
		"instType": "SPOT",
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
		Result []orderList `json:"data"`
	}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}

	return convertOrderList(answ.Result), nil
}

type orderList struct {
	InstId         string                     `json:"instId"`
	OrdId          string                     `json:"ordId"`
	ClOrdId        string                     `json:"clOrdId"`
	Px             string                     `json:"px"`
	Sz             string                     `json:"sz"`
	AttachAlgoOrds []orderList_attachAlgoOrds `json:"AttachAlgoOrds"`
	PosSide        string                     `json:"posSide"`
	OrdType        string                     `json:"ordType"`
	TdMode         string                     `json:"tdMode"`
	InstType       string                     `json:"instType"`
	Lever          string                     `json:"lever"`
	Side           string                     `json:"side"`
	State          string                     `json:"state"`
	IsTpLimit      string                     `json:"isTpLimit"`
	UTime          string                     `json:"uTime"`
	CTime          string                     `json:"cTime"`
}

type orderList_attachAlgoOrds struct {
	AttachAlgoId string `json:"attachAlgoId"`
	SlOrdPx      string `json:"slOrdPx"`
	SlTriggerPx  string `json:"slTriggerPx"`
	TpOrdPx      string `json:"tpOrdPx"`
	TpTriggerPx  string `json:"tpTriggerPx"`
}
