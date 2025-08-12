package bybit

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Betarost/onetrades/entity"
	"github.com/Betarost/onetrades/utils"
)

type futures_getPositionMode struct {
	callAPI func(ctx context.Context, r *utils.Request, opts ...utils.RequestOption) (data []byte, header *http.Header, err error)
	convert futures_converts
}

func (s *futures_getPositionMode) Do(ctx context.Context, opts ...utils.RequestOption) (res entity.Futures_PositionsMode, err error) {
	r := &utils.Request{
		Method:   http.MethodGet,
		Endpoint: "/v5/account/info",
		SecType:  utils.SecTypeSigned,
	}

	data, _, err := s.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}
	log.Println("=3c97f9=", string(data))
	var answ struct {
		Result futures_positionMode `json:"data"`
	}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}
	b, err := strconv.ParseBool(answ.Result.DualSidePosition)
	if err != nil {
		return res, err
	}
	return entity.Futures_PositionsMode{HedgeMode: b}, nil
}

type futures_positionMode struct {
	DualSidePosition string `json:"dualSidePosition"`
}
