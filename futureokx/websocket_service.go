package futureokx

import (
	"encoding/json"

	"github.com/Betarost/onetrades/entity"
)

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
