package shell

import (
    "fmt"
    "io"
    "strings"

    "github.com/chzyer/readline"
    "gitlab.com/neonsea/iopshell/internal/cmd"
    "gitlab.com/neonsea/iopshell/internal/connection"
    "gitlab.com/neonsea/iopshell/internal/setting"
)

func filterInput(r rune) (rune, bool) {
    switch r {
    case readline.CharCtrlZ:
        return r, false
    }
    return r, true
}

var Sv = &setting.Vars

func connectionHandler() {
    for {
        cmd := <-setting.Cmd
        switch cmd[0] {
        case "connect":
            Sv.Conn.Connect(cmd[1])
            Sv.UpdatePrompt()
        case "disconnect":
            Sv.Conn.Disconnect()
            Sv.UpdatePrompt()
        }
    }
}

func msgParser() {
    for {
        select {
        case input := <-setting.In:
            fmt.Println("Got", input)
        case output := <-setting.Out:
            Sv.Conn.Send(output)
        }
    }
}

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