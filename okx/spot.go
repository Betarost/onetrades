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
	return ConvertPlaceOrder(answ.Result), nil
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
		orderIDs := []OrdersIDs{}
		for _, item := range *s.orderIDs {
			orderIDs = append(orderIDs, OrdersIDs{
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
	return ConvertPlaceOrder(answ.Result), nil
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

// ==============GetInstrumentsInfo=================
type getInstrumentsInfo struct {
	callAPI func(ctx context.Context, r *utils.Request, opts ...utils.RequestOption) (data []byte, header *http.Header, err error)

	symbol *string
}

func (s *getInstrumentsInfo) Symbol(symbol string) *getInstrumentsInfo {
	s.symbol = &symbol
	return s
}

func (s *getInstrumentsInfo) Do(ctx context.Context, opts ...utils.RequestOption) (res []entity.InstrumentsInfo, err error) {
	r := &utils.Request{
		Method:   http.MethodGet,
		Endpoint: "/api/v5/public/instruments",
		SecType:  utils.SecTypeNone,
	}

	m := utils.Params{
		"instType": "SPOT",
	}

	if s.symbol != nil {
		m["instId"] = *s.symbol
	}

	r.SetParams(m)

	data, _, err := s.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}
	var answ struct {
		Result []instrumentsInfo `json:"data"`
	}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}
	return ConvertInstrumentsInfo(answ.Result), nil
}

type instrumentsInfo struct {
	InstId   string `json:"instId"`
	InstType string `json:"instType"`
	CtVal    string `json:"ctVal"`
	CtMult   string `json:"ctMult"`
	BaseCcy  string `json:"baseCcy"`
	QuoteCcy string `json:"quoteCcy"`
	CtValCcy string `json:"ctValCcy"`
	TickSz   string `json:"tickSz"`
	LotSz    string `json:"lotSz"`
	MinSz    string `json:"minSz"`
	Lever    string `json:"lever"`
	State    string `json:"state"`
	RuleType string `json:"ruleType"`
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

func (s *placeOrder) Do(ctx context.Context, opts ...utils.RequestOption) (res []entity.PlaceOrder, err error) {
	r := &utils.Request{
		Method:   http.MethodPost,
		Endpoint: "/api/v5/trade/order",
		SecType:  utils.SecTypeSigned,
	}

	m := utils.Params{
		"tdMode": "cash",
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
	return ConvertPlaceOrder(answ.Result), nil
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

	return ConvertOrderList(answ.Result), nil
}

type orderList struct {
	InstId  string `json:"instId"`
	OrdId   string `json:"ordId"`
	Px      string `json:"px"`
	Sz      string `json:"sz"`
	PosSide string `json:"posSide"`
	OrdType string `json:"ordType"`
	TdMode  string `json:"tdMode"`
	Side    string `json:"side"`
	State   string `json:"state"`
	UTime   string `json:"uTime"`
}
