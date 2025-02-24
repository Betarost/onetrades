package futurebinance

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Betarost/onetrades/entity"
	"github.com/Betarost/onetrades/utils"
)

// ==============GetAccountBalance=================
type GetAccountBalance struct {
	c *Client
}

func (s *GetAccountBalance) Do(ctx context.Context, opts ...utils.RequestOption) (res []entity.AccountBalance, err error) {
	r := &utils.Request{
		Method:     http.MethodGet,
		BaseURL:    s.c.BaseURL,
		Endpoint:   "/fapi/v3/balance",
		TimeOffset: s.c.TimeOffset,
		SecType:    utils.SecTypeSigned,
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}

	answ := make([]Balance, 0)
	err = json.Unmarshal(data, &answ)
	if err != nil {
		return res, err
	}
	return ConvertAccountBalance(answ), nil
}

type Balance struct {
	AccountAlias       string `json:"accountAlias"`
	Asset              string `json:"asset"`
	Balance            string `json:"balance"`
	CrossWalletBalance string `json:"crossWalletBalance"`
	CrossUnPnl         string `json:"crossUnPnl"`
	AvailableBalance   string `json:"availableBalance"`
	MaxWithdrawAmount  string `json:"maxWithdrawAmount"`
}

//==============GetAccountService=================

type GetAccountService struct {
	c *Client
}

func (s *GetAccountService) Do(ctx context.Context, opts ...utils.RequestOption) (res *Account, err error) {
	r := &utils.Request{
		Method:     http.MethodGet,
		BaseURL:    s.c.BaseURL,
		Endpoint:   "/fapi/v2/account",
		TimeOffset: s.c.TimeOffset,
		SecType:    utils.SecTypeSigned,
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res = new(Account)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type Account struct {
	Assets                      []*AccountAsset    `json:"assets"`
	FeeTier                     int                `json:"feeTier"`
	CanTrade                    bool               `json:"canTrade"`
	CanDeposit                  bool               `json:"canDeposit"`
	CanWithdraw                 bool               `json:"canWithdraw"`
	UpdateTime                  int64              `json:"updateTime"`
	MultiAssetsMargin           bool               `json:"multiAssetsMargin"`
	TotalInitialMargin          string             `json:"totalInitialMargin"`
	TotalMaintMargin            string             `json:"totalMaintMargin"`
	TotalWalletBalance          string             `json:"totalWalletBalance"`
	TotalUnrealizedProfit       string             `json:"totalUnrealizedProfit"`
	TotalMarginBalance          string             `json:"totalMarginBalance"`
	TotalPositionInitialMargin  string             `json:"totalPositionInitialMargin"`
	TotalOpenOrderInitialMargin string             `json:"totalOpenOrderInitialMargin"`
	TotalCrossWalletBalance     string             `json:"totalCrossWalletBalance"`
	TotalCrossUnPnl             string             `json:"totalCrossUnPnl"`
	AvailableBalance            string             `json:"availableBalance"`
	MaxWithdrawAmount           string             `json:"maxWithdrawAmount"`
	Positions                   []*AccountPosition `json:"positions"`
}

type AccountAsset struct {
	Asset                  string `json:"asset"`
	InitialMargin          string `json:"initialMargin"`
	MaintMargin            string `json:"maintMargin"`
	MarginBalance          string `json:"marginBalance"`
	MaxWithdrawAmount      string `json:"maxWithdrawAmount"`
	OpenOrderInitialMargin string `json:"openOrderInitialMargin"`
	PositionInitialMargin  string `json:"positionInitialMargin"`
	UnrealizedProfit       string `json:"unrealizedProfit"`
	WalletBalance          string `json:"walletBalance"`
	CrossWalletBalance     string `json:"crossWalletBalance"`
	CrossUnPnl             string `json:"crossUnPnl"`
	AvailableBalance       string `json:"availableBalance"`
	MarginAvailable        bool   `json:"marginAvailable"`
	UpdateTime             int64  `json:"updateTime"`
}

type AccountPosition struct {
	Isolated               bool   `json:"isolated"`
	Leverage               string `json:"leverage"`
	InitialMargin          string `json:"initialMargin"`
	MaintMargin            string `json:"maintMargin"`
	OpenOrderInitialMargin string `json:"openOrderInitialMargin"`
	PositionInitialMargin  string `json:"positionInitialMargin"`
	Symbol                 string `json:"symbol"`
	UnrealizedProfit       string `json:"unrealizedProfit"`
	EntryPrice             string `json:"entryPrice"`
	MaxNotional            string `json:"maxNotional"`
	PositionSide           string `json:"positionSide"`
	PositionAmt            string `json:"positionAmt"`
	Notional               string `json:"notional"`
	BidNotional            string `json:"bidNotional"`
	AskNotional            string `json:"askNotional"`
	UpdateTime             int64  `json:"updateTime"`
}

//===============================
