package commands

import (
	"fmt"

	"github.com/golang/glog"
	"github.com/grassedge/go-online-judge-tools/services/atcoder"
	"github.com/urfave/cli"

	"net/url"
	"os"
)

func doDownload(c *cli.Context) error {
	urlStr := c.Args().First()
	directory := c.String("directory")
	if urlStr == "" {
		cli.ShowCommandHelp(c, "download")
		os.Exit(1)
	}
	if directory == "" {
		directory = "test"
	}

	url, err := url.Parse(urlStr)
	if err != nil {
		glog.Fatal(err)
	}

	// TODO: Support multiple services
	p, err := atcoder.NewAtCoderProblemFromUrl(url)
	if err != nil {
		glog.Fatal(err)
	}

	samples := p.DownloadSampleCases()
	os.Mkdir(directory, 0744)

	for _, s := range samples {
		{
			file, _ := os.Create(fmt.Sprintf("%s/%s.in", directory, s.InputName))
			file.Write([]byte(s.InputData))
		}
		{
			file, _ := os.Create(fmt.Sprintf("%s/%s.out", directory, s.OutputName))
			file.Write([]byte(s.OutputData))
		}
	}
	return nil
}
