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
	"reflect"
	"strings"

	"gitlab.com/c-/iopshell/internal/cmd"
	"gitlab.com/c-/iopshell/internal/setting"
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
	e := reflect.ValueOf(&setting.Vars.Opts).Elem()

	for i := 0; i < e.NumField(); i++ {
		f := e.Field(i)
		fmt.Printf("%s: %v\n", e.Type().Field(i).Name, f.Interface())
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
	if t == reflect.Invalid {
		fmt.Printf("Unknown setting '%s'\n", param[0])
		return
	}

	switch len(param) {
	case 1:
		fmt.Printf("%s: %v\n", param[0], opt.Interface())
	case 2:
		switch t {
		case reflect.Bool:
			val, err := toBool(param[1])
			if err != nil {
				fmt.Printf("'%s' not a boolean", param[1])
				return
			}
			opt.SetBool(val)
		case reflect.String:
			opt.SetString(param[1])
		}
	}
}

func getOption(option string) (reflect.Kind, *reflect.Value) {
	option = strings.Title(strings.ToLower(option))
	r := reflect.ValueOf(&setting.Vars.Opts)
	opt := reflect.Indirect(r).FieldByName(option)
	exp := opt.Kind()
	return exp, &opt
}

func init() {
	set.Register()
}
