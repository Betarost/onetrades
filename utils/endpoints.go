package utils

var (
	Binance_spot    = "https://api.binance.com"
	Binance_futures = "https://fapi.binance.com"
	Bingx_spot      = "https://open-api.bingx.com"
	Bingx_futures   = "https://open-api.bingx.com"
	Bybit_spot      = "https://api.bybit.com"
	Bybit_futures   = "https://api.bybit.com"
	Gateio_spot     = "https://api.gateio.ws"
	Gateio_futures  = "https://api.gateio.ws"
	Mexc_spot       = "https://api.mexc.com"
	Mexc_futures    = "https://contract.mexc.com"
	Bitget_spot     = "https://api.bitget.com"
	Bitget_futures  = "https://api.bitget.com"
	Okx_spot        = "https://www.okx.com"
	Okx_futures     = "https://www.okx.com"
	Huobi_spot      = "https://api.huobi.pro"
	Huobi_futures   = "https://api.hbdm.com"
	//================
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

func GetEndpoint(trade string) string {
	switch trade {
	case "BINANCE_SPOT":
		return Binance_spot
	case "BINANCE_FUTURES":
		return Binance_futures
	case "BINGX_SPOT":
		return Bingx_spot
	case "BINGX_FUTURES":
		return Bingx_futures
	case "BYBIT_SPOT":
		return Bybit_spot
	case "BYBIT_FUTURES":
		return Bybit_futures
	case "GATEIO_SPOT":
		return Gateio_spot
	case "GATEIO_FUTURES":
		return Gateio_futures
	case "MEXC_SPOT":
		return Mexc_spot
	case "MEXC_FUTURES":
		return Mexc_futures
	case "BITGET_SPOT":
		return Bitget_spot
	case "BITGET_FUTURES":
		return Bitget_futures
	case "OKX_SPOT":
		return Okx_spot
	case "OKX_FUTURES":
		return Okx_futures
	case "HUOBI_SPOT":
		return Huobi_spot
	case "HUOBI_FUTURES":
		return Huobi_futures
	default:
		return ""
	}
}
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

func GetApiEndpointOption(trade string) string {
	switch trade {
	case "OKX":
		return OKXFutureApiMainUrl
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
