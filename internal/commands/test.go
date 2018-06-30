package commands

import (
    "fmt"
    "gitlab.com/neonsea/iopshell/internal/cmd"
)

var Test = cmd.Command {
    Name:        "test",
    UsageText:   "TEST",
    Description: "Test",
    Action:      Testfunc,
}

func Testfunc(param []string) {
    fmt.Printf("test %s\n", param)
}

func init() {
    Test.Register()
}