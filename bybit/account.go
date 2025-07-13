package bybit

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Betarost/onetrades/entity"
	"github.com/Betarost/onetrades/utils"
)

// ===================GetAccountInfo==================
type getAccountInfo struct {
	callAPI func(ctx context.Context, r *utils.Request, opts ...utils.RequestOption) (data []byte, header *http.Header, err error)
	convert account_converts
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
		Result accountInfo `json:"result"`
	}
	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}
	return s.convert.convertAccountInfo(answ.Result), nil
}

type accountInfo struct {
	UserID      int64    `json:"userID"`
	ID          string   `json:"id"`
	Note        string   `json:"note"`
	Ips         []string `json:"ips"`
	ReadOnly    int64    `json:"readOnly"`
	Permissions struct {
		Spot        []string `json:"Spot"`
		Derivatives []string `json:"Derivatives"`
		Wallet      []string `json:"Wallet"`
	} `json:"permissions"`
}

type tradingBalance struct {
	TotalEquity           string                  `json:"totalEquity"`
	TotalAvailableBalance string                  `json:"totalAvailableBalance"`
	TotalPerpUPL          string                  `json:"totalPerpUPL"`
	Coin                  []tradingBalanceDetails `json:"coin"`
}

type tradingBalanceDetails struct {
	Coin                string `json:"coin"`
	Equity              string `json:"equity"`
	WalletBalance       string `json:"walletBalance"`
	UnrealisedPnl       string `json:"unrealisedPnl"`
	AvailableToWithdraw string `json:"availableToWithdraw"`
}

type fundingBalance struct {
	Coin            string `json:"coin"`
	TransferBalance string `json:"transferBalance"`
	WalletBalance   string `json:"walletBalance"`
}
