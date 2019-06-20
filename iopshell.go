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

package main

import (
	"fmt"
	"os"

	_ "github.com/neonsea/iopshell/commands" // runs .Register() for each command
	"github.com/neonsea/iopshell/internal/shell"
)

func main() {
	if len(os.Args) == 1 {
		shell.Shell("")
	} else {
		err := shell.Shell(os.Args[1])
		if err != nil {
			fmt.Printf("Unable to run script '%s'\n", os.Args[1])
			os.Exit(-1)
		}
	}
}
