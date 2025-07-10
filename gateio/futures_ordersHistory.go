package gateio

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/Betarost/onetrades/entity"
	"github.com/Betarost/onetrades/utils"
)

type futures_ordersHistory struct {
	callAPI func(ctx context.Context, r *utils.Request, opts ...utils.RequestOption) (data []byte, header *http.Header, err error)
	convert futures_converts

	settle    *string
	symbol    *string
	orderType *entity.OrderType
	limit     *int
}

func (s *futures_ordersHistory) Settle(settle string) *futures_ordersHistory {
	s.settle = &settle
	return s
}

func (s *futures_ordersHistory) Symbol(symbol string) *futures_ordersHistory {
	s.symbol = &symbol
	return s
}

func (s *futures_ordersHistory) OrderType(orderType entity.OrderType) *futures_ordersHistory {
	s.orderType = &orderType
	return s
}

func (s *futures_ordersHistory) Limit(limit int) *futures_ordersHistory {
	s.limit = &limit
	return s
}

func (s *futures_ordersHistory) Do(ctx context.Context, opts ...utils.RequestOption) (res []entity.Futures_OrdersHistory, err error) {
	r := &utils.Request{
		Method:   http.MethodGet,
		Endpoint: "/api/v4/futures/{settle}/orders",
		SecType:  utils.SecTypeSigned,
	}

	settleDefault := "usdt"

	if s.settle == nil {
		s.settle = &settleDefault
	}

	r.Endpoint = strings.Replace(r.Endpoint, "{settle}", *s.settle, 1)

	m := utils.Params{
		"status": "finished",
	}

	if s.symbol != nil {
		m["contract"] = *s.symbol
	}

	if s.limit != nil {
		m["limit"] = *s.limit
	}

	r.SetParams(m)

	data, _, err := s.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}
	log.Println("=a30286=", string(data))
	answ := []futures_ordersHistory_Response{}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}

	return s.convert.convertOrdersHistory(answ), nil
}

type futures_ordersHistory_Response struct {
	ID          int64   `json:"id"`
	Contract    string  `json:"contract"`
	Status      string  `json:"status"`
	Size        int64   `json:"size"`
	Left        int64   `json:"left"`
	Is_close    bool    `json:"is_close"`
	Text        string  `json:"text"`
	Price       string  `json:"price"`
	Fill_price  string  `json:"fill_price"`
	Create_time float64 `json:"create_time"`
	Update_time float64 `json:"update_time"`
}
