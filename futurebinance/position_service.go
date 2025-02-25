package futurebinance

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Betarost/onetrades/entity"
	"github.com/Betarost/onetrades/utils"
)

// ==============GetPositions=================
type GetPositions struct {
	c *Client
}

func (s *GetPositions) Do(ctx context.Context, opts ...utils.RequestOption) (res []entity.Position, err error) {
	r := &utils.Request{
		Method:     http.MethodGet,
		BaseURL:    s.c.BaseURL,
		Endpoint:   "/fapi/v3/positionRisk",
		TimeOffset: s.c.TimeOffset,
		SecType:    utils.SecTypeSigned,
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}

	answ := make([]Position, 0)
	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}
	return ConvertPositions(answ), nil
}

type Position struct {
	Symbol           string `json:"symbol" bson:"symbol"`
	PositionSide     string `json:"positionSide" bson:"positionSide"`
	PositionAmt      string `json:"positionAmt" bson:"positionAmt"`
	EntryPrice       string `json:"entryPrice" bson:"entryPrice"`
	MarkPrice        string `json:"markPrice" bson:"markPrice"`
	UnRealizedProfit string `json:"unRealizedProfit" bson:"unRealizedProfit"`
	Notional         string `json:"notional" bson:"notional"`
	InitialMargin    string `json:"initialMargin" bson:"initialMargin"`
	UpdateTime       int64  `json:"updateTime" bson:"updateTime"`
}
