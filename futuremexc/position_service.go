package futuremexc

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
		Endpoint:   "/api/v1/private/position/open_positions",
		TimeOffset: s.c.TimeOffset,
		SecType:    utils.SecTypeSigned,
	}

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}

	var answ struct {
		Result []Position `json:"data"`
	}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}
	return ConvertPositions(answ.Result), nil
}

type Position struct {
	Symbol       string  `json:"symbol" bson:"symbol"`
	PositionId   int     `json:"positionId" bson:"positionId"`
	HoldVol      float64 `json:"holdVol" bson:"holdVol"`
	PositionType int     `json:"positionType" bson:"positionType"`
	HoldAvgPrice float64 `json:"holdAvgPrice" bson:"holdAvgPrice"`
	Oim          float64 `json:"oim" bson:"oim"`
	UpdateTime   int64   `json:"updateTime" bson:"updateTime"`
}
