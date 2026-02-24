package gateio

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
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
	hedgeMode     *bool
	settle        *string
	reduce        *bool
	tpOrder       *bool
	slOrder       *bool
}

func (s *futures_placeOrder) Reduce(reduce bool) *futures_placeOrder {
	s.reduce = &reduce
	return s
}

func (s *futures_placeOrder) TpOrder(v bool) *futures_placeOrder {
	s.tpOrder = &v
	return s
}

func (s *futures_placeOrder) SlOrder(v bool) *futures_placeOrder {
	s.slOrder = &v
	return s
}

func (s *futures_placeOrder) TradeMode(tradeMode entity.MarginModeType) *futures_placeOrder {
	s.tradeMode = &tradeMode
	return s
}
func (s *futures_placeOrder) HedgeMode(hedgeMode bool) *futures_placeOrder {
	s.hedgeMode = &hedgeMode
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
		Endpoint: "/api/v4/futures/{settle}/orders",
		SecType:  utils.SecTypeSigned,
	}

	settleDefault := "usdt"

	if s.settle == nil {
		s.settle = &settleDefault
	}

	r.Endpoint = strings.Replace(r.Endpoint, "{settle}", *s.settle, 1)

	isTP := s.tpOrder != nil && *s.tpOrder
	isSL := s.slOrder != nil && *s.slOrder

	// минимальная валидация: нельзя одновременно TP и SL
	if isTP && isSL {
		return res, errors.New("gateio futures_placeOrder: TpOrder and SlOrder cannot both be true")
	}

	if isTP || isSL {
		// price-triggered order endpoint
		r.Endpoint = "/api/v4/futures/{settle}/price_orders"
		r.Endpoint = strings.Replace(r.Endpoint, "{settle}", *s.settle, 1)

		// -------- initial (что выставим, когда триггер сработает) --------
		initial := utils.Params{}

		if s.symbol != nil {
			initial["contract"] = *s.symbol
		}

		// market close by default
		initial["price"] = "0"
		initial["tif"] = "ioc"

		// close/reduce flags (пусть биржа валидирует если что)
		initial["close"] = true
		initial["reduce_only"] = true

		// text (clientOrderID)
		if s.clientOrderID != nil {
			initial["text"] = *s.clientOrderID
		} else {
			// если не задан — можно не ставить вовсе, но оставим "api"
			initial["text"] = "api"
		}

		// size:
		// - если size не задан => size=0 (закрыть позицию целиком по доке)
		// - если size задан => используем его, а знак определяем по side (если он задан)
		if s.size != nil && *s.size != "" {
			if s.side != nil && *s.side == entity.SideTypeSell {
				initial["size"] = fmt.Sprintf("-%s", *s.size)
			} else if s.side != nil && *s.side == entity.SideTypeBuy {
				initial["size"] = *s.size
			} else {
				// side не задан — отправим как есть, биржа сама скажет
				initial["size"] = *s.size
			}
		} else {
			initial["size"] = 0
		}

		// -------- trigger (когда сработает) --------
		trigger := utils.Params{
			"strategy_type": 0,
			"price_type":    0,     // last price
			"expiration":    86400, // 24h
		}

		if s.price != nil {
			trigger["price"] = *s.price
		}

		// rule (если side известен — подставим разумно; иначе пусть биржа валидирует)
		// rule: 1 => >= price, 2 => <= price
		if s.side != nil {
			// трактуем side как "закрывающий" (как у Binance/Bybit в вашей логике)
			// SELL обычно закрывает LONG, BUY закрывает SHORT
			if *s.side == entity.SideTypeSell {
				// LONG close
				if isTP {
					trigger["rule"] = 1 // >= trigger
				} else {
					trigger["rule"] = 2 // <= trigger
				}
			} else if *s.side == entity.SideTypeBuy {
				// SHORT close
				if isTP {
					trigger["rule"] = 2 // <= trigger
				} else {
					trigger["rule"] = 1 // >= trigger
				}
			}
		}

		body := utils.Params{
			"initial": initial,
			"trigger": trigger,
		}

		r.SetFormParams(body)

		data, _, err := s.callAPI(ctx, r, opts...)
		if err != nil {
			return res, err
		}

		// ответ: { "id": 1432329 }
		var answ struct {
			ID int64 `json:"id"`
		}
		if err := json.Unmarshal(data, &answ); err != nil {
			return res, err
		}

		// адаптируем под текущий ответ create order
		out := futures_placeOrder_Response{
			Contract:    "",
			ID:          answ.ID,
			Text:        "",
			Create_time: 0,
		}
		if s.symbol != nil {
			out.Contract = *s.symbol
		}
		if s.clientOrderID != nil {
			out.Text = *s.clientOrderID
		}

		return s.convert.convertPlaceOrder(out), nil
	}
	m := utils.Params{}

	if s.price != nil {
		m["price"] = *s.price
	}

	if s.symbol != nil {
		m["contract"] = *s.symbol
	}

	if s.size != nil {
		m["size"] = *s.size
		// m["size"] = utils.StringToInt64(*s.size)
	}

	if s.side != nil && s.size != nil {
		if *s.side == entity.SideTypeSell {
			m["size"] = fmt.Sprintf("-%s", *s.size)
			// m["size"] = 0 - utils.StringToInt64(*s.size)
		}
	}

	if s.orderType != nil {
		if *s.orderType == entity.OrderTypeMarket {
			m["price"] = "0"
			m["tif"] = "ioc"
		}
	}

	if s.clientOrderID != nil {
		m["text"] = *s.clientOrderID
	}

	if s.hedgeMode != nil && s.positionSide != nil && s.side != nil {
		if *s.hedgeMode && ((strings.ToUpper(string(*s.positionSide)) == "LONG" && strings.ToUpper(string(*s.side)) == "SELL") || (strings.ToUpper(string(*s.positionSide)) == "SHORT" && strings.ToUpper(string(*s.side)) == "BUY")) {
			m["reduce_only"] = true
		}
	}

	if s.reduce != nil && *s.reduce == true {
		m["reduce_only"] = true
	}

	// if s.tradeMode != nil {
	// 	if *s.tradeMode == entity.MarginModeTypeCross {
	// 		m["tdMode"] = "cross"
	// 	} else if *s.tradeMode == entity.MarginModeTypeIsolated {
	// 		m["tdMode"] = "isolated"
	// 	}
	// }

	// if s.orderType != nil {
	// 	m["ordType"] = strings.ToLower(string(*s.orderType))
	// }

	// if s.positionSide != nil {
	// 	m["posSide"] = strings.ToLower(string(*s.positionSide))
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

	answ := futures_placeOrder_Response{}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}

	return s.convert.convertPlaceOrder(answ), nil
}

type futures_placeOrder_Response struct {
	Contract    string  `json:"contract"`
	ID          int64   `json:"id"`
	Text        string  `json:"text"`
	Create_time float64 `json:"create_time"`
}
