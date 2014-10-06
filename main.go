package main

import (
	"log"

	"github.com/sourcegraph/go-selenium"
	"github.com/visionmedia/go-debug"
)

var Debug = debug.Debug("screenplay")

func init() {
	selenium.Log = nil
	log.SetFlags(log.Lshortfile)
}

type Action map[string]string

type Scenario struct {
	Name    string   `name`
	URL     string   `url`
	Actions []Action `actions`
}

func main() {
	filePath := ParseCliArgs()
	scenarios, err := LoadSenarios(filePath)

	if err != nil {
		log.Fatal(err)
	}

	player := NewPlayer()
	player.Play(scenarios)
}
