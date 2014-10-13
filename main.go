package main

import (
	"log"
	"os"

	"github.com/sourcegraph/go-selenium"
)

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
