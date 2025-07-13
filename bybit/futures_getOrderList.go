package bybit

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Betarost/onetrades/entity"
	"github.com/Betarost/onetrades/utils"
)

type futures_getOrderList struct {
	callAPI func(ctx context.Context, r *utils.Request, opts ...utils.RequestOption) (data []byte, header *http.Header, err error)
	convert futures_converts

	symbol *string
}

func (s *futures_getOrderList) Symbol(symbol string) *futures_getOrderList {
	s.symbol = &symbol
	return s
}

func (s *futures_getOrderList) Do(ctx context.Context, opts ...utils.RequestOption) (res []entity.Futures_OrdersList, err error) {
	r := &utils.Request{
		Method:   http.MethodGet,
		Endpoint: "/v5/order/realtime",
		SecType:  utils.SecTypeSigned,
	}

	m := utils.Params{
		"category": "linear",
		"limit":    50,
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
		Result struct {
			List []futures_orderList `json:"list"`
		} `json:"result"`
	}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}

	return s.convert.convertOrderList(answ.Result.List), nil
}

type futures_orderList struct {
	Symbol      string `json:"symbol"`
	OrderId     string `json:"orderId"`
	OrderLinkId string `json:"orderLinkId"`
	Side        string `json:"side"`
	PositionIdx int64  `json:"positionIdx"`
	Qty         string `json:"qty"`
	CumExecQty  string `json:"cumExecQty"`
	Price       string `json:"price"`
	OrderType   string `json:"orderType"`
	OrderStatus string `json:"orderStatus"`
	CreatedTime string `json:"createdTime"`
	UpdatedTime string `json:"updatedTime"`
}
