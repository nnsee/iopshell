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
    Usage       string
    Description string
    Action      func([]string)
    MinArg      int
    MaxArg      int
}

func Usage(cmd *Command) {
    fmt.Println("Usage: ", cmd.Usage)
}

func Execute(cmd *Command, line string) {
    arg := strings.Split(line, " ")[1:]
    cmd.Action(arg)
}