package entity

type WsHandlerMarkPrice func(event *WsPublicMarkPriceEvent)

type WsPublicMarkPriceEvent struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
	Time   int64   `json:"time"`
}
