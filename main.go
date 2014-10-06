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

func main() {
	filePath := ParseCliArgs()
	scenarioFile, err := NewSenarioFile(filePath)

	if err != nil {
		log.Fatal(err)
	}

	player := NewPlayer(scenarioFile)
	player.Play()
}
