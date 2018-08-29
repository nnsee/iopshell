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
	// Cmd is a makeshift "API" to avoid imports
	Cmd = make(chan []string)
	// In receives incoming messages
	In = make(chan interface{})
	// Out houses outgoing messages
	Out = make(chan interface{})
)

var (
	// Host IP. In reality, this should be read from a file (if it exists)
	Host = "192.168.1.1"
)

// ShellVars house some important structs, so they can be accessed elsewhere
type ShellVars struct {
	Conn      *connection.Connection
	Completer readline.PrefixCompleter
	Instance  *readline.Instance
}

// UpdatePrompt refreshes the prompt and sets it according to current status
func (s *ShellVars) UpdatePrompt() {
	var prompt string
	if s.Conn.Ws == nil {
		// Not connected
		prompt = "\033[91miop\033[0;1m$\033[0m "
	} else {
		if s.Conn.User == "" {
			// Connected but not authenticated
			prompt = "\033[32miop\033[0;1m$\033[0m "
		} else {
			// Connected and authenticated
			prompt = fmt.Sprintf("\033[32miop\033[0m %s\033[0;1m$\033[0m ", s.Conn.User)
		}
	}
	s.Instance.SetPrompt(prompt)
	s.Instance.Refresh()
}

// UpdateCompleter adds commands registered with .Register() to the autocompleter
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

// Vars is an instance of ShellVars
var Vars ShellVars
