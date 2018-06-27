package shell

import (
    "github.com/chzyer/readline"
    "gitlab.com/neonsea/iopshell/internal/cmd"
)

var completer = readline.NewPrefixCompleter()

func updateCompleter() {
    commands := make([]string, len(cmd.CommandList))
    i := 0
    for c := range cmd.CommandList {
        commands[i] = c
        i++
    }
    for _, c := range commands {
        completer.Children = append(completer.Children, readline.PcItem(c))
    }
}