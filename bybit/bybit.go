package bybit

import (
	"fmt"
	"log"
	"os"

	"github.com/Betarost/onetrades/utils"
)

var (
	tradeName_Spot    = "BYBIT"
	tradeName_Futures = "BYBIT"
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

func NewSpotClient(apiKey, secretKey string) *spotClient {
	return &spotClient{
		apiKey:    apiKey,
		secretKey: secretKey,
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
