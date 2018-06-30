package commands

import (
    "gitlab.com/neonsea/iopshell/internal/setting"
    "gitlab.com/neonsea/iopshell/internal/cmd"
)

var Disconnect = cmd.Command {
    Name:        "disconnect",
    UsageText:   "disconnect",
    Description: "Disconnects from the currently connected host",
    Action:      disconnect,
    MaxArg:      1,
}

func disconnect(param []string) {
    setting.Cmd <- []string{"disconnect"}
}

func init() {
    Disconnect.Register()
}