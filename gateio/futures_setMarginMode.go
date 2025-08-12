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

type futures_setMarginMode struct {
	callAPI func(ctx context.Context, r *utils.Request, opts ...utils.RequestOption) (data []byte, header *http.Header, err error)
	convert futures_converts

	symbol       *string
	leverage     *string
	positionSide *entity.PositionSideType
	marginMode   *entity.MarginModeType
}

func (s *futures_setMarginMode) Symbol(symbol string) *futures_setMarginMode {
	s.symbol = &symbol
	return s
}

func (s *futures_setMarginMode) Leverage(leverage string) *futures_setMarginMode {
	s.leverage = &leverage
	return s
}

func (s *futures_setMarginMode) MarginMode(marginMode entity.MarginModeType) *futures_setMarginMode {
	s.marginMode = &marginMode
	return s
}

func (s *futures_setMarginMode) PositionSide(positionSide entity.PositionSideType) *futures_setMarginMode {
	s.positionSide = &positionSide
	return s
}

func (s *futures_setMarginMode) Do(ctx context.Context, opts ...utils.RequestOption) (res entity.Futures_MarginMode, err error) {
	r := &utils.Request{
		Method:   http.MethodPost,
		Endpoint: "/api/v4/futures/{settle}/positions/cross_mode",
		SecType:  utils.SecTypeSigned,
	}

	m := utils.Params{}

	settleDefault := "usdt"

	r.Endpoint = strings.Replace(r.Endpoint, "{settle}", settleDefault, 1)

	if s.symbol != nil {
		m["contract"] = *s.symbol
	}

	if s.marginMode != nil {
		switch *s.marginMode {
		case entity.MarginModeTypeCross:
			m["mode"] = "CROSS"
		case entity.MarginModeTypeIsolated:
			m["mode"] = "ISOLATED"
		}
	}

	r.SetFormParams(m)

	data, _, err := s.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}
	log.Println("=a9f605=", string(data))
	answ := futures_marginMode{}
	err = json.Unmarshal(data, &answ)
	if err != nil {
		// return res, err
		log.Println("=futures_setMarginMode ERR=", err)
		return entity.Futures_MarginMode{MarginMode: string(*s.marginMode)}, nil
	}

	res.MarginMode = "cross"
	if utils.StringToInt(answ.Leverage) != 0 {
		res.MarginMode = "isolated"
	}
	// return entity.Futures_MarginMode{MarginMode: string(*s.marginMode)}, nil
	return res, nil
}

type futures_marginMode struct {
	Mode     string `json:"mode"`
	Leverage string `json:"leverage"`
}
