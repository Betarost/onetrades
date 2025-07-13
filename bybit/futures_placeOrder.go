package bybit

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Betarost/onetrades/entity"
	"github.com/Betarost/onetrades/utils"
)

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
	tradeMode     *entity.MarginModeType
	tpPrice       *string
	slPrice       *string
}

func (s *futures_placeOrder) TradeMode(tradeMode entity.MarginModeType) *futures_placeOrder {
	s.tradeMode = &tradeMode
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
		Endpoint: "/v5/order/create",
		SecType:  utils.SecTypeSigned,
	}

	m := utils.Params{
		"category": "linear",
	}

	if s.symbol != nil {
		m["symbol"] = *s.symbol
	}

	if s.side != nil {
		if *s.side == entity.SideTypeBuy {
			m["side"] = "Buy"
		} else if *s.side == entity.SideTypeSell {
			m["side"] = "Sell"
		}
	}

	if s.price != nil {
		m["price"] = *s.price
	}

	if s.orderType != nil {
		m["orderType"] = strings.ToLower(string(*s.orderType))
		if *s.orderType == entity.OrderTypeMarket {
			m["marketUnit"] = "baseCoin"
		}
	}

	if s.size != nil {
		m["qty"] = *s.size
	}

	if s.clientOrderID != nil {
		m["orderLinkId"] = *s.clientOrderID
	}

	// if s.positionSide != nil {
	// 	m["posSide"] = strings.ToLower(string(*s.positionSide))
	// } else {
	// 	m["positionIdx"] = 0
	// }

	// if s.tradeMode != nil {
	// 	if *s.tradeMode == entity.MarginModeTypeCross {
	// 		m["tdMode"] = "cross"
	// 	} else if *s.tradeMode == entity.MarginModeTypeIsolated {
	// 		m["tdMode"] = "isolated"
	// 	}
	// }

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

	var answ struct {
		Result futures_placeOrder_Response `json:"result"`
	}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}

	return s.convert.convertPlaceOrder(answ.Result), nil
}

type futures_placeOrder_Response struct {
	OrderId     string `json:"orderId"`
	OrderLinkId string `json:"orderLinkId"`
}
