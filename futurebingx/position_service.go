package futurebingx

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
		Endpoint:   "/openApi/swap/v2/user/positions",
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
	Symbol             string  `json:"symbol" bson:"symbol"`
	PositionID         string  `json:"positionId" bson:"positionId"`
	PositionSide       string  `json:"positionSide" bson:"positionSide"`
	Isolated           bool    `json:"isolated" bson:"isolated"`
	PositionAmt        string  `json:"positionAmt" bson:"positionAmt"`
	AvailableAmt       string  `json:"availableAmt" bson:"availableAmt"`
	UnrealizedProfit   string  `json:"unrealizedProfit" bson:"unrealizedProfit"`
	RealisedProfit     string  `json:"realisedProfit" bson:"realisedProfit"`
	InitialMargin      string  `json:"initialMargin" bson:"initialMargin"`
	Margin             string  `json:"margin" bson:"margin"`
	AvgPrice           string  `json:"avgPrice" bson:"avgPrice"`
	LiquidationPrice   float64 `json:"liquidationPrice" bson:"liquidationPrice"`
	Leverage           int64   `json:"leverage" bson:"leverage"`
	PositionValue      string  `json:"positionValue" bson:"positionValue"`
	MarkPrice          string  `json:"markPrice" bson:"markPrice"`
	RiskRate           string  `json:"riskRate" bson:"riskRate"`
	MaxMarginReduction string  `json:"maxMarginReduction" bson:"maxMarginReduction"`
	PNLRatio           string  `json:"pnlRatio" bson:"pnlRatio"`
	UpdateTime         int64   `json:"updateTime" bson:"updateTime"`
}
