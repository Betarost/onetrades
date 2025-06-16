package bybit

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Betarost/onetrades/entity"
	"github.com/Betarost/onetrades/utils"
)

// ===================GetAccountInfo==================
type getAccountInfo struct {
	callAPI func(ctx context.Context, r *utils.Request, opts ...utils.RequestOption) (data []byte, header *http.Header, err error)
}

func (s *getAccountInfo) Do(ctx context.Context, opts ...utils.RequestOption) (res entity.AccountInformation, err error) {
	r := &utils.Request{
		Method:   http.MethodGet,
		Endpoint: "/v5/user/query-api",
		SecType:  utils.SecTypeSigned,
	}

	data, _, err := s.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}
	var answ struct {
		RetCode int64       `json:"retCode"`
		RetMsg  string      `json:"retMsg"`
		Result  accountInfo `json:"result"`
	}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}

	if answ.RetCode != 0 {
		return res, errors.New(answ.RetMsg)
	}
	return convertAccountInfo(answ.Result), nil
}

type accountInfo struct {
	UserID      int64           `json:"userID"`
	ID          string          `json:"id"`
	Note        string          `json:"note"`
	Ips         []string        `json:"ips"`
	ReadOnly    int64           `json:"readOnly"`
	Permissions permAccountInfo `json:"permissions"`
}

type permAccountInfo struct {
	Spot        []string `json:"Spot"`
	Derivatives []string `json:"Derivatives"`
	Wallet      []string `json:"Wallet"`
}
