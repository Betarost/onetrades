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
type spot_getAccountInfo struct {
	callAPI func(ctx context.Context, r *utils.Request, opts ...utils.RequestOption) (data []byte, header *http.Header, err error)
}

func (s *spot_getAccountInfo) Do(ctx context.Context, opts ...utils.RequestOption) (res entity.AccountInformation, err error) {
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

//===================getTradingAccountBalance==================

type spot_getTradingAccountBalance struct {
	callAPI func(ctx context.Context, r *utils.Request, opts ...utils.RequestOption) (data []byte, header *http.Header, err error)
}

func (s *spot_getTradingAccountBalance) Do(ctx context.Context, opts ...utils.RequestOption) (res []entity.TradingAccountBalance, err error) {
	r := &utils.Request{
		Method:   http.MethodGet,
		Endpoint: "/v5/account/wallet-balance",
		SecType:  utils.SecTypeSigned,
	}

	r.SetParam("accountType", "UNIFIED")

	data, _, err := s.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}

	var answ struct {
		RetCode int64  `json:"retCode"`
		RetMsg  string `json:"retMsg"`
		Result  struct {
			List []tradingBalance `json:"list"`
		} `json:"result"`
		Time int64 `json:"time"`
	}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}

	if answ.RetCode != 0 {
		return res, errors.New(answ.RetMsg)
	}

	return convertTradingAccountBalance(answ.Result.List), nil
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

//===================getFundingAccountBalance==================

type spot_getFundingAccountBalance struct {
	callAPI func(ctx context.Context, r *utils.Request, opts ...utils.RequestOption) (data []byte, header *http.Header, err error)
}

func (s *spot_getFundingAccountBalance) Do(ctx context.Context, opts ...utils.RequestOption) (res []entity.FundingAccountBalance, err error) {
	r := &utils.Request{
		Method:   http.MethodGet,
		Endpoint: "/v5/asset/transfer/query-account-coins-balance",
		SecType:  utils.SecTypeSigned,
	}

	r.SetParam("accountType", "FUND")

	data, _, err := s.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}

	var answ struct {
		RetCode int64  `json:"retCode"`
		RetMsg  string `json:"retMsg"`
		Result  struct {
			Balance []fundingBalance `json:"balance"`
		} `json:"result"`
		Time int64 `json:"time"`
	}

	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}

	if answ.RetCode != 0 {
		return res, errors.New(answ.RetMsg)
	}

	return convertFundingAccountBalance(answ.Result.Balance), nil
}

type fundingBalance struct {
	Coin            string `json:"coin"`
	TransferBalance string `json:"transferBalance"`
	WalletBalance   string `json:"walletBalance"`
}
