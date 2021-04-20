/*
   Copyright (c) 2021 Rasmus Moorats (neonsea)

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

package textmutate

import (
	"fmt"

	"github.com/TylerBrock/colorjson"
)

var (
	Verbose   = false
	formatter = colorjson.NewFormatter()
)

func init() {
	formatter.Indent = 2
}

// Pprint pretty-prints a json input (interface{})
// In the future, this should be replaced with a custom function
func Pprint(input interface{}) string {
	out, err := formatter.Marshal(input)
	if err == nil {
		return string(out)
	}
	return ""
}

// Vprint prints stuff if verbose is enabled
func Vprint(input string) {
	if Verbose {
		fmt.Println(input)
	}
}
