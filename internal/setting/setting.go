package setting

var (
    Cmd = make(chan []string)
    In  = make(chan interface{})
    Out = make(chan interface{})
)

var (
    Host = "192.168.1.1"
)