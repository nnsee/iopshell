package commands

import (
	"gitlab.com/neonsea/iopshell/internal/cmd"
	"gitlab.com/neonsea/iopshell/internal/setting"
)

var Auth = cmd.Command{
	Name:        "auth",
	UsageText:   "auth <user> <pass>",
	Description: "Authenticates as <user> with <pass>",
	Action:      auth,
	MinArg:      3,
	MaxArg:      3,
}

func auth(param []string) {
	setting.Vars.Conn.Call("session", "login", map[string]interface{}{"username": param[0],
		"password": param[1]})
}

func init() {
	Auth.Register()
}
