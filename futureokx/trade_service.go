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

// ==============TradePlaceOrder=================
type TradePlaceOrder struct {
	c            *Client
	symbol       *string
	positionSide *entity.PositionSideType
	side         *entity.SideType
	orderType    *entity.OrderType
	size         *float64
	price        *string
}

func (s *TradePlaceOrder) Size(size float64) *TradePlaceOrder {
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
