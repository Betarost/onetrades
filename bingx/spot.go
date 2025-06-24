package bingx

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

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
		Endpoint: "/openApi/spot/v1/common/symbols",
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

	var answ struct {
		Result spot_instrumentsInfo `json:"data"`
	}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}

	return s.convert.convertInstrumentsInfo(answ.Result), nil
}

type spot_instrumentsInfo struct {
	Symbols []struct {
		Symbol      string  `json:"symbol"`
		MinQty      float64 `json:"minQty"`
		MaxQty      float64 `json:"maxQty"`
		MinNotional float64 `json:"minNotional"`
		MaxNotional float64 `json:"maxNotional"`
		Status      int64   `json:"status"`
		TickSize    float64 `json:"tickSize"`
		StepSize    float64 `json:"stepSize"`
	} `json:"symbols"`
}

// ==============placeOrder=================
type spot_placeOrder struct {
	callAPI func(ctx context.Context, r *utils.Request, opts ...utils.RequestOption) (data []byte, header *http.Header, err error)
	convert spot_converts

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

func (s *spot_placeOrder) Symbol(symbol string) *spot_placeOrder {
	s.symbol = &symbol
	return s
}

func (s *spot_placeOrder) TradeMode(tradeMode entity.MarginModeType) *spot_placeOrder {
	s.tradeMode = &tradeMode
	return s
}

func (s *spot_placeOrder) SlPrice(slPrice string) *spot_placeOrder {
	s.slPrice = &slPrice
	return s
}

func (s *spot_placeOrder) TpPrice(tpPrice string) *spot_placeOrder {
	s.tpPrice = &tpPrice
	return s
}

func (s *spot_placeOrder) Side(side entity.SideType) *spot_placeOrder {
	s.side = &side
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

func (s *spot_placeOrder) OrderType(orderType entity.OrderType) *spot_placeOrder {
	s.orderType = &orderType
	return s
}

func (s *spot_placeOrder) ClientOrderID(clientOrderID string) *spot_placeOrder {
	s.clientOrderID = &clientOrderID
	return s
}

func (s *spot_placeOrder) PositionSide(positionSide entity.PositionSideType) *spot_placeOrder {
	s.positionSide = &positionSide
	return s
}

func (s *spot_placeOrder) Do(ctx context.Context, opts ...utils.RequestOption) (res []entity.PlaceOrder, err error) {
	r := &utils.Request{
		Method:   http.MethodPost,
		Endpoint: "/openApi/spot/v1/trade/order",
		SecType:  utils.SecTypeSigned,
	}

	m := utils.Params{}

	if s.symbol != nil {
		m["symbol"] = *s.symbol
	}

	if s.tradeMode != nil {
		if *s.tradeMode == entity.MarginModeTypeCross {
			m["tdMode"] = "cross"
		} else if *s.tradeMode == entity.MarginModeTypeIsolated {
			m["tdMode"] = "isolated"
		}
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

	// if s.tpPrice != nil || s.slPrice != nil {
	// 	attachAlgoOrds := []orderList_attachAlgoOrds{{}}
	// 	if s.tpPrice != nil {
	// 		attachAlgoOrds[0].TpTriggerPx = *s.tpPrice
	// 		attachAlgoOrds[0].TpOrdPx = "-1"
	// 	}

	// 	if s.slPrice != nil {
	// 		attachAlgoOrds[0].SlTriggerPx = *s.slPrice
	// 		attachAlgoOrds[0].SlOrdPx = "-1"
	// 	}
	// 	j, err := json.Marshal(attachAlgoOrds)
	// 	if err != nil {
	// 		return res, err
	// 	}

	// 	m["attachAlgoOrds"] = string(j)
	// }

	r.SetFormParams(m)

	data, _, err := s.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}

	log.Println("=e0cd43=", string(data))

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
	return
	// return convertPlaceOrder(answ.Result), nil
}

type placeOrder_Response struct {
	ClOrdId string `json:"clOrdId"`
	OrdId   string `json:"ordId"`
	Tag     string `json:"tag"`
	Ts      string `json:"ts"`
	SCode   string `json:"sCode"`
	SMsg    string `json:"sMsg"`
}
