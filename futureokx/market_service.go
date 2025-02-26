package futureokx

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Betarost/onetrades/entity"
	"github.com/Betarost/onetrades/utils"
)

// ==============GetPositions=================
type GetContractsInfo struct {
	c *Client
}

func (s *GetContractsInfo) Do(ctx context.Context, opts ...utils.RequestOption) (res []entity.ContractInfo, err error) {
	r := &utils.Request{
		Method:     http.MethodGet,
		BaseURL:    s.c.BaseURL,
		Endpoint:   "/api/v5/public/instruments",
		TimeOffset: s.c.TimeOffset,
		SecType:    utils.SecTypeNone,
	}

	m := utils.Params{
		"instType": "SWAP",
	}
	r.SetParams(m)

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}

	var answ struct {
		Result []ContractsInfo `json:"data"`
	}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}
	return ConvertContractsInfo(answ.Result), nil
}

type ContractsInfo struct {
	InstId string `json:"instId"`
	CtVal  string `json:"ctVal"`
	CtMult string `json:"ctMult"`
	TickSz string `json:"tickSz"`
	LotSz  string `json:"lotSz"`
}
