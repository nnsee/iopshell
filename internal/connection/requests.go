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

func (c *Connection) genUbusRequest(method, path, pmethod string, message map[string]interface{}) interface{} {
	request := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      c.Id,
		"method":  method,
		"params": []interface{}{
			c.Key,
			path,
			pmethod,
			message,
		},
	}

	return request
}

func (c *Connection) Call(path, method string, message map[string]interface{}) {
	request := c.genUbusRequest("call", path, method, message)
	c.Send(request)
}
