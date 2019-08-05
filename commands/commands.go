package commands

import (
	"github.com/urfave/cli"
)

var Commands = []cli.Command{
	CommandDownload,
	CommandTest,
}

var CommandDownload = cli.Command{
	Name: "download",
	Usage: "Download sample cases",
	Description: `
supported services:
  AtCoder
`,
	Action: doDownload,
	Flags: []cli.Flag{
		cli.StringFlag{ Name: "directory, dir, d", Value: "test", Usage: "A directory name for test cases (default: test/)" },
	},
}

var CommandTest = cli.Command{
	Name: "test",
	Usage: "Test your code",
	Description: `
tips:
  You can do similar things with shell: e.g. for f in test/*.in ; do echo $f ; diff <(./a.out < $f) ${f/.in/.out} ; done
`,
	Action: doTest,
	Flags: []cli.Flag{
		cli.StringFlag{ Name: "command, c", Value: "./a.out", Usage: "your solution to be tested. (default: \"./a.out\")" },
		cli.StringFlag{ Name: "directory, d", Value: "test", Usage: "A directory name for test cases (default: test/)" },
	},
}
