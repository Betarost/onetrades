package bitget

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

	symbol    *string
	orderType *entity.OrderType
	limit     *int
}

func (s *futures_getOrderList) Symbol(symbol string) *futures_getOrderList {
	s.symbol = &symbol
	return s
}

func (s *futures_getOrderList) OrderType(orderType entity.OrderType) *futures_getOrderList {
	s.orderType = &orderType
	return s
}

func (s *futures_getOrderList) Limit(limit int) *futures_getOrderList {
	s.limit = &limit
	return s
}

func (s *futures_getOrderList) Do(ctx context.Context, opts ...utils.RequestOption) (res []entity.Futures_OrdersList, err error) {
	r := &utils.Request{
		Method:   http.MethodGet,
		Endpoint: "/api/v2/mix/order/orders-pending",
		SecType:  utils.SecTypeSigned,
	}

	m := utils.Params{"productType": "USDT-FUTURES"}

	if s.symbol != nil {
		m["symbol"] = *s.symbol
	}

	r.SetParams(m)

	data, _, err := s.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}
	var answ struct {
		Result futures_orderList `json:"data"`
	}
	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}
	res = s.convert.convertOrderList(answ.Result)
	return res, nil
}

type futures_orderList struct {
	Orders []struct {
		Symbol        string `json:"symbol"`
		OrderId       string `json:"orderId"`
		ClientOrderId string `json:"clientOid"`
		// PositionID    int64  `json:"positionID"`
		Side         string `json:"side"`
		PositionSide string `json:"posSide"`
		Type         string `json:"orderType"`
		Size         string `json:"size"`
		BaseVolume   string `json:"baseVolume"`
		Price        string `json:"price"`
		AvgPrice     string `json:"priceAvg"`
		Leverage     string `json:"leverage"`
		MarginMode   string `json:"marginMode"`
		TradeSide    string `json:"tradeSide"`
		PosMode      string `json:"posMode"`
		Status       string `json:"status"`
		Time         string `json:"cTime"`
		UpdateTime   string `json:"uTime"`
	} `json:"entrustedList"`
}
