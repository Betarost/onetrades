package bybit

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

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

func (s *futures_setMarginMode) Do(ctx context.Context, opts ...utils.RequestOption) (res entity.Futures_Leverage, err error) {
	r := &utils.Request{
		Method:   http.MethodPost,
		Endpoint: "/v5/position/switch-isolated",
		SecType:  utils.SecTypeSigned,
	}

	m := utils.Params{"category": "linear"}
	if s.symbol != nil {
		m["symbol"] = *s.symbol
	}

	// if s.leverage != nil {
	// 	m["buyLeverage"] = *s.leverage
	// 	m["sellLeverage"] = *s.leverage
	// }

	if s.marginMode != nil {
		switch *s.marginMode {
		case entity.MarginModeTypeCross:
			m["tradeMode"] = 0
		case entity.MarginModeTypeIsolated:
			m["tradeMode"] = 1
		}
	}

	r.SetFormParams(m)

	data, _, err := s.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}

	log.Println("=a6ad66=", string(data))
	var answ struct {
		RetMsg string `json:"retMsg"`
		// Result futures_leverage `json:"data"`
	}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}

	if answ.RetMsg != "OK" {
		return res, errors.New("Wrong Answer")
	}
	if err != nil {
		return res, err
	}

	// return entity.Futures_MarginMode{MarginMode: string(*s.marginMode)}, nil
	return entity.Futures_Leverage{Symbol: *s.symbol, Leverage: *s.leverage}, nil
}

// type futures_leverage struct {
// 	Symbol           string `json:"symbol"`
// 	LongLeverage     int    `json:"longLeverage"`
// 	MaxLongLeverage  int    `json:"maxLongLeverage"`
// 	MaxShortLeverage int    `json:"maxShortLeverage"`
// }
