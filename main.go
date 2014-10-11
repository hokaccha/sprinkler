package main

import (
	"log"
	"os"

	"github.com/sourcegraph/go-selenium"
	"github.com/visionmedia/go-debug"
)

var Debug = debug.Debug("sprinkler")

type Action map[string]string
type Actions []Action

type Scenario struct {
	Name    string   `name`
	Actions Actions  `actions`
	Include string   `include`
	Tags    []string `tags`
}

type Scenarios []Scenario

func init() {
	selenium.Log = nil
	log.SetFlags(log.Lshortfile)
}

func main() {
	app := NewCliApp()
	err := app.Run(os.Args)

	if err != nil {
		os.Exit(1)
	}
}
