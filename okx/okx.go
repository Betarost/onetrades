package okx

import (
	"fmt"
	"log"
	"os"

	"github.com/Betarost/onetrades/utils"
)

var (
	tradeName_Spot    = "OKX_SPOT"
	tradeName_Futures = "OKX_FUTURES"
)

// ===============SPOT=================

type spotClient struct {
	apiKey     string
	secretKey  string
	memo       string
	keyType    string
	BaseURL    string
	UserAgent  string
	Debug      bool
	logger     *log.Logger
	TimeOffset int64
}

func (c *spotClient) debug(format string, v ...interface{}) {
	if c.Debug {
		c.logger.Printf(format, v...)
	}
}

func NewSpotClient(apiKey, secretKey, memo string) *spotClient {
	return &spotClient{
		apiKey:    apiKey,
		secretKey: secretKey,
		memo:      memo,
		keyType:   utils.KeyTypeHmac,
		BaseURL:   utils.GetEndpoint(tradeName_Spot),
		UserAgent: "Onetrades/golang",
		logger:    log.New(os.Stderr, fmt.Sprintf("%s-onetrades ", tradeName_Spot), log.LstdFlags),
	}
}

func (c *spotClient) NewGetInstrumentsInfo() *spot_getInstrumentsInfo {
	return &spot_getInstrumentsInfo{callAPI: c.callAPI}
}

func (c *spotClient) NewGetAccountInfo() *getAccountInfo {
	return &getAccountInfo{callAPI: c.callAPI}
}

func (c *spotClient) NewGetTradingAccountBalance() *getTradingAccountBalance {
	return &getTradingAccountBalance{callAPI: c.callAPI}
}

func (c *spotClient) NewGetFundingAccountBalance() *getFundingAccountBalance {
	return &getFundingAccountBalance{callAPI: c.callAPI}
}

func (c *spotClient) NewGetOrderList() *spot_getOrderList {
	return &spot_getOrderList{callAPI: c.callAPI}
}

func (c *spotClient) NewCancelOrder() *spot_cancelOrder {
	return &spot_cancelOrder{callAPI: c.callAPI}
}

func (c *spotClient) NewMultiCancelOrders() *multiCancelOrders {
	return &multiCancelOrders{callAPI: c.callAPI}
}

func (c *spotClient) NewAmendOrder() *spot_amendOrder {
	return &spot_amendOrder{callAPI: c.callAPI}
}

func (c *spotClient) NewPlaceOrder() *spot_placeOrder {
	return &spot_placeOrder{callAPI: c.callAPI}
}

func (c *spotClient) NewGetBalance() *spot_getBalance {
	return &spot_getBalance{callAPI: c.callAPI}
}

// ===============FUTURES=================

type futuresClient struct {
	apiKey     string
	secretKey  string
	memo       string
	keyType    string
	BaseURL    string
	UserAgent  string
	Debug      bool
	logger     *log.Logger
	TimeOffset int64
}

func (c *futuresClient) debug(format string, v ...interface{}) {
	if c.Debug {
		c.logger.Printf(format, v...)
	}
}

func NewFuturesClient(apiKey, secretKey, memo string) *futuresClient {
	return &futuresClient{
		apiKey:    apiKey,
		secretKey: secretKey,
		memo:      memo,
		keyType:   utils.KeyTypeHmac,
		BaseURL:   utils.GetEndpoint(tradeName_Futures),
		UserAgent: "Onetrades/golang",
		logger:    log.New(os.Stderr, fmt.Sprintf("%s-onetrades ", tradeName_Futures), log.LstdFlags),
	}
}

func (c *futuresClient) NewGetInstrumentsInfo() *futures_getInstrumentsInfo {
	return &futures_getInstrumentsInfo{callAPI: c.callAPI}
}

func (c *futuresClient) NewGetAccountInfo() *getAccountInfo {
	return &getAccountInfo{callAPI: c.callAPI}
}

func (c *futuresClient) NewGetBalance() *futures_getBalance {
	return &futures_getBalance{callAPI: c.callAPI}
}

func (c *futuresClient) NewSetPositionMode() *setPositionMode {
	return &setPositionMode{callAPI: c.callAPI}
}

func (c *futuresClient) NewSetLeverage() *setLeverage {
	return &setLeverage{callAPI: c.callAPI}
}

func (c *futuresClient) NewGetPositions() *getPositions {
	return &getPositions{callAPI: c.callAPI}
}

func (c *futuresClient) NewGetOrderList() *futures_getOrderList {
	return &futures_getOrderList{callAPI: c.callAPI}
}

func (c *futuresClient) NewCancelOrder() *cancelOrder {
	return &cancelOrder{callAPI: c.callAPI}
}

func (c *futuresClient) NewAmendOrder() *amendOrder {
	return &amendOrder{callAPI: c.callAPI}
}

func (c *futuresClient) NewPlaceOrder() *placeOrder {
	return &placeOrder{callAPI: c.callAPI}
}
