package onetrades

import (
	"github.com/Betarost/onetrades/binance"
	"github.com/Betarost/onetrades/bingx"
	"github.com/Betarost/onetrades/bitget"
	"github.com/Betarost/onetrades/bullish"
	"github.com/Betarost/onetrades/bybit"
	"github.com/Betarost/onetrades/gateio"
	"github.com/Betarost/onetrades/huobi"
	"github.com/Betarost/onetrades/kucoin"
	"github.com/Betarost/onetrades/mexc"
	"github.com/Betarost/onetrades/okx"
)

// Общие креды для всех бирж
type Credentials struct {
	APIKey    string
	SecretKey string
	Memo      string
}

// Общие опции
type Options struct {
	COINM      bool   // для фьючерсов у поддерживаемых бирж
	Proxy      string // http(s)://host:port
	UserAgent  string // кастомный UA
	Debug      bool
	BrokerID   string
	TimeOffset int64
}

type HasProxy interface{ SetProxy(string) }
type HasUserAgent interface{ SetUserAgent(string) }
type HasDebug interface{ SetDebug(bool) }
type HasBrokerID interface{ SetBrokerID(string) }
type HasTimeOffset interface{ SetTimeOffset(int64) }
type HasCOINM interface{ SetCOINM(bool) }

func applyOptions(cli any, opts Options) {
	if v, ok := cli.(HasProxy); ok && opts.Proxy != "" {
		v.SetProxy(opts.Proxy)
	}
	if v, ok := cli.(HasUserAgent); ok && opts.UserAgent != "" {
		v.SetUserAgent(opts.UserAgent)
	}
	if v, ok := cli.(HasDebug); ok {
		v.SetDebug(opts.Debug)
	}
	if v, ok := cli.(HasBrokerID); ok && opts.BrokerID != "" {
		v.SetBrokerID(opts.BrokerID)
	}
	if v, ok := cli.(HasTimeOffset); ok && opts.TimeOffset != 0 {
		v.SetTimeOffset(opts.TimeOffset)
	}
	if v, ok := cli.(HasCOINM); ok {
		v.SetCOINM(opts.COINM)
	}
}

// ===================== BINANCE =====================

func NewBinanceSpot(cred Credentials) *binance.SpotClient {
	return binance.NewSpotClient(cred.APIKey, cred.SecretKey)
}

func NewBinanceFutures(cred Credentials, opts Options) *binance.FuturesClient {
	c := binance.NewFuturesClient(cred.APIKey, cred.SecretKey)
	if opts.COINM {
		c.IsCOINM(true)
	}
	return c
}

// ===================== BYBIT =====================

func NewBybitSpot(cred Credentials) *bybit.SpotClient {
	return bybit.NewSpotClient(cred.APIKey, cred.SecretKey)
}

func NewBybitFutures(cred Credentials) *bybit.FuturesClient {
	return bybit.NewFuturesClient(cred.APIKey, cred.SecretKey)
}

// ===================== OKX =====================

func NewOKXSpot(cred Credentials) *okx.SpotClient {
	return okx.NewSpotClient(cred.APIKey, cred.SecretKey, cred.Memo)
}

func NewOKXFutures(cred Credentials) *okx.FuturesClient {
	return okx.NewFuturesClient(cred.APIKey, cred.SecretKey, cred.Memo)
}

// ===================== BINGX =====================

func NewBingXSpot(cred Credentials) *bingx.SpotClient {
	return bingx.NewSpotClient(cred.APIKey, cred.SecretKey)
}

func NewBingXFutures(cred Credentials) *bingx.FuturesClient {
	return bingx.NewFuturesClient(cred.APIKey, cred.SecretKey)
}

// ===================== BITGET =====================

func NewBitgetSpot(cred Credentials) *bitget.SpotClient {
	// У Bitget обычно требуется passphrase/memo
	return bitget.NewSpotClient(cred.APIKey, cred.SecretKey, cred.Memo)
}

func NewBitgetFutures(cred Credentials) *bitget.FuturesClient {
	return bitget.NewFuturesClient(cred.APIKey, cred.SecretKey, cred.Memo)
}

// ===================== BULLISH =====================

func NewBullishFutures(cred Credentials) *bullish.FuturesClient {
	return bullish.NewFuturesClient(cred.APIKey, cred.SecretKey, cred.Memo)
}

// ===================== GATE.IO =====================

func NewGateIOSpot(cred Credentials) *gateio.SpotClient {
	return gateio.NewSpotClient(cred.APIKey, cred.SecretKey)
}

func NewGateIOFutures(cred Credentials) *gateio.FuturesClient {
	return gateio.NewFuturesClient(cred.APIKey, cred.SecretKey)
}

// ===================== HUOBI =====================

func NewHuobiSpot(cred Credentials) *huobi.SpotClient {
	return huobi.NewSpotClient(cred.APIKey, cred.SecretKey)
}

func NewHuobiFutures(cred Credentials) *huobi.FuturesClient {
	return huobi.NewFuturesClient(cred.APIKey, cred.SecretKey)
}

// ===================== KUCOIN =====================

func NewKucoinSpot(cred Credentials) *kucoin.SpotClient {
	// KuCoin требует passphrase
	return kucoin.NewSpotClient(cred.APIKey, cred.SecretKey, cred.Memo)
}

func NewKucoinFutures(cred Credentials) *kucoin.FuturesClient {
	return kucoin.NewFuturesClient(cred.APIKey, cred.SecretKey, cred.Memo)
}

// ===================== MEXC =====================

func NewMEXCSpot(cred Credentials) *mexc.SpotClient {
	return mexc.NewSpotClient(cred.APIKey, cred.SecretKey)
}
