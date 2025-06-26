package gateio

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/Betarost/onetrades/entity"
	"github.com/Betarost/onetrades/utils"
)

// ==============SetPositionMode=================
type futures_setPositionMode struct {
	callAPI func(ctx context.Context, r *utils.Request, opts ...utils.RequestOption) (data []byte, header *http.Header, err error)
	convert futures_converts

	settle *string
	mode   *entity.PositionModeType
}

func (s *futures_setPositionMode) Settle(settle string) *futures_setPositionMode {
	s.settle = &settle
	return s
}

func (s *futures_setPositionMode) Mode(mode entity.PositionModeType) *futures_setPositionMode {
	s.mode = &mode
	return s
}

func (s *futures_setPositionMode) Do(ctx context.Context, opts ...utils.RequestOption) (res bool, err error) {
	r := &utils.Request{
		Method:   http.MethodPost,
		Endpoint: "/api/v4/futures/{settle}/dual_mode",
		SecType:  utils.SecTypeSigned,
	}

	m := utils.Params{}

	settleDefault := "usdt"

	if s.settle == nil {
		s.settle = &settleDefault
	}

	r.Endpoint = strings.Replace(r.Endpoint, "{settle}", *s.settle, 1)

	if s.mode != nil {
		if *s.mode == entity.PositionModeTypeHedge {
			m["dual_mode"] = true
		} else if *s.mode == entity.PositionModeTypeOneWay {
			m["dual_mode"] = false
		}
	}
	r.SetParams(m)

	data, _, err := s.callAPI(ctx, r, opts...)
	if err != nil {
		return false, err
	}

	log.Println("=d6461a=", string(data))
	answ := positionMode{}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return false, err
	}

	return true, nil
}

type positionMode struct {
	In_dual_mode bool `json:"in_dual_mode"`
}
