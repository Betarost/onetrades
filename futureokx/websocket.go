package futureokx

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/Betarost/onetrades/utils"
	"github.com/gorilla/websocket"
)

var (
	WebsocketTimeout   = time.Second * 25
	WebsocketKeepalive = true
)

type WsHandler func(message []byte)

type ErrHandler func(err error)

type MapsHandler struct {
	Channel    string
	WsHandler  WsHandler
	ErrHandler ErrHandler
}

type Ws struct {
	apiKey     string
	secretKey  string
	memo       string
	KeyType    string
	BaseURL    string
	Debug      bool
	Logger     *log.Logger
	c          *websocket.Conn
	doneC      chan struct{}
	stopC      chan struct{}
	mapsEvents []MapsHandler
}

type WsEventMessage struct {
	Arg   WsEvent `json:"arg"`
	Event string  `json:"event"`
}

type WsEvent struct {
	Channel string `json:"channel"`
	InstId  string `json:"instId"`
}

type ArgsLogin struct {
	ApiKey     string `json:"apiKey"`
	Passphrase string `json:"passphrase"`
	Timestamp  string `json:"timestamp"`
	Sign       string `json:"sign"`
}

func (w *Ws) signature(time int64) (sig string) {
	sf, err := utils.SignFunc(utils.KeyTypeHmacBase64)
	if err != nil {
		return sig
	}
	raw := fmt.Sprintf("%dGET/users/self/verify", time)
	sign, err := sf(w.secretKey, raw)
	if err != nil {
		return sig
	}
	return *sign
}

func (w *Ws) newConnect(endpoint string) error {
	dialer := websocket.Dialer{
		HandshakeTimeout:  45 * time.Second,
		EnableCompression: false,
	}
	c, _, err := dialer.Dial(endpoint, nil)
	if err != nil {
		return err
	}
	w.c = c
	w.c.SetReadLimit(655350)
	w.doneC = make(chan struct{})
	w.stopC = make(chan struct{})
	go func() {
		defer close(w.doneC)
		if WebsocketKeepalive {
			w.keepAlive(c, WebsocketTimeout)
		}
		silent := false
		go func() {
			select {
			case <-w.stopC:
				silent = true
			case <-w.doneC:
			}
			c.Close()
		}()
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				if !silent {
					for _, item := range w.mapsEvents {
						item.ErrHandler(err)
					}
				}
				return
			}
			// log.Println("=917bb7=", string(message))
			mess := WsEventMessage{}
			json.Unmarshal(message, &mess)
			if mess.Event == "" {
				for _, item := range w.mapsEvents {
					if item.Channel == mess.Arg.Channel {
						item.WsHandler(message)
					}
				}
			}
		}
	}()

	if w.apiKey != "" {
		timestamp := time.Now().Unix()
		sig := w.signature(timestamp)
		mess := struct {
			Op   string      `json:"op"`
			Args []ArgsLogin `json:"args"`
		}{
			Op: "login",
			Args: []ArgsLogin{
				{
					ApiKey:     w.apiKey,
					Passphrase: w.memo,
					Timestamp:  fmt.Sprintf("%d", timestamp),
					Sign:       sig,
				},
			},
		}

		if w.c != nil {
			w.c.WriteJSON(mess)
		}
	}
	return nil
}

func (w *Ws) keepAlive(c *websocket.Conn, timeout time.Duration) {
	ticker := time.NewTicker(timeout)
	c.SetPongHandler(func(msg string) error {
		if err := c.WriteMessage(websocket.PongMessage, []byte{}); err != nil {
			log.Println("=keepAlive WriteMessage err=", err)
			return err
		}
		return nil
	})
	go func() {
		defer ticker.Stop()
		for {
			if c == nil {
				return
			}
			<-ticker.C
			if c == nil {
				return
			}
			if err := c.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}()
}
