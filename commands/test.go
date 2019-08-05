package commands

import (
	"fmt"

	"github.com/urfave/cli"
)

func doTest(c *cli.Context) error {
	fmt.Println("hogehoge")
	return nil
}
