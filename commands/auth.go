/*
   Copyright (c) 2018 Rasmus Moorats (nns)

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

package commands

import (
	"fmt"

	"github.com/nnsee/iopshell/internal/cmd"
	"github.com/nnsee/iopshell/internal/setting"
)

var auth = cmd.Command{
	Name:      "auth",
	UsageText: "auth [user [pass]]",
	Description: "Authenticates as [user] with [pass]. If none are specified, uses the values from settings.\n" +
		"\t\tIf only the user is specified, the shell prompts for a password.",
	Action: authRun,
	MinArg: 1,
	MaxArg: 3,
}

func authRun(param []string) {
	var user, pass string
	switch len(param) {
	case 0:
		user, _ = setting.Vars.GetS("user")
		pass, _ = setting.Vars.GetS("pass")
	case 1:
		user = param[0]
		passb, err := setting.Vars.Instance.ReadPassword("Password: ")
		if err != nil {
			fmt.Println("Unable to read password")
			return
		}
		pass = string(passb)
	case 2:
		user = param[0]
		pass = param[1]
	}

	// Authenticating is just another call (but key needs to be zeroed out)
	request := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      setting.Vars.Conn.ID,
		"method":  "call",
		"params": []interface{}{
			"00000000000000000000000000000000",
			"session",
			"login",
			map[string]interface{}{
				"username": user,
				"password": pass,
			},
		},
	}
	setting.Vars.Conn.Send(request)
}

func init() {
	auth.Register()
}
