package bybit

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Betarost/onetrades/entity"
	"github.com/Betarost/onetrades/utils"
)

// ==============GetInstrumentsInfo=================
type futures_getInstrumentsInfo struct {
	callAPI func(ctx context.Context, r *utils.Request, opts ...utils.RequestOption) (data []byte, header *http.Header, err error)

	symbol *string
}

func (s *futures_getInstrumentsInfo) Symbol(symbol string) *futures_getInstrumentsInfo {
	s.symbol = &symbol
	return s
}

func (s *futures_getInstrumentsInfo) Do(ctx context.Context, opts ...utils.RequestOption) (res []entity.InstrumentsInfo, err error) {
	r := &utils.Request{
		Method:   http.MethodGet,
		Endpoint: "/v5/market/instruments-info",
		SecType:  utils.SecTypeNone,
	}

	m := utils.Params{
		"category": "linear",
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
			List []futures_instrumentsInfo `json:"list"`
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
	return futures_convertInstrumentsInfo(answ.Result.List, "FUTURES"), nil
}

type futures_instrumentsInfo struct {
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
