package utils

var (
	BinanceFutureApiMainUrl = "https://fapi.binance.com"
	BybitFutureApiMainUrl   = "https://api.bybit.com"
)

func GetApiEndpoint(trade string) string {
	switch trade {
	case "BINANCE":
		return BinanceFutureApiMainUrl
	case "BYBIT":
		return BybitFutureApiMainUrl
	default:
		return ""
	}
}
