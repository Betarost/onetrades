package utils

var (
	BinanceFutureApiMainUrl = "https://fapi.binance.com"
	BybitFutureApiMainUrl   = "https://api.bybit.com"
	MexcFutureApiMainUrl    = "https://futures.mexc.com"
)

func GetApiEndpoint(trade string) string {
	switch trade {
	case "BINANCE":
		return BinanceFutureApiMainUrl
	case "BYBIT":
		return BybitFutureApiMainUrl
	case "MEXC":
		return MexcFutureApiMainUrl
	default:
		return ""
	}
}
