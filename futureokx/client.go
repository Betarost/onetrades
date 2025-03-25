package futureokx

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Betarost/onetrades/utils"
)

var (
	TradeName = "OKX"
)

type Client struct {
	apiKey     string
	secretKey  string
	memo       string
	KeyType    string
	BaseURL    string
	UserAgent  string
	Debug      bool
	Logger     *log.Logger
	TimeOffset int64
	WsPublic   *Ws
	WsPrivate  *Ws
}

func (c *Client) debug(format string, v ...interface{}) {
	if c.Debug {
		c.Logger.Printf(format, v...)
	}
}

func (c *Client) callAPI(ctx context.Context, r *utils.Request, opts ...utils.RequestOption) (data []byte, header *http.Header, err error) {
	r.TmpApi = c.apiKey
	r.TmpSig = c.secretKey
	r.TmpMemo = c.memo
	opts = append(opts, CreateFullURL, CreateBody, CreateSign, CreateHeaders, CreateReq)
	err = r.ParseRequest(opts...)
	if err != nil {
		return []byte{}, &http.Header{}, err

	}
	c.debug("FullURL %s\n", r.FullURL)
	c.debug("Body %s\n", r.BodyString)
	c.debug("Sign %s\n", r.Sign)
	c.debug("Headers %+v\n", r.Header)

	// return []byte{}, &http.Header{}, err

	req, err := http.NewRequest(r.Method, r.FullURL, r.Body)
	if err != nil {
		return []byte{}, &http.Header{}, err
	}
	req = req.WithContext(ctx)
	req.Header = r.Header
	c.debug("request: %#v\n", req)

	res, err, conn := r.DoFunc(req)
	if err != nil {
		return []byte{}, &http.Header{}, err
	}
	data, err = r.ReadAllBody(res.Body, res.Header.Get("Content-Encoding"))
	if err != nil {
		return []byte{}, &http.Header{}, err
	}
	defer func() {
		cerr := res.Body.Close()
		if conn != nil {
			conn.Close()
		}
		if err == nil && cerr != nil {
			err = cerr
		}
	}()

	c.debug("response: %#v\n", res)
	c.debug("response body: %s\n", string(data))
	c.debug("response status code: %d\n", res.StatusCode)

	apiErr := new(APIError)
	e := json.Unmarshal(data, apiErr)
	if e != nil {
		c.debug("failed to unmarshal json: %s\n", e)
	}
	if !apiErr.IsValid() {
		c.debug("Answer Not Walid: %+v\n", apiErr)
		apiErr.Response = data
		return nil, &res.Header, apiErr
	}
	return data, &res.Header, nil
}

func NewClient(apiKey, secretKey, memo string) *Client {
	return &Client{
		apiKey:    apiKey,
		secretKey: secretKey,
		memo:      memo,
		KeyType:   utils.KeyTypeHmac,
		BaseURL:   utils.GetApiEndpoint(TradeName),
		UserAgent: "Onetrades/golang",
		Logger:    log.New(os.Stderr, fmt.Sprintf("%s-onetrades ", TradeName), log.LstdFlags),
	}
}

func (c *Client) NewWebSocketPublicClient() *Ws {
	ws := &Ws{
		Debug:      c.Debug,
		Logger:     c.Logger,
		BaseURL:    utils.GetWsPublicEndpoint(TradeName),
		mapsEvents: []MapsHandler{},
	}
	ws.newConnect(ws.BaseURL)
	c.WsPublic = ws
	return ws
}

func (c *Client) NewWebSocketPrivateClient() *Ws {
	ws := &Ws{
		apiKey:     c.apiKey,
		secretKey:  c.secretKey,
		memo:       c.memo,
		KeyType:    c.KeyType,
		Debug:      c.Debug,
		Logger:     c.Logger,
		BaseURL:    utils.GetWsPrivateEndpoint(TradeName),
		mapsEvents: []MapsHandler{},
	}
	ws.newConnect(ws.BaseURL)
	c.WsPrivate = ws
	return ws
}

func (c *Client) NewGetAccountBalance() *GetAccountBalance {
	return &GetAccountBalance{c: c}
}

func (c *Client) NewGetAccountValuation() *GetAccountValuation {
	return &GetAccountValuation{c: c}
}

func (c *Client) NewGetSubAccountsLists() *GetSubAccountsLists {
	return &GetSubAccountsLists{c: c}
}

func (c *Client) NewGetSubAccountBalance() *GetSubAccountBalance {
	return &GetSubAccountBalance{c: c}
}

func (c *Client) NewGetSubAccountFundingBalance() *GetSubAccountFundingBalance {
	return &GetSubAccountFundingBalance{c: c}
}

func (c *Client) NewGetAccountInfo() *GetAccountInfo {
	return &GetAccountInfo{c: c}
}

func (c *Client) NewGetPositions() *GetPositions {
	return &GetPositions{c: c}
}

func (c *Client) NewGetHistoryPositions() *GetHistoryPositions {
	return &GetHistoryPositions{c: c}
}

func (c *Client) NewGetTradeHistoryOrder() *TradeHistoryOrder {
	return &TradeHistoryOrder{c: c}
}

func (c *Client) NewGetContractsInfo() *GetContractsInfo {
	return &GetContractsInfo{c: c}
}

func (c *Client) NewSetAccountMode() *SetAccountMode {
	return &SetAccountMode{c: c}
}

func (c *Client) NewSetPositionMode() *SetPositionMode {
	return &SetPositionMode{c: c}
}

func (c *Client) NewSetLeverage() *SetLeverage {
	return &SetLeverage{c: c}
}

func (c *Client) NewGetMarkPrice() *GetMarkPrice {
	return &GetMarkPrice{c: c}
}

func (c *Client) NewGetMarkPrices() *GetMarkPrices {
	return &GetMarkPrices{c: c}
}

func (c *Client) NewTradePlaceOrder() *TradePlaceOrder {
	return &TradePlaceOrder{c: c}
}

func (c *Client) NewGetOrderList() *GetOrderList {
	return &GetOrderList{c: c}
}

func (c *Client) NewTradeCancelOrders() *TradeCancelOrders {
	return &TradeCancelOrders{c: c}
}

func (c *Client) NewGetKline() *GetKline {
	return &GetKline{c: c}
}
