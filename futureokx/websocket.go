package futureokx

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

var (
	WebsocketTimeout   = time.Second * 60
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
					mess := WsEventMessage{}
					json.Unmarshal(message, &mess)
					for _, item := range w.mapsEvents {
						if item.Channel == mess.Arg.Channel {
							item.ErrHandler(err)
						}
					}
				}
				return
			}
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
	return nil
}

func (w *Ws) keepAlive(c *websocket.Conn, timeout time.Duration) {
	ticker := time.NewTicker(timeout)
	c.SetPongHandler(func(msg string) error {
		if err := c.WriteMessage(websocket.PongMessage, []byte{}); err != nil {
			return err
		}
		return nil
	})
	go func() {
		defer ticker.Stop()
		for {
			if c == nil {
				log.Println("=bc159f=", 2)

				return
			}
			<-ticker.C
			if c == nil {
				log.Println("=bc159f=", 1)
				return
			}
			if err := c.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}()
}
