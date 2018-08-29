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

var connect = cmd.Command{
	Name:        "connect",
	UsageText:   "connect [host]",
	Description: "Connects to [host]. If none specified, uses values from config",
	Action:      connectRun,
	MaxArg:      2,
}

func connectRun(param []string) {
	var addr string
	if len(param) == 0 {
		addr = setting.Host
	} else {
		addr = param[0]
	}
	setting.Cmd <- []string{"connect", addr}
}

func init() {
	connect.Register()
}
