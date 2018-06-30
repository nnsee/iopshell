package connection

import (
    "github.com/gorilla/websocket"
    "fmt"
    "net/url"
    "time"
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

func (c *Connection) Connect(addr string) {
    u := url.URL{Scheme: "ws", Host: addr, Path: "/"}

    con, _, err := dialer.Dial(u.String(), nil)
    if err != nil {
        fmt.Println(err)
    }
    c.Ws = con
    if c.Key == "" {
        c.Key = "00000000000000000000000000000000"
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

func (c *Connection) Recv() interface{} {
    var response interface{}
    c.Ws.ReadJSON(&response)
    return response
}