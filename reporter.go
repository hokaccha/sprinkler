package main

import (
	"fmt"

	"github.com/visionmedia/go-debug"
)

var Debug = debug.Debug("screenplay")

func OK(message string, args ...interface{}) {
	fmt.Printf("\033[32mOK\033[0m - "+message+"\n", args...)
}

func NG(message string, args ...interface{}) {
	fmt.Printf("\033[31mNG\033[0m - "+message+"\n", args...)
}
