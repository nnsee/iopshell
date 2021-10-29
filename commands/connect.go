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
	"github.com/nnsee/iopshell/internal/cmd"
	"github.com/nnsee/iopshell/internal/setting"
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
		addr, _ = setting.Vars.GetS("host")
	} else {
		addr = param[0]
	}
	setting.ConnChannel <- []string{"connect", addr}
}

func init() {
	connect.Register()
}
