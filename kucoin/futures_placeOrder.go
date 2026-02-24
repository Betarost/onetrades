package kucoin

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
	leverage      *string

	positionSide *entity.PositionSideType
	marginMode   *entity.MarginModeType

	reduce  *bool
	tpOrder *bool
	slOrder *bool
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

func (s *futures_placeOrder) Leverage(leverage string) *futures_placeOrder {
	s.leverage = &leverage
	return s
}

func (s *futures_placeOrder) PositionSide(positionSide entity.PositionSideType) *futures_placeOrder {
	s.positionSide = &positionSide
	return s
}

func (s *futures_placeOrder) MarginMode(marginMode entity.MarginModeType) *futures_placeOrder {
	s.marginMode = &marginMode
	return s
}

func (s *futures_placeOrder) Do(ctx context.Context, opts ...utils.RequestOption) (res []entity.PlaceOrder, err error) {
	r := &utils.Request{
		Method:   http.MethodPost,
		Endpoint: "/api/v1/orders",
		SecType:  utils.SecTypeSigned,
	}

	m := utils.Params{}

	// --- TP / SL separate order for existing futures position (KuCoin) ---
	isTP := s.tpOrder != nil && *s.tpOrder
	isSL := s.slOrder != nil && *s.slOrder

	// минимальная валидация: нельзя одновременно TP и SL
	if isTP && isSL {
		return res, errors.New("kucoin futures_placeOrder: TpOrder and SlOrder cannot both be true")
	}

	if isTP || isSL {
		// KuCoin отдельный endpoint для TPSL
		r.Endpoint = "/api/v1/st-orders"

		// Базовые поля такие же как у place order по доке :contentReference[oaicite:1]{index=1}
		if s.symbol != nil {
			m["symbol"] = *s.symbol
		}
		if s.side != nil {
			m["side"] = strings.ToLower(string(*s.side))
		}
		if s.size != nil {
			m["size"] = *s.size
		}
		if s.clientOrderID != nil {
			m["clientOid"] = *s.clientOrderID
		}
		if s.leverage != nil {
			m["leverage"] = *s.leverage
		}
		if s.marginMode != nil {
			switch *s.marginMode {
			case entity.MarginModeTypeCross:
				m["marginMode"] = "CROSS"
			case entity.MarginModeTypeIsolated:
				m["marginMode"] = "ISOLATED"
			}
		}
		if s.positionSide != nil {
			m["positionSide"] = strings.ToUpper(string(*s.positionSide))
		}

		// Для нашего унифицированного TP/SL: делаем исполнение MARKET, а s.price используем как trigger.
		m["type"] = "market"

		// stopPriceType: "TP" (обычно trade/last price) — как в примере доки :contentReference[oaicite:2]{index=2}
		m["stopPriceType"] = "TP"

		// trigger price:
		// KuCoin ожидает triggerStopUpPrice и/или triggerStopDownPrice. В примере присутствуют оба. :contentReference[oaicite:3]{index=3}
		// Мы ставим один (по направлению), чтобы было однозначно.
		if s.price != nil {
			// Логика направления, если side указан:
			// SELL (обычно закрытие LONG): TP вверх, SL вниз
			// BUY  (обычно закрытие SHORT): TP вниз, SL вверх
			if s.side != nil {
				if *s.side == entity.SideTypeSell {
					if isTP {
						m["triggerStopUpPrice"] = *s.price
					} else {
						m["triggerStopDownPrice"] = *s.price
					}
				} else if *s.side == entity.SideTypeBuy {
					if isTP {
						m["triggerStopDownPrice"] = *s.price
					} else {
						m["triggerStopUpPrice"] = *s.price
					}
				}
			} else {
				// если side не задан — пусть биржа вернёт ошибку/объяснение;
				// но чтобы запрос был “какой-то”, ставим TP вверх, SL вниз
				if isTP {
					m["triggerStopUpPrice"] = *s.price
				} else {
					m["triggerStopDownPrice"] = *s.price
				}
			}
		}

		// reduceOnly: для TP/SL логично true (чтобы не открывать), но без жёсткой валидации:
		if s.reduce != nil && *s.reduce == true {
			m["reduceOnly"] = true
		} else {
			m["reduceOnly"] = true
		}

		r.SetFormParams(m)

		data, _, err := s.callAPI(ctx, r, opts...)
		if err != nil {
			return res, err
		}

		var answ struct {
			Result placeOrder_Response `json:"data"`
		}

		err = json.Unmarshal(data, &answ)
		if err != nil {
			return res, err
		}

		return s.convert.convertPlaceOrder(answ.Result), nil
	}
	// --- end TP / SL branch ---

	if s.symbol != nil {
		m["symbol"] = *s.symbol
	}

	if s.side != nil {
		m["side"] = strings.ToLower(string(*s.side))
	}

	if s.size != nil {
		m["size"] = *s.size
	}

	if s.orderType != nil {
		m["type"] = strings.ToLower(string(*s.orderType))
	}

	if s.price != nil {
		m["price"] = *s.price
	}

	if s.clientOrderID != nil {
		m["clientOid"] = *s.clientOrderID
	}

	if s.leverage != nil {
		m["leverage"] = *s.leverage
	}

	if s.marginMode != nil {
		switch *s.marginMode {
		case entity.MarginModeTypeCross:
			m["marginMode"] = "CROSS"
		case entity.MarginModeTypeIsolated:
			m["marginMode"] = "ISOLATED"
		}
	}

	if s.positionSide != nil {
		m["positionSide"] = strings.ToUpper(string(*s.positionSide))
	}

	if s.reduce != nil && *s.reduce == true {
		m["reduceOnly"] = true
	}

	r.SetFormParams(m)

	data, _, err := s.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}
	var answ struct {
		Result placeOrder_Response `json:"data"`
	}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}

	return s.convert.convertPlaceOrder(answ.Result), nil
}
