package main

import (
	"os"
	"strings"

	"github.com/codegangsta/cli"
)

const helpTemplate = `
Usage: sprinkler [options] playscript.yml...

Options:
{{range .Flags}}  {{.}}
{{end}}`

func NewCliApp() *cli.App {
	cli.AppHelpTemplate = helpTemplate[1:]
	app := cli.NewApp()
	app.Name = "sprinkler"
	app.HideHelp = true
	app.Flags = []cli.Flag{
		browserFlag,
		remoteFlag,
		tagFlag,
		skipTagFlag,
		cli.HelpFlag,
	}
	app.Action = doMain

	return app
}

func doMain(c *cli.Context) {
	args := c.Args()

	if len(args) == 0 {
		cli.ShowAppHelp(c)
		os.Exit(1)
	}

	opts := &PlayerOpts{
		Browser:   c.String("browser"),
		RemoteUrl: c.String("remote"),
		Tags:      splitString(c.String("tags")),
		SkipTags:  splitString(c.String("skip-tags")),
	}

	statusCode := NewPlayer(opts).Run(args)

	os.Exit(statusCode)
}

var browserFlag = cli.StringFlag{
	Name:  "browser, b",
	Value: "firefox",
	Usage: "browser name",
}

var remoteFlag = cli.StringFlag{
	Name:  "remote, r",
	Value: "http://localhost:4444/wd/hub",
	Usage: "RemoteWebDriver server URL",
}

var tagFlag = cli.StringFlag{
	Name:  "tags, t",
	Value: "",
	Usage: "only run scenarios tagged with these values",
}

var skipTagFlag = cli.StringFlag{
	Name:  "skip-tags",
	Value: "",
	Usage: "only run scenarios whose tags do not match these values",
}

func splitString(str string) []string {
	if str == "" {
		return []string{}
	}

	s := strings.Split(str, ",")

	for key, val := range s {
		s[key] = strings.Trim(val, " ")
	}

	return s
}
