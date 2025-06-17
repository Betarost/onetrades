package bybit

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
		Endpoint: "/v5/market/instruments-info",
		SecType:  utils.SecTypeNone,
	}

	m := utils.Params{
		"category": "spot",
	}

	if s.symbol != nil {
		m["symbol"] = *s.symbol
	}

	r.SetParams(m)

	data, _, err := s.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}

	var answ struct {
		RetCode int64  `json:"retCode"`
		RetMsg  string `json:"retMsg"`
		Result  struct {
			List []instrumentsInfo `json:"list"`
		} `json:"result"`
		Time int64 `json:"time"`
	}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}

	if answ.RetCode != 0 {
		return res, errors.New(answ.RetMsg)
	}
	return convertInstrumentsInfo(answ.Result.List, "SPOT"), nil
}

type instrumentsInfo struct {
	Symbol        string `json:"symbol"`
	BaseCoin      string `json:"baseCoin"`
	QuoteCoin     string `json:"quoteCoin"`
	Status        string `json:"status"`
	LotSizeFilter struct {
		BasePrecision  string `json:"basePrecision"`
		QuotePrecision string `json:"quotePrecision"`
		MinOrderQty    string `json:"minOrderQty"`
		MaxOrderQty    string `json:"maxOrderQty"`
		MinOrderAmt    string `json:"minOrderAmt"`
		MaxOrderAmt    string `json:"maxOrderAmt"`
	} `json:"lotSizeFilter"`
	PriceFilter struct {
		TickSize string `json:"tickSize"`
	} `json:"priceFilter"`
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
	positionSide  *entity.PositionSideType
	tradeMode     *entity.MarginModeType
	tpPrice       *string
	slPrice       *string
}

func (s *placeOrder) TradeMode(tradeMode entity.MarginModeType) *placeOrder {
	s.tradeMode = &tradeMode
	return s
}

func (s *placeOrder) SlPrice(slPrice string) *placeOrder {
	s.slPrice = &slPrice
	return s
}

func (s *placeOrder) TpPrice(tpPrice string) *placeOrder {
	s.tpPrice = &tpPrice
	return s
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

func (s *placeOrder) PositionSide(positionSide entity.PositionSideType) *placeOrder {
	s.positionSide = &positionSide
	return s
}

func (s *placeOrder) Do(ctx context.Context, opts ...utils.RequestOption) (res []entity.PlaceOrder, err error) {
	r := &utils.Request{
		Method:   http.MethodPost,
		Endpoint: "/v5/order/create",
		SecType:  utils.SecTypeSigned,
	}

	m := utils.Params{
		"category": "spot",
	}

	if s.symbol != nil {
		m["symbol"] = *s.symbol
	}

	if s.side != nil {
		m["side"] = strings.ToLower(string(*s.side))
	}

	if s.orderType != nil {
		m["orderType"] = strings.ToLower(string(*s.orderType))
	}

	if s.size != nil {
		m["qty"] = *s.size
	}

	// if s.tradeMode != nil {
	// 	if *s.tradeMode == entity.MarginModeTypeCross {
	// 		m["tdMode"] = "cross"
	// 	} else if *s.tradeMode == entity.MarginModeTypeIsolated {
	// 		m["tdMode"] = "isolated"
	// 	}
	// }

	// if s.price != nil {
	// 	m["px"] = *s.price
	// }

	// if s.clientOrderID != nil {
	// 	m["clOrdId"] = *s.clientOrderID
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

	log.Println("=c2eced=", string(data))

	var answ struct {
		RetCode int64               `json:"retCode"`
		RetMsg  string              `json:"retMsg"`
		Result  placeOrder_Response `json:"result"`
		Time    int64               `json:"time"`
	}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}

	if answ.RetCode != 0 {
		return res, errors.New(answ.RetMsg)
	}

	return convertPlaceOrder(answ.Result), nil
}

type placeOrder_Response struct {
	OrderId     string `json:"orderId"`
	OrderLinkId string `json:"orderLinkId"`
}
