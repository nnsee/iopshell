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
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
	"gitlab.com/c-/iopshell/internal/textmutate"
)

var dialer = websocket.Dialer{
	Subprotocols:     []string{"ubus-json"},
	ReadBufferSize:   512,
	WriteBufferSize:  512,
	HandshakeTimeout: 15 * time.Second,
}

// Connection houses the current connection's details
type Connection struct {
	Ws   *websocket.Conn
	Key  string
	User string
	ID   int
}

// Connect to the specified host
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

// Disconnect from the host
func (c *Connection) Disconnect() {
	if c.Ws != nil {
		c.Ws.Close()
		c.Ws = nil
	}
}

// Send the specified interface{} as json
func (c *Connection) Send(request interface{}) int {
	if c.Ws != nil {
		s, _ := json.Marshal(request)
		textmutate.Vprint(fmt.Sprintf("> %s", string(s)))
		c.Ws.WriteJSON(request)
		c.ID++
		return c.ID - 1
	}
	return -1
}

// Recv reads a json response into a response struct
func (c *Connection) Recv() Response {
	if c.Ws != nil {
		var r Response
		c.Ws.ReadJSON(&r)
		s, _ := json.Marshal(r)
		textmutate.Vprint(fmt.Sprintf("\n< %s", string(s)))
		return r
	}
	return Response{}
}
