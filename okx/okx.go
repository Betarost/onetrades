package okx

import (
	"fmt"
	"log"
	"os"

	"github.com/Betarost/onetrades/utils"
)

var (
	tradeName_Spot    = "OKX"
	tradeName_Futures = "OKX"
)

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
	// WsPublic   *Ws
	// WsPrivate  *Ws
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
		BaseURL:   utils.GetApiEndpoint(tradeName_Spot),
		UserAgent: "Onetrades/golang",
		logger:    log.New(os.Stderr, fmt.Sprintf("%s-onetrades ", tradeName_Spot), log.LstdFlags),
	}
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

func (c *spotClient) NewGetInstrumentsInfo() *getInstrumentsInfo {
	return &getInstrumentsInfo{callAPI: c.callAPI}
}

func (c *spotClient) NewGetOrderList() *getOrderList {
	return &getOrderList{callAPI: c.callAPI}
}

func (c *spotClient) NewCancelOrder() *cancelOrder {
	return &cancelOrder{callAPI: c.callAPI}
}

func (c *spotClient) NewMultiCancelOrders() *multiCancelOrders {
	return &multiCancelOrders{callAPI: c.callAPI}
}

func (c *spotClient) NewAmendOrder() *amendOrder {
	return &amendOrder{callAPI: c.callAPI}
}

func (c *spotClient) NewPlaceOrder() *placeOrder {
	return &placeOrder{callAPI: c.callAPI}
}
