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

type response struct {
	Id      int
	Result  []interface{}
	Jsonrpc string
}

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

func resultToStr(r int) string {
	switch r {
	case 0:
		return "Success"
	case 1:
		return "Invalid command"
	case 2:
		return "Invalid argument"
	case 3:
		return "Method not found"
	case 4:
		return "Not found"
	case 5:
		return "No data"
	case 6:
		return "Permission denied"
	case 7:
		return "Timeout"
	case 8:
		return "Not supported"
	case 10:
		return "Connection failed"
	}
	return "Unknown error"
}

func ParseResponse(r *response) (int, string, map[string]interface{}) {
	var result string
	var data map[string]interface{}
	rLen := len(r.Result)
	if rLen > 0 {
		result = resultToStr(int(r.Result[0].(float64)))
	}
	if rLen > 1 {
		data = r.Result[1].(map[string]interface{})
	}
	return rLen, result, data
}
