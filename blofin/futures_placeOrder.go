package blofin

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
	reduceOnly    *bool
	tpTriggerPx   *string
	tpOrderPx     *string
	slTriggerPx   *string
	slOrderPx     *string

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

func (s *futures_placeOrder) ReduceOnly(reduceOnly bool) *futures_placeOrder {
	s.reduceOnly = &reduceOnly
	return s
}

func (s *futures_placeOrder) TpTriggerPrice(px string) *futures_placeOrder {
	s.tpTriggerPx = &px
	return s
}

func (s *futures_placeOrder) TpOrderPrice(px string) *futures_placeOrder {
	s.tpOrderPx = &px
	return s
}

func (s *futures_placeOrder) SlTriggerPrice(px string) *futures_placeOrder {
	s.slTriggerPx = &px
	return s
}

func (s *futures_placeOrder) SlOrderPrice(px string) *futures_placeOrder {
	s.slOrderPx = &px
	return s
}

func (s *futures_placeOrder) Do(ctx context.Context, opts ...utils.RequestOption) (res []entity.PlaceOrder, err error) {
	// --- базовые проверки обязательных полей ---
	if s.symbol == nil || *s.symbol == "" {
		return res, errors.New("symbol (instId) is required")
	}
	if s.marginMode == nil {
		return res, errors.New("marginMode is required")
	}
	if s.positionSide == nil {
		// на Blofin positionSide обязателен даже в one-way, по доке
		return res, errors.New("positionSide is required")
	}
	if s.side == nil {
		return res, errors.New("side is required")
	}
	if s.orderType == nil {
		return res, errors.New("orderType is required")
	}
	if s.size == nil || *s.size == "" {
		return res, errors.New("size is required")
	}

	ordTypeStr := strings.ToLower(string(*s.orderType))
	if ordTypeStr != "market" && s.price == nil {
		return res, errors.New("price is required for non-market orders")
	}

	r := &utils.Request{
		Method:   http.MethodPost,
		Endpoint: "/api/v1/trade/order",
		SecType:  utils.SecTypeSigned,
	}

	m := utils.Params{}

	// instId
	m["instId"] = *s.symbol

	// marginMode: cross / isolated
	switch *s.marginMode {
	case entity.MarginModeTypeCross:
		m["marginMode"] = "cross"
	case entity.MarginModeTypeIsolated:
		m["marginMode"] = "isolated"
	default:
		return res, errors.New("unsupported marginMode: " + string(*s.marginMode))
	}

	// positionSide: net / long / short
	if s.positionSide != nil {
		ps := strings.ToLower(string(*s.positionSide))
		// как и в setLeverage: ONE_WAY/BOTH -> net
		if ps == "both" || ps == "one_way" {
			ps = "net"
		}
		m["positionSide"] = ps
	}

	// side: buy / sell
	if s.side != nil {
		m["side"] = strings.ToLower(string(*s.side))
	}

	// orderType
	m["orderType"] = ordTypeStr

	// price (кроме market)
	if s.price != nil && ordTypeStr != "market" {
		m["price"] = *s.price
	}

	// size (контракты)
	m["size"] = *s.size

	// reduceOnly
	if s.reduceOnly != nil {
		if *s.reduceOnly {
			m["reduceOnly"] = "true"
		} else {
			m["reduceOnly"] = "false"
		}
	}

	// clientOrderId
	if s.clientOrderID != nil {
		m["clientOrderId"] = *s.clientOrderID
	}

	// TP/SL – по доке: tpTriggerPrice/tpOrderPrice/slTriggerPrice/slOrderPrice
	if s.tpTriggerPx != nil && s.tpOrderPx != nil {
		m["tpTriggerPrice"] = *s.tpTriggerPx
		m["tpOrderPrice"] = *s.tpOrderPx
	}
	if s.slTriggerPx != nil && s.slOrderPx != nil {
		m["slTriggerPrice"] = *s.slTriggerPx
		m["slOrderPrice"] = *s.slOrderPx
	}

	r.SetFormParams(m)

	data, _, err := s.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}

	var answ struct {
		Result []placeOrder_Response `json:"data"`
	}

	if err = json.Unmarshal(data, &answ); err != nil {
		return res, err
	}

	if len(answ.Result) == 0 {
		return res, errors.New("Zero Answer")
	}

	// проверяем внутренний code по каждому ордеру
	for _, item := range answ.Result {
		if item.Code != "0" {
			// если есть текст — вернём его
			if item.Msg != "" {
				return res, errors.New(item.Msg)
			}
			return res, errors.New("place order failed with code " + item.Code)
		}
	}

	return s.convert.convertPlaceOrder(answ.Result), nil
}

// структура под ответ Blofin Place Order
type placeOrder_Response struct {
	OrderId       string `json:"orderId"`
	ClientOrderId string `json:"clientOrderId"`
	Msg           string `json:"msg"`
	Code          string `json:"code"`
}
