package shell

import (
    "fmt"
    "io"
    "strings"

    "github.com/chzyer/readline"
    "gitlab.com/neonsea/iopshell/internal/cmd"
    "gitlab.com/neonsea/iopshell/internal/connection"
)

type shellVars struct {
    Conn      *connection.Connection
    Completer readline.PrefixCompleter
    Instance  *readline.Instance
}

func filterInput(r rune) (rune, bool) {
    switch r {
    case readline.CharCtrlZ:
        return r, false
    }
    return r, true
}

func (s *shellVars) UpdatePrompt() {
    var prompt string
    if s.Conn.C == nil {
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

var Sv shellVars

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
    Sv.UpdatePrompt()
    Sv.UpdateCompleter()

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
            val.Execute(line)
        } else if command == "" {
            continue
        } else {
            fmt.Printf("Unknown command '%s'\n", line)
        }
    }
}