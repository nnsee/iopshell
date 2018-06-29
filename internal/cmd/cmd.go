package cmd

import (
    "strings"
    "fmt"
)

var CommandList = map[string]Command{
    "test": Test,
}

type Command struct {
    Name        string
    UsageText   string
    Description string
    Action      func([]string)
    MinArg      int
    MaxArg      int
}

func (cmd *Command) Usage() {
    fmt.Println("Usage: ", cmd.UsageText)
}

func (cmd *Command) Execute(line string) {
    arg := strings.Split(line, " ")[1:]
    cmd.Action(arg)
}