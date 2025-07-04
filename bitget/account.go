package bitget

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Betarost/onetrades/entity"
	"github.com/Betarost/onetrades/utils"
)

type getAccountInfo struct {
	callAPI func(ctx context.Context, r *utils.Request, opts ...utils.RequestOption) (data []byte, header *http.Header, err error)
	convert account_converts
}

func (s *getAccountInfo) Do(ctx context.Context, opts ...utils.RequestOption) (res entity.AccountInformation, err error) {
	r := &utils.Request{
		Method:   http.MethodGet,
		Endpoint: "/api/v2/spot/account/info",
		SecType:  utils.SecTypeSigned,
	}

	data, _, err := s.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}

	var answ struct {
		Result accountInfo `json:"data"`
	}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}
	return s.convert.convertAccountInfo(answ.Result), nil
}

type accountInfo struct {
	UserId      string   `json:"userId"`
	Ips         string   `json:"ips"`
	Authorities []string `json:"authorities"`
}
