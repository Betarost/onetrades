package binance

import (
	"fmt"
	"log"
	"os"

	"github.com/Betarost/onetrades/utils"
)

var (
	tradeName_Spot    = "BINANCE_SPOT"
	tradeName_Futures = "BINANCE_FUTURES"
)

// ===============SPOT=================

type SpotClient struct {
	apiKey     string
	secretKey  string
	keyType    string
	BaseURL    string
	UserAgent  string
	Debug      bool
	logger     *log.Logger
	TimeOffset int64
}

func (c *SpotClient) debug(format string, v ...interface{}) {
	if c.Debug {
		c.logger.Printf(format, v...)
	}
}

func NewSpotClient(apiKey, secretKey string) *SpotClient {
	return &SpotClient{
		apiKey:    apiKey,
		secretKey: secretKey,
		keyType:   utils.KeyTypeHmac,
		BaseURL:   utils.GetEndpoint(tradeName_Spot),
		UserAgent: "Onetrades/golang",
		logger:    log.New(os.Stderr, fmt.Sprintf("%s-onetrades ", tradeName_Spot), log.LstdFlags),
	}
}

func (c *SpotClient) NewGetInstrumentsInfo() *spot_getInstrumentsInfo {
	return &spot_getInstrumentsInfo{callAPI: c.callAPI}
}

func (c *SpotClient) NewGetAccountInfo() *getAccountInfo {
	return &getAccountInfo{callAPI: c.callAPI}
}

// ===============FUTURES=================

type FuturesClient struct {
	apiKey     string
	secretKey  string
	keyType    string
	BaseURL    string
	UserAgent  string
	Debug      bool
	logger     *log.Logger
	TimeOffset int64
}

func (c *FuturesClient) debug(format string, v ...interface{}) {
	if c.Debug {
		c.logger.Printf(format, v...)
	}
}

func NewFuturesClient(apiKey, secretKey string) *FuturesClient {
	return &FuturesClient{
		apiKey:    apiKey,
		secretKey: secretKey,
		keyType:   utils.KeyTypeHmac,
		BaseURL:   utils.GetEndpoint(tradeName_Futures),
		UserAgent: "Onetrades/golang",
		logger:    log.New(os.Stderr, fmt.Sprintf("%s-onetrades ", tradeName_Futures), log.LstdFlags),
	}
}

func (c *FuturesClient) NewGetInstrumentsInfo() *futures_getInstrumentsInfo {
	return &futures_getInstrumentsInfo{callAPI: c.callAPI}
}
