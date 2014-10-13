package main

import (
	"log"
	"os"
	"strings"

	"github.com/codegangsta/cli"
)

const helpTemplate = `
Usage: sprinkler scenario.yml

Options:
  --tags=TAGS, -t TAGS  only run scenarios tagged with these values
  --skip-tags=TAGS      only run scenarios whose tags do not match these values
  --version, -v         print the version
  --help, -h            show help
`

func NewCliApp() *cli.App {
	cli.AppHelpTemplate = helpTemplate[1:]
	app := cli.NewApp()
	app.Name = "sprinkler"
	app.HideHelp = true
	app.Flags = []cli.Flag{
		tagFlag,
		skipTagFlag,
		cli.HelpFlag,
	}
	app.Action = action

	return app
}

func action(c *cli.Context) {
	args := c.Args()

	if len(args) == 0 {
		cli.ShowAppHelp(c)
		os.Exit(1)
	}

	playscript, err := NewPlayscript(args[0])

	if err != nil {
		log.Fatal(err)
	}

	opts := &PlayerOpts{
		Tags:     splitString(c.String("tags")),
		SkipTags: splitString(c.String("skip-tags")),
	}

	statusCode := NewPlayer(playscript, opts).Play()

	os.Exit(statusCode)
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
