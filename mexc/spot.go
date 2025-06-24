package mexc

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Betarost/onetrades/entity"
	"github.com/Betarost/onetrades/utils"
)

// ==============GetInstrumentsInfo=================
type spot_getInstrumentsInfo struct {
	callAPI func(ctx context.Context, r *utils.Request, opts ...utils.RequestOption) (data []byte, header *http.Header, err error)
	convert spot_converts
	symbol  *string
}

func (s *spot_getInstrumentsInfo) Symbol(symbol string) *spot_getInstrumentsInfo {
	s.symbol = &symbol
	return s
}

func (s *spot_getInstrumentsInfo) Do(ctx context.Context, opts ...utils.RequestOption) (res []entity.InstrumentsInfo, err error) {
	r := &utils.Request{
		Method:   http.MethodGet,
		Endpoint: "/api/v3/exchangeInfo",
		SecType:  utils.SecTypeNone,
	}

	m := utils.Params{}
	if s.symbol != nil {
		m["symbol"] = *s.symbol
	}
	r.SetParams(m)

	data, _, err := s.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}
	log.Println("=2b295c=", string(data))
	answ := spot_instrumentsInfo{}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}

	return s.convert.convertInstrumentsInfo(answ), nil
}

type spot_instrumentsInfo struct {
	Symbols []struct {
		Symbol             string `json:"symbol"`
		Status             string `json:"status"`
		BaseAsset          string `json:"baseAsset"`
		QuoteAsset         string `json:"quoteAsset"`
		BaseAssetPrecision int64  `json:"baseAssetPrecision"`
		QuotePrecision     int64  `json:"quotePrecision"`

		// MinQty      float64 `json:"minQty"`
		// MaxQty      float64 `json:"maxQty"`
		// MinNotional float64 `json:"minNotional"`
		// MaxNotional float64 `json:"maxNotional"`
		// TickSize    float64 `json:"tickSize"`
		// StepSize    float64 `json:"stepSize"`
	} `json:"symbols"`
}

// ===================GetBalance==================
type spot_getBalance struct {
	callAPI func(ctx context.Context, r *utils.Request, opts ...utils.RequestOption) (data []byte, header *http.Header, err error)
	convert spot_converts
}

func (s *spot_getBalance) Do(ctx context.Context, opts ...utils.RequestOption) (res []entity.AssetsBalance, err error) {
	r := &utils.Request{
		Method:   http.MethodGet,
		Endpoint: "/api/v3/account",
		SecType:  utils.SecTypeSigned,
	}

	data, _, err := s.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}
	answ := spot_Balance{}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}

	return s.convert.convertBalance(answ), nil
}

type spot_Balance struct {
	Balances []struct {
		Asset  string `json:"asset"`
		Free   string `json:"free"`
		Locked string `json:"locked"`
	} `json:"balances"`
}

// ==============placeOrder=================
type spot_placeOrder struct {
	callAPI func(ctx context.Context, r *utils.Request, opts ...utils.RequestOption) (data []byte, header *http.Header, err error)
	convert spot_converts

	symbol        *string
	side          *entity.SideType
	orderType     *entity.OrderType
	size          *string
	price         *string
	clientOrderID *string

	tpPrice *string
	slPrice *string
}

func (s *spot_placeOrder) Symbol(symbol string) *spot_placeOrder {
	s.symbol = &symbol
	return s
}

func (s *spot_placeOrder) Side(side entity.SideType) *spot_placeOrder {
	s.side = &side
	return s
}

func (s *spot_placeOrder) OrderType(orderType entity.OrderType) *spot_placeOrder {
	s.orderType = &orderType
	return s
}

func (s *spot_placeOrder) Size(size string) *spot_placeOrder {
	s.size = &size
	return s
}

func (s *spot_placeOrder) Price(price string) *spot_placeOrder {
	s.price = &price
	return s
}

func (s *spot_placeOrder) ClientOrderID(clientOrderID string) *spot_placeOrder {
	s.clientOrderID = &clientOrderID
	return s
}

func (s *spot_placeOrder) Do(ctx context.Context, opts ...utils.RequestOption) (res []entity.PlaceOrder, err error) {
	r := &utils.Request{
		Method:   http.MethodPost,
		Endpoint: "/api/v3/order",
		SecType:  utils.SecTypeSigned,
	}

	m := utils.Params{}

	if s.symbol != nil {
		m["symbol"] = *s.symbol
	}

	if s.side != nil {
		m["side"] = string(*s.side)
	}

	if s.orderType != nil {
		m["type"] = string(*s.orderType)
	}

	if s.size != nil {
		m["quantity"] = *s.size
	}

	if s.price != nil {
		m["price"] = *s.price
	}

	if s.clientOrderID != nil {
		m["newClientOrderId"] = *s.clientOrderID
	}

	r.SetFormParams(m)

	data, _, err := s.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}

	answ := placeOrder_Response{}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}

	return s.convert.convertPlaceOrder(answ), nil
}

type placeOrder_Response struct {
	Symbol       string `json:"symbol"`
	OrderId      string `json:"orderId"`
	Price        string `json:"price"`
	OrigQty      string `json:"origQty"`
	Type         string `json:"type"`
	Side         string `json:"side"`
	TransactTime int64  `json:"transactTime"`
}

// ==============TradeCancelOrders=================

type spot_cancelOrder struct {
	callAPI func(ctx context.Context, r *utils.Request, opts ...utils.RequestOption) (data []byte, header *http.Header, err error)
	convert spot_converts

	symbol  *string
	orderID *string
}

func (s *spot_cancelOrder) Symbol(symbol string) *spot_cancelOrder {
	s.symbol = &symbol
	return s
}

func (s *spot_cancelOrder) OrderID(orderID string) *spot_cancelOrder {
	s.orderID = &orderID
	return s
}

func (s *spot_cancelOrder) Do(ctx context.Context, opts ...utils.RequestOption) (res []entity.PlaceOrder, err error) {
	r := &utils.Request{
		Method:   http.MethodDelete,
		Endpoint: "/api/v3/order",
		SecType:  utils.SecTypeSigned,
	}

	m := utils.Params{}

	if s.symbol != nil {
		m["symbol"] = *s.symbol
	}

	if s.orderID != nil {
		m["orderId"] = *s.orderID
	}

	r.SetFormParams(m)

	data, _, err := s.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}

	answ := placeOrder_Response{}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}

	return s.convert.convertPlaceOrder(answ), nil
}

// ==============spot_getOrderList=================
type spot_getOrderList struct {
	callAPI func(ctx context.Context, r *utils.Request, opts ...utils.RequestOption) (data []byte, header *http.Header, err error)
	convert spot_converts

	symbol *string
}

func (s *spot_getOrderList) Symbol(symbol string) *spot_getOrderList {
	s.symbol = &symbol
	return s
}

func (s *spot_getOrderList) Do(ctx context.Context, opts ...utils.RequestOption) (res []entity.OrdersPendingList, err error) {
	r := &utils.Request{
		Method:   http.MethodGet,
		Endpoint: "/api/v3/openOrders",
		SecType:  utils.SecTypeSigned,
	}

	m := utils.Params{}
	if s.symbol != nil {
		m["symbol"] = *s.symbol
	}
	r.SetParams(m)

	data, _, err := s.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}

	log.Println("=4c00f6=", string(data))
	answ := []spot_orderList{}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}

	return s.convert.convertOrderList(answ), nil
}

type spot_orderList struct {
	Symbol              string `json:"symbol"`
	OrderId             string `json:"orderId"`
	ClientOrderId       string `json:"clientOrderId"`
	Price               string `json:"price"`
	OrigQty             string `json:"origQty"`
	ExecutedQty         string `json:"executedQty"`
	CummulativeQuoteQty string `json:"cummulativeQuoteQty"`
	Status              string `json:"status"`
	Type                string `json:"type"`
	Side                string `json:"side"`
	StopPrice           string `json:"stopPrice"`
	Time                int64  `json:"time"`
	UpdateTime          int64  `json:"updateTime"`
	IsWorking           bool   `json:"isWorking"`
	OrigQuoteOrderQty   string `json:"origQuoteOrderQty"`
}
