// Copyright 2016 Google Inc. All Rights Reserved.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Command 127 provides a programmers debugging web server.
package main

import (
	"os"
	"path/filepath"

	"github.com/aliafshar/toylog"
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/googlesamples/onetwoseven"
)

var (
	app      = kingpin.New("127", "Programmer's test http server")
	hostFlag = app.Flag("host",
		"Host and port to start server on").
		Short('h').
		Default("127.0.0.1:8190").
		String()
	pathFlag = app.Flag("path",
		"path to serve.").
		Short('p').
		Default(".").
		String()
	verboseFlag = app.Flag("verbose",
		"Print debugging output.").
		Short('d').
		Bool()

)

const (
	logo1 = "━┃ ━━┃┏━┃"
	logo2 = " ┃ ┏━┛  ┃"
	logo3 = "━━┛━━┛  ┛"
)

func main() {
  app.Arg("path", "path to serve.").StringVar(pathFlag)
	_, err := app.Parse(os.Args[1:])
	if err != nil {
		toylog.Fatalln(err)
	}
  path, err := filepath.Abs(*pathFlag)
  if err != nil {
    toylog.Fatalln("path looks suspicious", *pathFlag, err)
  }
	cfg := &onetwoseven.Config{Path: path, HostPort: *hostFlag}
	toylog.Infoln(logo1)
	toylog.Infoln(logo2)
	toylog.Infoln(logo3)
	toylog.Infoln("starting", cfg)
	onetwoseven.Serve(cfg)
}
