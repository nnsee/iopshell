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

package commands

import (
	"gitlab.com/neonsea/iopshell/internal/cmd"
	"gitlab.com/neonsea/iopshell/internal/setting"
)

var auth = cmd.Command{
	Name:        "auth",
	UsageText:   "auth <user> <pass>",
	Description: "Authenticates as <user> with <pass>",
	Action:      authRun,
	MinArg:      3,
	MaxArg:      3,
}

func authRun(param []string) {
	// Authenticating is just another call
	setting.Vars.Conn.Call("session", "login", map[string]interface{}{"username": param[0],
		"password": param[1]})
}

func init() {
	auth.Register()
}
