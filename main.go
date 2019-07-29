package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/golang/glog"
	"github.com/grassedge/go-online-judge-tools/commands"
	"github.com/urfave/cli"
)

const version = "0.0.1"

var revision = "HEAD"

func main() {
	flag.Parse()

	if err := newApp().Run(os.Args); err != nil {
		exitCode := 1
		if excoder, ok := err.(cli.ExitCoder); ok {
			exitCode = excoder.ExitCode()
		}
		glog.Error("error", err.Error())
		os.Exit(exitCode)
	}
}

func newApp() *cli.App {
	app := cli.NewApp()
	app.Name = "goj"
	app.Usage = "Tools for Online Judges"
	app.Version = fmt.Sprintf("%s (rev:%s)", version, revision)
	app.Author = "grassedge"
	app.Email = "grassedge@gmail.com"
	app.Commands = commands.Commands
	return app
}
