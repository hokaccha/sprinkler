package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/hokaccha/sprinkler/action"
	"github.com/hokaccha/sprinkler/utils"
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
	PlayScript   *PlayScript
	SuccessCount int
	FailCount    int
	Opts         *PlayerOpts
}

func NewPlayer(opts *PlayerOpts) *Player {
	return &Player{
		Opts: opts,
	}
}

type Actions []map[string]interface{}

type Scenario struct {
	Name    string      `name`
	Actions Actions     `actions`
	Include string      `include`
	Tags    interface{} `tags`
}

type Scenarios []Scenario

type PlayScript struct {
	OrigPath   string
	FullPath   string
	BaseName   string
	BaseDir    string
	Scenarios  Scenarios `scenarios`
	Before     Actions   `before`
	After      Actions   `after`
	BeforeEach Actions   `before_each`
	AfterEach  Actions   `after_each`
}

func NewPlayScript(inputFilePath string) (*PlayScript, error) {
	fullPath, err := filepath.Abs(inputFilePath)

	if err != nil {
		return nil, err
	}

	playscript := &PlayScript{}

	playscript.OrigPath = inputFilePath
	playscript.FullPath = fullPath
	playscript.BaseDir = filepath.Dir(fullPath)
	playscript.BaseName = filepath.Base(fullPath)

	err = utils.LoadYAML(fullPath, &playscript)

	if err != nil {
		return nil, err
	}

	return playscript, nil
}

func (player *Player) Run(filePaths []string) (statusCode int) {
	caps := selenium.Capabilities(map[string]interface{}{
		"browserName": player.Opts.Browser,
	})

	wd, err := selenium.NewRemote(caps, player.Opts.RemoteUrl)

	if err != nil {
		log.Fatal(err)
	}

	player.Wd = wd
	defer wd.Quit()

	for _, filePath := range filePaths {
		err = player.Play(filePath)

		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			return 1
		}
	}

	if player.FailCount == 0 {
		fmt.Printf("\nResult: %s\n", utils.Green("SUCCESS"))
	} else {
		statusCode = 1
		fmt.Printf("\nResult: %s\n", utils.Red("FAIL"))
	}

	fmt.Printf("Success: %d  Fail: %d\n", player.SuccessCount, player.FailCount)

	return statusCode
}

func (player *Player) Play(filePath string) error {
	playscript, err := NewPlayScript(filePath)

	if err != nil {
		return fmt.Errorf("%s: %s", err.Error(), filePath)
	}

	utils.Debug("Run playscript: %s", playscript.FullPath)

	player.PlayScript = playscript

	err = player.PlayActions(player.PlayScript.Before)
	if err != nil {
		return err
	}

	err = player.PlayScenarios(player.PlayScript.Scenarios)
	if err != nil {
		return err
	}

	err = player.PlayActions(player.PlayScript.After)
	if err != nil {
		return err
	}

	return nil
}

func (player *Player) LoadInclude(path string, scenarios *Scenarios) error {
	fullPath := filepath.Join(player.PlayScript.BaseDir, path)
	return utils.LoadYAML(fullPath, scenarios)
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
	tags := utils.StringSlice(scenario.Tags)

	if len(player.Opts.Tags) > 0 && !utils.HasIntersection(tags, player.Opts.Tags) {
		return nil
	}

	if utils.HasIntersection(tags, player.Opts.SkipTags) {
		return nil
	}

	fmt.Printf("\n## %s\n\n", scenario.Name)

	err := player.PlayActions(player.PlayScript.BeforeEach)
	if err != nil {
		return err
	}

	err = player.PlayActions(scenario.Actions)
	if err != nil {
		return err
	}

	err = player.PlayActions(player.PlayScript.AfterEach)
	if err != nil {
		return err
	}

	return nil
}

func (player *Player) PlayActions(actions Actions) error {
	for _, actionMap := range actions {
		for name, params := range actionMap {
			opts := &action.ActionOpts{
				Wd:      player.Wd,
				BaseDir: player.PlayScript.BaseDir,
				Name:    name,
				Params:  params,
			}
			result, err := action.RunAction(opts)

			if err != nil {
				actionYAML, _ := utils.MarshalYAML(actionMap)
				format := `%s

Failed File
-----------
%s

Failed Action
-------------
%s`
				return fmt.Errorf(format, utils.Red("[Error] "+err.Error()), player.PlayScript.OrigPath, actionYAML)
			}

			player.HandleActionResult(result)
		}
	}

	return nil
}

func (player *Player) HandleActionResult(result *action.ActionResult) {
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
