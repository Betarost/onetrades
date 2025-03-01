package futurebybit

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
		Endpoint:   "/v5/position/list",
		TimeOffset: s.c.TimeOffset,
		SecType:    utils.SecTypeSigned,
	}

	r.SetParam("category", "linear")
	r.SetParam("settleCoin", "USDT")

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}

	var answ struct {
		Result struct {
			List []Position `json:"list"`
		} `json:"result"`
	}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}
	return ConvertPositions(answ.Result.List), nil
}

type Position struct {
	Symbol          string `json:"symbol" bson:"symbol"`
	PositionIdx     int    `json:"positionIdx" bson:"positionIdx"`
	Side            string `json:"side" bson:"side"`
	Size            string `json:"size" bson:"size"`
	AvgPrice        string `json:"avgPrice" bson:"avgPrice"`
	PositionValue   string `json:"positionValue" bson:"positionValue"`
	MarkPrice       string `json:"markPrice" bson:"markPrice"`
	UnrealisedPnl   string `json:"unrealisedPnl" bson:"unrealisedPnl"`
	CurRealisedPnl  string `json:"curRealisedPnl" bson:"curRealisedPnl"`
	PositionBalance string `json:"positionBalance" bson:"positionBalance"`
	UpdatedTime     string `json:"updatedTime" bson:"updatedTime"`
}
