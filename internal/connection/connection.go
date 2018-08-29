/*
   Copyright (c) 2018 Rasmus Moorats (neonsea)

   This file is part of iopshell.

   iopshell is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   iopshell is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with iopshell. If not, see <https://www.gnu.org/licenses/>.
*/

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
