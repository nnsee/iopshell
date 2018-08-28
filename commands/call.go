package commands

import (
	"strings"

	"gitlab.com/neonsea/iopshell/internal/cmd"
	"gitlab.com/neonsea/iopshell/internal/setting"
	"gitlab.com/neonsea/iopshell/internal/textmutate"
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
		message := strings.Join(param[2:], " ")
		mmap, _ := textmutate.StrToMap(message)
		setting.Vars.Conn.Call(param[0], param[1], mmap)
	}
}

func init() {
	Call.Register()
}
