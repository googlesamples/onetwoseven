package main

import (
	"os"
	"path/filepath"

	"github.com/aliafshar/toylog"
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/aliafshar/onetwoseven"
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
