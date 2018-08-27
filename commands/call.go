package commands

import (
	"gitlab.com/neonsea/iopshell/internal/cmd"
	"gitlab.com/neonsea/iopshell/internal/setting"
)

var Call = cmd.Command{
	Name:        "call",
	UsageText:   "call <path> <method> [message]",
	Description: "Calls <method> from <path>. Additionally, [message] is passed to the call if set",
	Action:      call,
	MinArg:      3,
	MaxArg:      4,
}

func call(param []string) {
	if len(param) == 2 {
		setting.Vars.Conn.Call(param[0], param[1], make(map[string]interface{}))
	} else {
		setting.Vars.Conn.Call(param[0], param[1], *new(map[string]interface{}))
	}
}

func init() {
	Call.Register()
}
