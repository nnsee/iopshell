# iopshell

__iopshell__ is a client written in Go, which allows the user to communicate with Inteno's IOPSYS devices, or other OpenWRT devices using the `owsd` server. Its primary goals are usability and to have a layer of abstraction handling all communications between the user and the device. __iopshell__ is also capable of executing simple scripts, which can be used to automate certain tasks.

## Installation
Before installing __iopshell__, make sure you've installed Go >1.10 from [here](https://golang.org/dl/) or using your distribution's package manager, and your `$GOPATH` and `$GOBIN` are properly set. You can then `go get` the dependencies and __iopshell__:
```sh
$ go get github.com/gorilla/websocket
$ go get github.com/chzyer/readline
$ go get gitlab.com/c-/iopshell
```
In the future, prebuilt binaries will be provided.

## Usage
Run a script:
```sh
$ $GOBIN/iopshell scripts/example.iop
```
Run the program as a shell:
```sh
$ $GOBIN/iopshell
```
You will be dropped to a shell prompt. Some example usage might look like this:
```
iop$ connect
iop$ auth admin admin
iop admin$ call uci get config:firewall
```

Note the custom paramater `config:firewall`. __iopshell__ uses a custom parser for call params. This is to prevent having to write out json-formatted text. For example, take the following json: 
```json
{"key1": {"key1-1": "value1-1", "key1-2": "value1-2"}, "key2": "value2"}
```
In __iopshell__, this can be formatted as:
```
key1:key1-1:value1-1,key1-2:value1-2;key2:value2
```
which is a lot more painless to write out. Of course, if you prefer, you can still use fully JSON-formatted text - either works.

See the `help` command for further info.

## Reporting problems
If you encounter an issue, unintended behaviour or a crash, [create a new issue](https://gitlab.com/c-/iopshell/issues). Don't forget to include:
1. Steps to reproduce the issue
2. Error log (if any)
3. Host platform info: Operating system, CPU architecture 

## Writing custom commands
Implementing custom commands is easy. Save your custom command as `commands/<command>.go`. The file needs to be a part of the package `commands` and has to import `internal/cmd`. To get an idea of the structure of a command, take a look at [auth.go](commands/auth.go). Don't forget to call `.Register()` in `init()`. Recompile and run.

## Built with
* [chyzer's readline](https://github.com/chzyer/readline) - for providing a shell interface
* [Gorilla WebSocket](https://github.com/gorilla/websocket) - for communicating with the device

## [License](LICENSE)
```
This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
```