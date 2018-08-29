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

package shell

import (
	"fmt"
	"io"
	"strings"

	"github.com/chzyer/readline"
	"gitlab.com/neonsea/iopshell/internal/cmd"
	"gitlab.com/neonsea/iopshell/internal/connection"
	"gitlab.com/neonsea/iopshell/internal/setting"
	"gitlab.com/neonsea/iopshell/internal/textmutate"
)

func filterInput(r rune) (rune, bool) {
	switch r {
	case readline.CharCtrlZ:
		return r, false
	}
	return r, true
}

// Sv is just Vars for easy access
var Sv = &setting.Vars

func connectionHandler() {
	for {
		cmd := <-setting.Cmd
		switch cmd[0] {
		case "connect":
			Sv.Conn.Connect(cmd[1])
			go msgListener()
			Sv.UpdatePrompt()
		case "disconnect":
			Sv.Conn.Disconnect()
			Sv.UpdatePrompt()
		}
	}
}

func msgParser() {
	for {
		output := <-setting.Out
		Sv.Conn.Send(output)
	}
}

func msgListener() {
	for Sv.Conn.Ws != nil {
		response := Sv.Conn.Recv()
		if response.Jsonrpc != "" {
			rLen, err, rData := connection.ParseResponse(&response)
			fmt.Printf("\n%d: %s\n", response.ID, err)
			if rLen > 1 {
				fmt.Println(textmutate.Pprint(rData))
				if key, ok := rData["ubus_rpc_session"]; ok {
					// See if we have a session key
					Sv.Conn.Key = key.(string)
				}
				if data, ok := rData["data"]; ok {
					if user, ok := data.(map[string]interface{})["username"]; ok {
						// If we just logged in, we get our username
						Sv.Conn.User = user.(string)
					}
				}
			}
			Sv.UpdatePrompt()
		} else {
			return
		}
	}
	return
}

// Shell provides the CLI interface
func Shell() {
	l, err := readline.NewEx(&readline.Config{
		HistoryFile:     "/tmp/iop.tmp",
		AutoComplete:    &Sv.Completer,
		InterruptPrompt: "^C",
		EOFPrompt:       "^D",

		HistorySearchFold:   true,
		FuncFilterInputRune: filterInput,
	})
	if err != nil {
		panic(err)
	}

	Sv.Instance = l
	Sv.Conn = new(connection.Connection)
	defer l.Close()

	go connectionHandler()
	go msgParser()

	Sv.UpdatePrompt()
	Sv.UpdateCompleter(cmd.CommandList)

	for {
		line, err := l.Readline()
		if err == io.EOF {
			break
		} else if err == readline.ErrInterrupt {
			continue
		}

		line = strings.TrimSpace(line)
		command := strings.Split(line, " ")[0]
		if val, k := cmd.CommandList[command]; k {
			val.Execute(line)
		} else if command == "" {
			continue
		} else {
			fmt.Printf("Unknown command '%s'\n", line)
		}
	}
}
