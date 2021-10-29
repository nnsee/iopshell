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
	"encoding/json"
	"strings"

	"github.com/nnsee/iopshell/internal/cmd"
	"github.com/nnsee/iopshell/internal/setting"
	"github.com/nnsee/iopshell/internal/textmutate"
)

var call = cmd.Command{
	Name:        "call",
	UsageText:   "call <path> <method> [message]",
	Description: "Calls <method> from <path>. Additionally, [message] is passed to the call if set",
	Action:      callRun,
	MinArg:      3,
}

func callRun(param []string) {
	if len(param) == 2 {
		setting.Vars.Conn.Call(param[0], param[1], make(map[string]interface{}))
	} else {
		message := strings.Join(param[2:], " ")
		mmap := callMessage(message)
		setting.Vars.Conn.Call(param[0], param[1], mmap)
	}
}

func callMessage(message string) map[string]interface{} {
	out := make(map[string]interface{})
	err := json.Unmarshal([]byte(message), &out)
	if err != nil {
		out, _ := textmutate.StrToMap(message)
		return out
	}
	return out
}

func init() {
	call.Register()
}
