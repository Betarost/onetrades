package gateio

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/Betarost/onetrades/entity"
	"github.com/Betarost/onetrades/utils"
)

// ==============GetInstrumentsInfo=================
type futures_getInstrumentsInfo struct {
	callAPI func(ctx context.Context, r *utils.Request, opts ...utils.RequestOption) (data []byte, header *http.Header, err error)
	convert futures_converts

	symbol *string
	settle *string
}

func (s *futures_getInstrumentsInfo) Symbol(symbol string) *futures_getInstrumentsInfo {
	s.symbol = &symbol
	return s
}

func (s *futures_getInstrumentsInfo) Settle(settle string) *futures_getInstrumentsInfo {
	s.settle = &settle
	return s
}

func (s *futures_getInstrumentsInfo) Do(ctx context.Context, opts ...utils.RequestOption) (res []entity.InstrumentsInfo, err error) {
	r := &utils.Request{
		Method:   http.MethodGet,
		Endpoint: "/api/v4/futures/{settle}/contracts",
		SecType:  utils.SecTypeNone,
	}

	settleDefault := "usdt"

	if s.settle == nil {
		s.settle = &settleDefault
	}

	r.Endpoint = strings.Replace(r.Endpoint, "{settle}", *s.settle, 1)

	if s.symbol != nil {
		r.Endpoint = fmt.Sprintf("%s/%s", r.Endpoint, *s.symbol)
	}

	data, _, err := s.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}
	log.Println("=8099fe=", string(data))
	if s.symbol != nil {
		answ := futures_instrumentsInfo{}

		err = json.Unmarshal(data, &answ)
		if err != nil {
			return res, err
		}

		return s.convert.convertInstrumentsInfo([]futures_instrumentsInfo{answ}), nil

	} else {
		answ := []futures_instrumentsInfo{}

		err = json.Unmarshal(data, &answ)
		if err != nil {
			return res, err
		}

		return s.convert.convertInstrumentsInfo(answ), nil
	}

}

type futures_instrumentsInfo struct {
	Name                string `json:"name"`
	Type                string `json:"type"`
	Quanto_multiplier   string `json:"quanto_multiplier"`
	Leverage_min        string `json:"leverage_min"`
	Leverage_max        string `json:"leverage_max"`
	Maintenance_rate    string `json:"maintenance_rate"`
	Mark_type           string `json:"mark_type"`
	Mark_price          string `json:"mark_price"`
	Index_price         string `json:"index_price"`
	Last_price          string `json:"last_price"`
	Maker_fee_rate      string `json:"maker_fee_rate"`
	Taker_fee_rate      string `json:"taker_fee_rate"`
	Order_price_round   string `json:"order_price_round"`
	Mark_price_round    string `json:"mark_price_round"`
	Funding_rate        string `json:"funding_rate"`
	Order_size_min      int64  `json:"order_size_min"`
	Order_size_max      int64  `json:"order_size_max"`
	Trade_id            int64  `json:"trade_id"`
	Trade_size          int64  `json:"trade_size"`
	Position_size       int64  `json:"position_size"`
	Order_price_deviate string `json:"order_price_deviate"`
	In_delisting        bool   `json:"in_delisting"`
	Orders_limit        int64  `json:"orders_limit"`
}

// ==============SetPositionMode=================
type futures_setPositionMode struct {
	callAPI func(ctx context.Context, r *utils.Request, opts ...utils.RequestOption) (data []byte, header *http.Header, err error)
	convert futures_converts

	settle *string
	mode   *entity.PositionModeType
}

func (s *futures_setPositionMode) Settle(settle string) *futures_setPositionMode {
	s.settle = &settle
	return s
}

func (s *futures_setPositionMode) Mode(mode entity.PositionModeType) *futures_setPositionMode {
	s.mode = &mode
	return s
}

func (s *futures_setPositionMode) Do(ctx context.Context, opts ...utils.RequestOption) (res bool, err error) {
	r := &utils.Request{
		Method:   http.MethodPost,
		Endpoint: "/api/v4/futures/{settle}/dual_mode",
		SecType:  utils.SecTypeSigned,
	}

	m := utils.Params{}

	settleDefault := "usdt"

	if s.settle == nil {
		s.settle = &settleDefault
	}

	r.Endpoint = strings.Replace(r.Endpoint, "{settle}", *s.settle, 1)

	if s.mode != nil {
		if *s.mode == entity.PositionModeTypeHedge {
			m["dual_mode"] = true
		} else if *s.mode == entity.PositionModeTypeOneWay {
			m["dual_mode"] = false
		}
	}
	r.SetParams(m)

	data, _, err := s.callAPI(ctx, r, opts...)
	if err != nil {
		return false, err
	}

	log.Println("=d6461a=", string(data))
	answ := positionMode{}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return false, err
	}

	return true, nil
}

type positionMode struct {
	In_dual_mode bool `json:"in_dual_mode"`
}
