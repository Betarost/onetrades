package utils

var (
	BinanceFutureApiMainUrl = "https://fapi.binance.com"
	BybitFutureApiMainUrl   = "https://api.bybit.com"
	MexcFutureApiMainUrl    = "https://futures.mexc.com"
	BingxFutureApiMainUrl   = "https://open-api.bingx.com"
	GateFutureApiMainUrl    = "https://api.gateio.ws"
	BitgetFutureApiMainUrl  = "https://api.bitget.com"
	OKXFutureApiMainUrl     = "https://www.okx.com"
	HuobiFutureApiMainUrl   = "https://api.hbdm.com"

	OKXFutureWsPublicUrl  = "wss://ws.okx.com:8443/ws/v5/public"
	OKXFutureWsPrivateUrl = "wss://ws.okx.com:8443/ws/v5/private"
)

func GetApiEndpoint(trade string) string {
	switch trade {
	case "BINANCE":
		return BinanceFutureApiMainUrl
	case "BYBIT":
		return BybitFutureApiMainUrl
	case "MEXC":
		return MexcFutureApiMainUrl
	case "BINGX":
		return BingxFutureApiMainUrl
	case "GATE":
		return GateFutureApiMainUrl
	case "BITGET":
		return BitgetFutureApiMainUrl
	case "OKX":
		return OKXFutureApiMainUrl
	case "HUOBI":
		return HuobiFutureApiMainUrl
	default:
		return ""
	}
}

func GetWsPublicEndpoint(trade string) string {
	switch trade {
	case "OKX":
		return OKXFutureWsPublicUrl
	default:
		return ""
	}
}

func GetWsPrivateEndpoint(trade string) string {
	switch trade {
	case "OKX":
		return OKXFutureWsPrivateUrl
	default:
		return ""
	}
}
