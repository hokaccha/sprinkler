package main

import (
	"log"

	"github.com/sourcegraph/go-selenium"
)

func init() {
	selenium.Log = nil
	log.SetFlags(log.Lshortfile)
}

type Scenario struct {
	Name    string              `name`
	URL     string              `url`
	Actions []map[string]string `actions`
}

func main() {
	Debug("start main")

	scenarios, err := LoadSenarios("fixture/google_search.yml")

	if err != nil {
		log.Fatal(err)
	}

	player := NewPlayer()
	player.Play(scenarios)

	Debug("end main")
}
