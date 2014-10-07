package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

const USAGE = `
Usage: sprinkler scenario.yml

Options:
  --tags=TAGS, -t TAGS  only run scenarios tagged with these values
  --skip-tags=TAGS      only run scenarios whose tags do not match these values
  --version, -v         print the version
  --help, -h            show help
`

var version = flag.Bool("version", false, "")
var tags = flag.String("tags", "", "")
var skipTags = flag.String("skip-tags", "", "")

func init() {
	flag.BoolVar(version, "v", false, "")
	flag.StringVar(tags, "t", "", "")
}

type Options struct {
	Tags     []string
	SkipTags []string
}

func ParseCliArgs() (string, *Options) {
	opts := &Options{}

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, USAGE[1:])
	}

	flag.Parse()

	if *version == true {
		fmt.Fprintf(os.Stderr, "sprinkler version: %s\n", Version)
		os.Exit(0)
	}

	opts.Tags = splitString(*tags)
	opts.SkipTags = splitString(*skipTags)

	args := flag.Args()

	if len(args) == 0 {
		flag.Usage()
		os.Exit(2)
	}

	return args[0], opts
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
