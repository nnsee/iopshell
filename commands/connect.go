package commands

import (
    "gitlab.com/neonsea/iopshell/internal/cmd"
    "gitlab.com/neonsea/iopshell/internal/setting"
)

var Connect = cmd.Command{
    Name:        "connect",
    UsageText:   "connect [host]",
    Description: "Connects to [host]. If none specified, uses values from config",
    Action:      connect,
    MaxArg:      2,
}

func connect(param []string) {
    var addr string
    if len(param) == 0 {
        addr = setting.Host
    } else {
        addr = param[0]
    }
    setting.Cmd <- []string{"connect", addr}
}

func init() {
    Connect.Register()
}
