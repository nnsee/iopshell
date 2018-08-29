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

package cmd

import (
	"fmt"
	"strings"
)

var CommandList = make(map[string]Command)

type Command struct {
	Name        string
	UsageText   string
	Description string
	Action      func([]string)
	MinArg      int // Includes command name. 0 for unlimited
	MaxArg      int // Same as above
}

func (cmd *Command) Register() {
	CommandList[cmd.Name] = *cmd
}

func (cmd *Command) Usage() {
	fmt.Println("Usage:", cmd.UsageText)
}

func (cmd *Command) Execute(line string) {
	arg := strings.Split(line, " ")
	switch {
	case cmd.MaxArg != 0 && len(arg) > cmd.MaxArg:
		fmt.Println("Too many arguments")
		cmd.Usage()
	case cmd.MinArg != 0 && len(arg) < cmd.MinArg:
		fmt.Println("Too few arguments")
		cmd.Usage()
	default:
		cmd.Action(arg[1:])
	}
}
