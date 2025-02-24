package utils

var (
	BinanceFutureApiMainUrl = "https://fapi.binance.com"
)

func GetApiEndpoint(trade string) string {
	switch trade {
	case "BINANCE":
		return BinanceFutureApiMainUrl
	default:
		return ""
	}
}
