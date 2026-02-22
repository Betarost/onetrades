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

type futures_placeOrder struct {
	callAPI  func(ctx context.Context, r *utils.Request, opts ...utils.RequestOption) (data []byte, header *http.Header, err error)
	brokerID string

	convert futures_converts

	symbol        *string
	side          *entity.SideType
	size          *string
	price         *string
	orderType     *entity.OrderType
	clientOrderID *string
	positionSide  *entity.PositionSideType
	marginMode    *entity.MarginModeType
	// tpPrice       *string
	// slPrice       *string

	reduce *bool
}

func (s *futures_placeOrder) Reduce(reduce bool) *futures_placeOrder {
	s.reduce = &reduce
	return s
}

func (s *futures_placeOrder) MarginMode(marginMode entity.MarginModeType) *futures_placeOrder {
	s.marginMode = &marginMode
	return s
}

func (s *futures_placeOrder) Symbol(symbol string) *futures_placeOrder {
	s.symbol = &symbol
	return s
}

func (s *futures_placeOrder) Side(side entity.SideType) *futures_placeOrder {
	s.side = &side
	return s
}

func (s *futures_placeOrder) Size(size string) *futures_placeOrder {
	s.size = &size
	return s
}

func (s *futures_placeOrder) Price(price string) *futures_placeOrder {
	s.price = &price
	return s
}

func (s *futures_placeOrder) OrderType(orderType entity.OrderType) *futures_placeOrder {
	s.orderType = &orderType
	return s
}

func (s *futures_placeOrder) ClientOrderID(clientOrderID string) *futures_placeOrder {
	s.clientOrderID = &clientOrderID
	return s
}

func (s *futures_placeOrder) PositionSide(positionSide entity.PositionSideType) *futures_placeOrder {
	s.positionSide = &positionSide
	return s
}

func (s *futures_placeOrder) Do(ctx context.Context, opts ...utils.RequestOption) (res []entity.PlaceOrder, err error) {
	r := &utils.Request{
		Method:   http.MethodPost,
		Endpoint: "/api/v5/trade/order",
		SecType:  utils.SecTypeSigned,
	}

	m := utils.Params{}

	if s.marginMode != nil {
		switch *s.marginMode {
		case entity.MarginModeTypeCross:
			m["tdMode"] = "cross"
		case entity.MarginModeTypeIsolated:
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

	if s.reduce != nil && *s.reduce == true {
		m["reduceOnly"] = true
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

	if s.brokerID != "" {
		m["tag"] = s.brokerID
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
	return s.convert.convertPlaceOrder(answ.Result), nil
}

// type futures_placeOrder_Response struct {
// 	ClOrdId string `json:"clOrdId"`
// 	OrdId   string `json:"ordId"`
// 	Tag     string `json:"tag"`
// 	Ts      string `json:"ts"`
// 	SCode   string `json:"sCode"`
// 	SMsg    string `json:"sMsg"`
// }
