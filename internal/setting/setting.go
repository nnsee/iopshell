package setting

import (
    "gitlab.com/neonsea/iopshell/internal/connection"
)

type Status struct {
    Key       string
    Curr_user string
    Last_ret  int
    Last_id   int
    Auth      bool
    Conn      *connection.Connection
    In        chan interface{}
    Out       chan interface{}
}