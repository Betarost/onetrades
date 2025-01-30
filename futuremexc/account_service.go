package futuremexc

import (
	"context"
	"encoding/json"
	"net/http"
)

type GetAccountAssets struct {
	c *Client
}

// Do send request
func (s *GetAccountAssets) Do(ctx context.Context, opts ...RequestOption) (res []AccountAssets, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v1/private/account/assets",
		secType:  secTypeSigned,
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}

	var answer struct {
		Success bool            `json:"success"`
		Code    int64           `json:"code"`
		Message string          `json:"message"`
		Res     []AccountAssets `json:"data"` // Assign the body value when the Code and Message fields are invalid.
	}

	err = json.Unmarshal(data, &answer)
	if err != nil {
		return res, err
	}
	return answer.Res, nil
}

type AccountAssets struct {
	Currency         string  `json:"currency"`
	PositionMargin   float64 `json:"positionMargin"`
	FrozenBalance    float64 `json:"frozenBalance"`
	AvailableBalance float64 `json:"availableBalance"`
	CashBalance      float64 `json:"cashBalance"`
	Equity           float64 `json:"equity"`
	Unrealized       float64 `json:"unrealized"`
}
