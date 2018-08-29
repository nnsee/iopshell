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

	"gitlab.com/neonsea/iopshell/internal/cmd"
)

var Help = cmd.Command{
	Name:        "help",
	UsageText:   "help [command]",
	Description: "Prints some help info. If [command] is specified, prints info on that",
	Action:      help,
	MaxArg:      2,
}

func help(param []string) {
	if len(param) == 0 {
		fmt.Println("Available commands:")
		for cmd := range cmd.CommandList {
			fmt.Printf("%s\t\t", cmd)
		}
		fmt.Println("\n\nSee 'help [command]' for more info")
	} else {
		if cmd, ok := cmd.CommandList[param[0]]; ok {
			fmt.Printf("Name:\t\t%s\n", cmd.Name)
			fmt.Printf("Usage:\t\t%s\n", cmd.UsageText)
			fmt.Printf("Description:\t%s\n", cmd.Description)
		} else {
			fmt.Printf("Unknown command '%s'\n", param[0])
		}
	}
}

func init() {
	Help.Register()
}
