package bitget

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Betarost/onetrades/entity"
	"github.com/Betarost/onetrades/utils"
)

// ==============futures_placeOrder=================
type futures_placeOrder struct {
	callAPI func(ctx context.Context, r *utils.Request, opts ...utils.RequestOption) (data []byte, header *http.Header, err error)
	convert futures_converts

	symbol        *string
	side          *entity.SideType
	size          *string
	price         *string
	orderType     *entity.OrderType
	clientOrderID *string
	positionSide  *entity.PositionSideType
	marginMode    *entity.MarginModeType
	hedgeMode     *bool
	tpPrice       *string
	slPrice       *string
}

func (s *futures_placeOrder) MarginMode(marginMode entity.MarginModeType) *futures_placeOrder {
	s.marginMode = &marginMode
	return s
}

func (s *futures_placeOrder) HedgeMode(hedgeMode bool) *futures_placeOrder {
	s.hedgeMode = &hedgeMode
	return s
}

func (s *futures_placeOrder) SlPrice(slPrice string) *futures_placeOrder {
	s.slPrice = &slPrice
	return s
}

func (s *futures_placeOrder) TpPrice(tpPrice string) *futures_placeOrder {
	s.tpPrice = &tpPrice
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
		Endpoint: "/api/v2/mix/order/place-order",
		SecType:  utils.SecTypeSigned,
	}

	m := utils.Params{"productType": "USDT-FUTURES", "marginCoin": "USDT"}

	if s.symbol != nil {
		m["symbol"] = *s.symbol
	}

	if s.size != nil {
		m["size"] = *s.size
	}

	if s.price != nil {
		m["price"] = *s.price
	}

	if s.side != nil {
		m["side"] = strings.ToLower(string(*s.side))
	}

	if s.marginMode != nil {
		switch *s.marginMode {
		case entity.MarginModeTypeCross:
			m["marginMode"] = "crossed"
		case entity.MarginModeTypeIsolated:
			m["marginMode"] = "isolated"
		}
	}

	if s.orderType != nil {
		m["orderType"] = strings.ToLower(string(*s.orderType))
		if *s.orderType == entity.OrderTypeLimit {

		}
	}

	if s.clientOrderID != nil {
		m["clientOid"] = *s.clientOrderID
	}

	if s.hedgeMode != nil && *s.hedgeMode {
		if s.positionSide != nil {
			switch *s.positionSide {
			case entity.PositionSideTypeLong:
				if s.side != nil && *s.side == entity.SideTypeBuy {
					m["tradeSide"] = "open"
				} else if s.side != nil && *s.side == entity.SideTypeSell {
					m["tradeSide"] = "close"
					m["side"] = strings.ToLower(string(entity.SideTypeBuy))
				}
			case entity.PositionSideTypeShort:
				if s.side != nil && *s.side == entity.SideTypeSell {
					m["tradeSide"] = "open"
				} else if s.side != nil && *s.side == entity.SideTypeBuy {
					m["tradeSide"] = "close"
					m["side"] = strings.ToLower(string(entity.SideTypeSell))
				}
			}
		}
	}

	r.SetFormParams(m)

	data, _, err := s.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}

	var answ struct {
		Result futures_placeOrder_Response `json:"data"`
	}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}

	return s.convert.convertPlaceOrder(answ.Result), nil
}

type futures_placeOrder_Response struct {
	Symbol    string `json:"symbol"`
	ClientOid string `json:"clientOid"`
	OrderId   string `json:"orderId"`
}
