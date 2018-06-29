package shell

import (
    "github.com/chzyer/readline"
    "gitlab.com/neonsea/iopshell/internal/cmd"
)

func (s *shellVars) UpdateCompleter() {
    s.Completer = *readline.NewPrefixCompleter()
    s.Completer.SetChildren(*new([]readline.PrefixCompleterInterface))

    commands := make([]string, len(cmd.CommandList))
    i := 0
    for c := range cmd.CommandList {
        commands[i] = c
        i++
    }
    for _, c := range commands {
        s.Completer.Children = append(s.Completer.Children, readline.PcItem(c))
    }
}