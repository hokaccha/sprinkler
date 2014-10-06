package main

import (
	"log"
	"os"

	"github.com/sourcegraph/go-selenium"
	"github.com/visionmedia/go-debug"
)

var Debug = debug.Debug("sprinkler")

func init() {
	selenium.Log = nil
	log.SetFlags(log.Lshortfile)
}

func main() {
	filePath := ParseCliArgs()
	playscript, err := NewPlayscript(filePath)

	if err != nil {
		log.Fatal(err)
	}

	statusCode := NewPlayer(playscript).Play()

	os.Exit(statusCode)
}
