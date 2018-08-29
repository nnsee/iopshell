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

package setting

import (
	"fmt"

	"github.com/chzyer/readline"
	"gitlab.com/neonsea/iopshell/internal/cmd"
	"gitlab.com/neonsea/iopshell/internal/connection"
)

var (
	Cmd = make(chan []string)
	In  = make(chan interface{})
	Out = make(chan interface{})
)

var (
	Host = "192.168.1.1"
)

type ShellVars struct {
	Conn      *connection.Connection
	Completer readline.PrefixCompleter
	Instance  *readline.Instance
}

func (s *ShellVars) UpdatePrompt() {
	var prompt string
	if s.Conn.Ws == nil {
		prompt = "\033[91miop\033[0;1m$\033[0m "
	} else {
		if s.Conn.User == "" {
			prompt = "\033[32miop\033[0;1m$\033[0m "
		} else {
			prompt = fmt.Sprintf("\033[32miop\033[0m %s\033[0;1m$\033[0m ", s.Conn.User)
		}
	}
	s.Instance.SetPrompt(prompt)
	s.Instance.Refresh()
}

func (s *ShellVars) UpdateCompleter(cmdlist map[string]cmd.Command) {
	s.Completer = *readline.NewPrefixCompleter()
	s.Completer.SetChildren(*new([]readline.PrefixCompleterInterface))

	commands := make([]string, len(cmdlist))
	i := 0
	for c := range cmdlist {
		commands[i] = c
		i++
	}
	for _, c := range commands {
		s.Completer.Children = append(s.Completer.Children, readline.PcItem(c))
	}
}

var Vars ShellVars
