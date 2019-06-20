/*
   Copyright (c) 2018 Rasmus Moorats (neonsea)

   This file is part of iopshell.

   iopshell is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   iopshell is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with iopshell. If not, see <https://www.gnu.org/licenses/>.
*/

package shell

import (
	"bufio"
	"os"
	"os/user"
	"time"

	"github.com/neonsea/iopshell/internal/setting"
	"github.com/neonsea/iopshell/internal/textmutate"
)

// GetRCFile fetches the absolute .ioprc file location
func GetRCFile() string {
	usr, err := user.Current()
	if err != nil {
		return ""
	}
	loc := usr.HomeDir + "/.ioprc"
	if _, err := os.Stat(loc); !os.IsNotExist(err) {
		return loc
	}
	return ""
}

// RunScript opens a .iop script, parses it and returns an error
func runScript(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		textmutate.Vprint(scanner.Text())
		parseLine(scanner.Text())
		del, _ := setting.Vars.GetF("script_delay")
		time.Sleep(time.Duration(del) * time.Second)
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
