package shell

import (
    "fmt"
    "io"
    "strings"

    "github.com/chzyer/readline"
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

func UpdatePrompt(l *readline.Instance, c *connection.Connection) {
    var prompt string
    if c.C == nil {
        prompt = "\033[91miop\033[0;1m$\033[0m "
    } else {
        if c.User == "" {
            prompt = "\033[32miop\033[0;1m$\033[0m "
        } else {
            prompt = fmt.Sprintf("\033[32miop\033[0m %s\033[0;1m$\033[0m ", c.User)
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

    conn := connection.Connection{}

    updateCompleter()

    for {
        line, err := l.Readline()
        UpdatePrompt(l, &conn)
        if err == io.EOF {
            break
        } else if err == readline.ErrInterrupt {
            continue
        }

        line     = strings.TrimSpace(line)
        command := strings.Split(line, " ")[0]
        if val, k := cmd.CommandList[command]; k {
            cmd.Execute(&val, line)
        } else if command == "" {
            continue
        } else {
            fmt.Printf("Unknown command '%s'\n", line)
        }
    }
}