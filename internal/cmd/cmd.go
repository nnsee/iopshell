package cmd

import (
	"fmt"
	"strings"
)

var CommandList = make(map[string]Command)

type Command struct {
	Name        string
	UsageText   string
	Description string
	Action      func([]string)
	MinArg      int // Includes command name. 0 for unlimited
	MaxArg      int // Same as above
}

func (cmd *Command) Register() {
	CommandList[cmd.Name] = *cmd
}

func (cmd *Command) Usage() {
	fmt.Println("Usage:", cmd.UsageText)
}

func (cmd *Command) Execute(line string) {
	arg := strings.Split(line, " ")
	switch {
	case cmd.MaxArg != 0 && len(arg) > cmd.MaxArg:
		fmt.Println("Too many arguments")
		cmd.Usage()
	case cmd.MinArg != 0 && len(arg) < cmd.MinArg:
		fmt.Println("Too few arguments")
		cmd.Usage()
	default:
		cmd.Action(arg[1:])
	}
}
