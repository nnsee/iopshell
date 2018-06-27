package cmd

import (
	"fmt"
)

var Test = Command {
	Name: "test",
	Usage: "TEST",
	Description: "Test",
	Action: Testfunc,
}

func Testfunc(param []string) {
	fmt.Printf("test %s\n", param)
}