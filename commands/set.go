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
	"strconv"
	"strings"

	"github.com/neonsea/iopshell/internal/cmd"
	"github.com/neonsea/iopshell/internal/setting"
)

var set = cmd.Command{
	Name:        "set",
	UsageText:   "set [option] [value]",
	Description: "Sets [option] to [value]. For a list of options and their values, use just 'set' or 'set [option]'.",
	Action:      setRun,
	MinArg:      1,
	MaxArg:      3,
}

func printOpts() {
	for k, v := range setting.Vars.Opts {
		fmt.Printf("%s: %v\n", k, v.Val)
	}
}

func toBool(i string) (bool, error) {
	switch strings.ToLower(i) {
	case "y", "yes", "true", "1":
		return true, nil
	case "n", "no", "false", "0":
		return false, nil
	}
	return false, fmt.Errorf("toBool: unknown type '%s'", i)
}

func setRun(param []string) {
	if len(param) == 0 {
		printOpts()
		return
	}

	t, opt := getOption(param[0])
	if t == "unknown" {
		fmt.Printf("Unknown setting '%s'\n", param[0])
		return
	}

	switch len(param) {
	case 1:
		fmt.Printf("%s: %v (%s)\n", param[0], opt.Val, opt.Description)
	case 2:
		switch t {
		case "float64":
			val, err := strconv.ParseFloat(param[1], 64)
			if err != nil {
				fmt.Printf("'%s' not a float", param[1])
				return
			}
			setting.Vars.Set(param[0], val)
		case "bool":
			val, err := toBool(param[1])
			if err != nil {
				fmt.Printf("'%s' not a boolean", param[1])
				return
			}
			setting.Vars.Set(param[0], val)
		case "string":
			setting.Vars.Set(param[0], param[1])
		}
	}
}

func getOption(option string) (string, *setting.Opt) {
	kind := "unknown"
	o, ok := setting.Vars.Get(option)

	if !ok {
		return kind, &setting.Opt{}
	}
	switch o.Val.(type) {
	case float64:
		kind = "float64"
	case bool:
		kind = "bool"
	case string:
		kind = "string"
	}

	return kind, o
}

func init() {
	set.Register()
}
