package binance

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Betarost/onetrades/entity"
	"github.com/Betarost/onetrades/utils"
)

type futures_getInstrumentsInfo struct {
	callAPI func(ctx context.Context, r *utils.Request, opts ...utils.RequestOption) (data []byte, header *http.Header, err error)
	convert futures_converts

	symbol *string
}

func (s *futures_getInstrumentsInfo) Symbol(symbol string) *futures_getInstrumentsInfo {
	s.symbol = &symbol
	return s
}

func (s *futures_getInstrumentsInfo) Do(ctx context.Context, opts ...utils.RequestOption) (res []entity.Futures_InstrumentsInfo, err error) {
	r := &utils.Request{
		Method:   http.MethodGet,
		Endpoint: "/fapi/v1/exchangeInfo",
		SecType:  utils.SecTypeNone,
	}

	data, _, err := s.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}

	answ := futures_instrumentsInfo{}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}

	if s.symbol != nil {
		a := futures_instrumentsInfo{}
		for _, item := range answ.Symbols {
			if item.Symbol == *s.symbol {
				a.Symbols = append(a.Symbols, struct {
					Symbol             string        "json:\"symbol\""
					Status             string        "json:\"status\""
					BaseAsset          string        "json:\"baseAsset\""
					QuoteAsset         string        "json:\"quoteAsset\""
					BaseAssetPrecision int64         "json:\"baseAssetPrecision\""
					Filters            []interface{} "json:\"filters\""
				}{
					Symbol:             item.Symbol,
					Status:             item.Status,
					BaseAsset:          item.BaseAsset,
					QuoteAsset:         item.QuoteAsset,
					BaseAssetPrecision: item.BaseAssetPrecision,
					Filters:            item.Filters,
				})
				break
			}
		}
		return s.convert.convertInstrumentsInfo(a), nil
	}

	return s.convert.convertInstrumentsInfo(answ), nil
}

type futures_instrumentsInfo struct {
	ServerTime int64 `json:"serverTime"`
	Symbols    []struct {
		Symbol             string        `json:"symbol"`
		Status             string        `json:"status"`
		BaseAsset          string        `json:"baseAsset"`
		QuoteAsset         string        `json:"quoteAsset"`
		BaseAssetPrecision int64         `json:"baseAssetPrecision"`
		Filters            []interface{} `json:"filters"`
	} `json:"symbols"`
}
