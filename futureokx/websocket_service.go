package futureokx

import (
	"encoding/json"

	"github.com/Betarost/onetrades/entity"
)

// ==============PrivateOrders=================
type ArgsPrivateOrders struct {
	Channel  string `json:"channel"`
	InstType string `json:"instType"`
	// InstId   string `json:"instId"`
}

type PrivateOrders struct {
	InstType    string `json:"instType"`
	InstId      string `json:"instId"`
	OrdId       string `json:"ordId"`
	Px          string `json:"px"`
	LastPx      string `json:"lastPx"`
	Sz          string `json:"sz"`
	NotionalUsd string `json:"notionalUsd"`
	OrdType     string `json:"ordType"`
	Side        string `json:"side"`
	PosSide     string `json:"posSide"`
	State       string `json:"state"` // live filled
	Lever       string `json:"lever"`
	CTime       string `json:"cTime"`
	UTime       string `json:"uTime"`

	AccFillSz       string `json:"accFillSz"`
	FillNotionalUsd string `json:"fillNotionalUsd"`
	AvgPx           string `json:"avgPx"`
	Pnl             string `json:"pnl"`
	FeeCcy          string `json:"feeCcy"`
	Fee             string `json:"fee"`
	FillPx          string `json:"fillPx"`
	FillSz          string `json:"fillSz"`
}

type AnswPrivateOrders struct {
	Arg  ArgsSymbol      `json:"arg"`
	Data []PrivateOrders `json:"data"`
}

func (w *Ws) NewPrivateOrders(handler entity.WsHandlerPrivateOrders, errHandler ErrHandler) error {

	channel := "orders"
	h := func(message []byte) {
		mess := AnswPrivateOrders{}
		json.Unmarshal(message, &mess)
		for _, item := range mess.Data {
			d := ConvertWsPrivateOrders(item)
			handler(&d)
		}
	}
	w.mapsEvents = append(w.mapsEvents, MapsHandler{
		Channel:    channel,
		WsHandler:  h,
		ErrHandler: errHandler,
	})

	mess := struct {
		Op   string              `json:"op"`
		Args []ArgsPrivateOrders `json:"args"`
	}{
		Op: "subscribe",
		Args: []ArgsPrivateOrders{
			{
				Channel:  channel,
				InstType: "SWAP",
				// InstId:   symbol,
			},
		},
	}

	if w.c != nil {
		w.c.WriteJSON(mess)
	}
	return nil
}

// ==============PrivatePositions=================
type ArgsPrivatePositions struct {
	Channel  string `json:"channel"`
	InstType string `json:"instType"`
	// InstId   string `json:"instId"`
}

func (w *Ws) NewPrivatePositions(symbol string, handler entity.WsHandlerPrivatePositions, errHandler ErrHandler) error {

	channel := "positions"
	// h := func(message []byte) {
	// 	mess := AnswPublicMarkPrice{}
	// 	json.Unmarshal(message, &mess)
	// 	for _, item := range mess.Data {
	// 		d := ConvertWsMarkPrice(item)
	// 		handler(&d)
	// 	}
	// }
	// w.mapsEvents = append(w.mapsEvents, MapsHandler{
	// 	Channel:    channel,
	// 	WsHandler:  h,
	// 	ErrHandler: errHandler,
	// })

	args := []ArgsPrivatePositions{
		{
			Channel:  channel,
			InstType: "SWAP",
			// InstId:   symbol,
		},
	}

	mess := struct {
		Op   string                 `json:"op"`
		Args []ArgsPrivatePositions `json:"args"`
	}{
		Op:   "subscribe",
		Args: args,
	}

	if w.c != nil {
		w.c.WriteJSON(mess)
	}
	return nil
}

// ==============PublicMarkPrice=================
type ArgsSymbol struct {
	Channel string `json:"channel"`
	InstId  string `json:"instId"`
}

type PublicMarkPrice struct {
	InstId   string `json:"instId"`
	InstType string `json:"instType"`
	MarkPx   string `json:"markPx"`
	Ts       string `json:"ts"`
}

type AnswPublicMarkPrice struct {
	Arg  ArgsSymbol        `json:"arg"`
	Data []PublicMarkPrice `json:"data"`
}

func (w *Ws) NewPublicMarkPrice(symbol string, handler entity.WsHandlerMarkPrice, errHandler ErrHandler) error {

	channel := "mark-price"
	h := func(message []byte) {
		mess := AnswPublicMarkPrice{}
		json.Unmarshal(message, &mess)
		for _, item := range mess.Data {
			d := ConvertWsMarkPrice(item)
			handler(&d)
		}
	}
	w.mapsEvents = append(w.mapsEvents, MapsHandler{
		Channel:    channel,
		WsHandler:  h,
		ErrHandler: errHandler,
	})

	args := []ArgsSymbol{
		{
			Channel: channel,
			InstId:  symbol,
		},
	}

	mess := struct {
		Op   string       `json:"op"`
		Args []ArgsSymbol `json:"args"`
	}{
		Op:   "subscribe",
		Args: args,
	}

	if w.c != nil {
		w.c.WriteJSON(mess)
	}
	return nil
}
