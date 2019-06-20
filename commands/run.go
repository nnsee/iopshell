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
	"fmt"

	"github.com/neonsea/iopshell/internal/cmd"
	"github.com/neonsea/iopshell/internal/setting"
)

var run = cmd.Command{
	Name:        "run",
	UsageText:   "run <path>",
	Description: "Runs script located at <path>",
	Action:      runRun,
	MinArg:      2,
	MaxArg:      2,
}

func runRun(path []string) {
	setting.ScriptIn <- []string{"run", path[0]}
	err := <-setting.ScriptRet // don't return until script has finished
	if err != nil {
		fmt.Printf("Failed to run script '%s'\n", path[0])
	}
}

func init() {
	run.Register()
}
