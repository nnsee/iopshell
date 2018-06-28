package shell

import (
    "fmt"
    "io"
    "log"
    "strings"

    "github.com/chzyer/readline"
    "gitlab.com/neonsea/iopshell/internal/setting"
    "gitlab.com/neonsea/iopshell/internal/cmd"
    "gitlab.com/neonsea/iopshell/internal/connection"
)

func filterInput(r rune) (rune, bool) {
    switch r {
    case readline.CharCtrlZ:
        return r, false
    }
    return r, true
}

func UpdatePrompt(l *readline.Instance, s *setting.Status) {
    var prompt string
    if s.Conn == nil {
        prompt = "\033[91miop\033[0;1m$\033[0m "
    } else {
        if !s.Auth {
            prompt = "\033[32miop\033[0;1m$\033[0m "
        } else {
            prompt = fmt.Sprintf("\033[32miop\033[0m %s\033[0;1m$\033[0m ", s.Curr_user)
        }
    }
    l.SetPrompt(prompt)
    l.Refresh()
}

func Shell() {
    l, err := readline.NewEx(&readline.Config{
        Prompt:          "\033[91miop\033[0;1m$\033[0m ",
        HistoryFile:     "/tmp/iop.tmp",
        AutoComplete:    completer,
        InterruptPrompt: "^C",
        EOFPrompt:       "^D",

        HistorySearchFold:   true,
        FuncFilterInputRune: filterInput,
    })
    if err != nil {
        panic(err)
    }
    defer l.Close()

    updateCompleter()

    for {
        line, err := l.Readline()
        if err == io.EOF {
            break
        } else if err == readline.ErrInterrupt {
            continue
        }

        line     = strings.TrimSpace(line)
        command := strings.Split(line, " ")[0]
        if val, k := cmd.CommandList[command]; k {
            cmd.Execute(&val, line)
        } else {
            fmt.Printf("Unknown command '%s'\n", line)
        }
    }
}