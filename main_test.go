package main

import (
	"os"
	"os/exec"
	"testing"
)

var err error

func TestMain(t *testing.T) {
	liveTest := os.Getenv("LIVE_TEST")

	if liveTest == "" {
		return
	}

	cmd := exec.Command("go", "build")
	err = cmd.Run()

	if err != nil {
		t.Fatal(err)
	}

	cmd = exec.Command("sh", "-c", "./sprinkler --skip-tags notest ./example/*.yml")
	err = cmd.Run()

	if err != nil {
		t.Error(err)
	}
}
