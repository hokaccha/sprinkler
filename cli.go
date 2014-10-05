package main

import (
	"flag"
	"fmt"
	"os"
)

const USAGE = `
  Usage: screenplay scenario.yml
  
  Options:
    --version, -v        print the version
    --help, -h           show help
`

var version = flag.Bool("version", false, "")

func init() {
	flag.BoolVar(version, "v", false, "")
}

func ParseCliArgs() string {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, USAGE[1:])
	}

	flag.Parse()

	if *version == true {
		fmt.Fprintf(os.Stderr, "screenplay version: %s\n", Version)
		os.Exit(0)
	}

	args := flag.Args()

	if len(args) == 0 {
		flag.Usage()
		os.Exit(2)
	}

	return args[0]
}
