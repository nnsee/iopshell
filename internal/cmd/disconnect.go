package cmd

import (
    "gitlab.com/neonsea/iopshell/internal/setting"
)

var Disconnect = Command {
    Name:        "disconnect",
    UsageText:   "disconnect",
    Description: "Disconnects from the currently connected host",
    Action:      disconnect,
    MaxArg:      1,
}

func disconnect(param []string) {
    setting.Cmd <- []string{"disconnect"}
}