package futureokx

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Betarost/onetrades/entity"
)

// ==============PrivatePlaceOrders=================

func (w *Ws) NewPrivateCancelOrders(data []map[string]string) error {

	channel := "batch-cancel-orders"

	mess := struct {
		Id   string              `json:"id"`
		Op   string              `json:"op"`
		Args []map[string]string `json:"args"`
	}{
		Id:   fmt.Sprintf("%d", time.Now().Unix()),
		Op:   channel,
		Args: data,
	}

	if w.c != nil {
		w.c.WriteJSON(mess)
	}
	return nil
}

// ==============PrivatePlaceOrders=================

func (w *Ws) NewPrivatePlaceOrders(data []map[string]string) error {

	channel := "batch-orders"

	mess := struct {
		Id   string              `json:"id"`
		Op   string              `json:"op"`
		Args []map[string]string `json:"args"`
	}{
		Id:   fmt.Sprintf("%d", time.Now().Unix()),
		Op:   channel,
		Args: data,
	}

	if w.c != nil {
		w.c.WriteJSON(mess)
	}
	return nil
}

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
	State       string `json:"state"` // live filled canceled
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
}

type PrivatePositions struct {
	InstID      string `json:"instId"`
	PosId       string `json:"posId"`
	PosSide     string `json:"posSide"`
	Ccy         string `json:"ccy"`
	Pos         string `json:"pos"`
	AvailPos    string `json:"availPos,omitempty"`
	NotionalUsd string `json:"notionalUsd"`
	AvgPx       string `json:"avgPx"`
	BePx        string `json:"bePx"`
	Fee         string `json:"fee"`
	FundingFee  string `json:"fundingFee"`
	Imr         string `json:"imr,omitempty"`
	Last        string `json:"last"`
	MarkPx      string `json:"markPx,omitempty"`
	Lever       string `json:"lever"`
	Pnl         string `json:"pnl"`
	RealizedPnl string `json:"realizedPnl"`
	Upl         string `json:"upl"`
	CTime       string `json:"cTime"`
	UTime       string `json:"uTime"`
}

type AnswPrivatePositions struct {
	Arg  ArgsSymbol         `json:"arg"`
	Data []PrivatePositions `json:"data"`
}

func (w *Ws) NewPrivatePositions(handler entity.WsHandlerPrivatePositions, errHandler ErrHandler) error {

	channel := "positions"
	h := func(message []byte) {
		mess := AnswPrivatePositions{}
		json.Unmarshal(message, &mess)
		for _, item := range mess.Data {
			d := ConvertWsPrivatePositions(item)
			handler(&d)
		}
	}
	w.mapsEvents = append(w.mapsEvents, MapsHandler{
		Channel:    channel,
		WsHandler:  h,
		ErrHandler: errHandler,
	})

	mess := struct {
		Op   string                 `json:"op"`
		Args []ArgsPrivatePositions `json:"args"`
	}{
		Op: "subscribe",
		Args: []ArgsPrivatePositions{
			{
				Channel:  channel,
				InstType: "SWAP",
			},
		},
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

func (w *Ws) NewPublicMarkPrice(symbols []string, handler entity.WsHandlerMarkPrice, errHandler ErrHandler) error {

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
	args := []ArgsSymbol{}
	for _, item := range symbols {
		args = append(args, ArgsSymbol{

			Channel: channel,
			InstId:  item,
		})
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
