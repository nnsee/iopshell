package cmd

import (
    "fmt"
)

var Test = Command {
    Name:        "test",
    UsageText:   "TEST",
    Description: "Test",
    Action:      Testfunc,
}

func Testfunc(param []string) {
    fmt.Printf("test %s\n", param)
}