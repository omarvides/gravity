/*
Copyright 2018 Gravitational, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package common

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/gravitational/gravity/lib/constants"

	teleutils "github.com/gravitational/teleport/lib/utils"
	"github.com/gravitational/trace"
	"golang.org/x/crypto/ssh/terminal"
	"gopkg.in/alecthomas/kingpin.v2"
)

// ReadUserPass reads a username and password from the console
func ReadUserPass() (string, string, error) {
	fmt.Printf("username: ")
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadSlice('\n')
	if err != nil {
		return "", "", trace.Wrap(err)
	}
	fmt.Printf("password: ")
	password, err := terminal.ReadPassword(0)
	if err != nil {
		return "", "", trace.Wrap(err)
	}
	return strings.TrimSpace(string(line)), strings.TrimSpace(string(password)), nil
}

// GetReader returns the reader for the provided file or stdin if no filename
// was provided
func GetReader(filename string) (io.ReadCloser, error) {
	if filename == "" {
		return ioutil.NopCloser(os.Stdin), nil
	}
	return teleutils.OpenFile(filename)
}

// Format is the CLI parser for output format flag
func Format(s kingpin.Settings) *constants.Format {
	var f constants.Format
	s.SetValue(&f)
	return &f
}
