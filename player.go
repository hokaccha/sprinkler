package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	. "github.com/hokaccha/sprinkler/action"
	. "github.com/hokaccha/sprinkler/utils"
	"github.com/sourcegraph/go-selenium"
)

type PlayerOpts struct {
	Tags      []string
	SkipTags  []string
	Browser   string
	RemoteUrl string
}

type Player struct {
	Wd           selenium.WebDriver
	Playscript   *Playscript
	SuccessCount int
	FailCount    int
	Opts         *PlayerOpts
}

func NewPlayer(playscript *Playscript, opts *PlayerOpts) *Player {
	return &Player{
		Playscript: playscript,
		Opts:       opts,
	}
}

func (player *Player) Play() (statusCode int) {
	caps := selenium.Capabilities(map[string]interface{}{"browserName": player.Opts.Browser})
	wd, err := selenium.NewRemote(caps, player.Opts.RemoteUrl)

	if err != nil {
		log.Fatal(err)
	}

	player.Wd = wd
	defer wd.Quit()

	err = player.PlayActions(player.Playscript.Before)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return 1
	}

	err = player.PlayScenarios(player.Playscript.Scenarios)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return 1
	}

	err = player.PlayActions(player.Playscript.After)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return 1
	}

	fmt.Print("\n")
	if player.FailCount == 0 {
		fmt.Println("Result: \033[32mSUCCESS\033[0m")
	} else {
		statusCode = 1
		fmt.Println("Result: \033[31mFAIL\033[0m")
	}

	fmt.Printf("Success: %d  Fail: %d\n", player.SuccessCount, player.FailCount)

	return statusCode
}

func (player *Player) LoadInclude(path string, scenarios *Scenarios) error {
	fullPath := filepath.Join(player.Playscript.BaseDir, path)
	return LoadYAML(fullPath, scenarios)
}

func (player *Player) PlayScenarios(scenarios Scenarios) error {
	var err error

	for _, scenario := range scenarios {
		if scenario.Include != "" {
			scenarios = Scenarios{}
			err = player.LoadInclude(scenario.Include, &scenarios)
			if err != nil {
				return err
			}

			err = player.PlayScenarios(scenarios)
			if err != nil {
				return err
			}
		} else {
			err = player.PlayScenario(scenario)

			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (player *Player) PlayScenario(scenario Scenario) error {
	if len(player.Opts.Tags) > 0 && !HasIntersection(scenario.Tags, player.Opts.Tags) {
		return nil
	}

	if HasIntersection(scenario.Tags, player.Opts.SkipTags) {
		return nil
	}

	fmt.Printf("\n## %s\n\n", scenario.Name)

	err := player.PlayActions(player.Playscript.BeforeEach)
	if err != nil {
		return err
	}

	err = player.PlayActions(scenario.Actions)
	if err != nil {
		return err
	}

	err = player.PlayActions(player.Playscript.AfterEach)
	if err != nil {
		return err
	}

	return nil
}

func (player *Player) PlayActions(actions Actions) error {
	for _, actionMap := range actions {
		for name, params := range actionMap {
			opts := &ActionOpts{
				Wd:      player.Wd,
				BaseDir: player.Playscript.BaseDir,
				Name:    name,
				Params:  params,
			}
			result, err := RunAction(opts)

			if err != nil {
				actionYAML, _ := ToYAML(actionMap)
				return fmt.Errorf("%s\n%s", Red("[Error] "+err.Error()), actionYAML)
			}

			player.HandleActionResult(result)
		}
	}

	return nil
}

func (player *Player) HandleActionResult(result *ActionResult) {
	if result == nil || result.IsAssert == false {
		return
	}

	if result.Successed == true {
		player.SuccessCount++
	} else {
		player.FailCount++
	}

	fmt.Println(result.Message)
}
