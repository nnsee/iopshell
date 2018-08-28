package connection

import (
	"fmt"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

var dialer = websocket.Dialer{
	Subprotocols:     []string{"ubus-json"},
	ReadBufferSize:   512,
	WriteBufferSize:  512,
	HandshakeTimeout: 15 * time.Second,
}

type Connection struct {
	Ws   *websocket.Conn
	Key  string
	User string
	Id   int
}

type response struct {
	Id      int
	Result  []interface{}
	Jsonrpc string
}

func (c *Connection) Connect(addr string) {
	u := url.URL{Scheme: "ws", Host: addr, Path: "/"}

	con, _, err := dialer.Dial(u.String(), nil)
	if err != nil {
		fmt.Println(err)
	} else {
		c.Ws = con
		if c.Key == "" {
			c.Key = "00000000000000000000000000000000"
		}
	}
}

func (c *Connection) Disconnect() {
	if c.Ws != nil {
		c.Ws.Close()
		c.Ws = nil
	}
}

func (c *Connection) Send(request interface{}) {
	if c.Ws != nil {
		c.Ws.WriteJSON(request)
		c.Id++
	}
}

func (c *Connection) Recv() response {
	if c.Ws != nil {
		var r response
		c.Ws.ReadJSON(&r)
		return r
	}
	return response{}
}
