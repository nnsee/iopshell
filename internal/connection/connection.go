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
    c   *websocket.Conn
    In  chan interface{}
    Out chan interface{}
}

func (conn *Connection) Connect(addr string) {
    u := url.URL{Scheme: "ws", Host: addr, Path: "/"}

    con, _, err := dialer.Dial(u.String(), nil)
    if err != nil {
        fmt.Println(err)
    }
    conn.c = con
}

func (conn *Connection) Send(request interface{}) {
    conn.c.WriteJSON(request)
}

func (conn *Connection) Recv() interface{} {
    var response interface{}
    conn.c.ReadJSON(&response)
    return response
}

