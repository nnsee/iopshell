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
    C    *websocket.Conn
    Key  string
    User string
    Id   int
}

func (conn *Connection) Connect(addr string) {
    u := url.URL{Scheme: "ws", Host: addr, Path: "/"}

    con, _, err := dialer.Dial(u.String(), nil)
    if err != nil {
        fmt.Println(err)
    }
    conn.C = con
}

func (conn *Connection) Disconnect() {
    conn.C.Close()
    conn.C = nil
}

func (conn *Connection) Send(request interface{}) {
    conn.C.WriteJSON(request)
}

func (conn *Connection) Recv() interface{} {
    var response interface{}
    conn.C.ReadJSON(&response)
    return response
}

