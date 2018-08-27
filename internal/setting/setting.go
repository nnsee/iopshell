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
