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
	"strconv"
	"time"

	"github.com/neonsea/iopshell/internal/cmd"
)

var wait = cmd.Command{
	Name:        "wait",
	UsageText:   "wait [time]",
	Description: "Wait for [time] seconds. If none specified, wait for half a second",
	Action:      waitRun,
	MaxArg:      2,
}

func waitRun(param []string) {
	t := 0.5
	if len(param) > 0 {
		tconv, err := strconv.ParseFloat(param[0], 32)
		if err == nil {
			t = tconv
		}
	}
	time.Sleep(time.Duration(t * 1000 * 1000 * 1000))
}

func init() {
	wait.Register()
}
