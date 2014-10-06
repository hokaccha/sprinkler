package main

import (
	"fmt"
	"log"
	"path/filepath"
	"reflect"

	"github.com/sourcegraph/go-selenium"
)

type Player struct {
	wd           selenium.WebDriver
	playscript   *Playscript
	successCount int
	failCount    int
}

func NewPlayer(playscript *Playscript) *Player {
	return &Player{playscript: playscript}
}

func (player *Player) Play() (statusCode int) {
	caps := selenium.Capabilities(map[string]interface{}{"browserName": "chrome"})
	url := "http://localhost:4444/wd/hub"
	wd, err := selenium.NewRemote(caps, url)

	if err != nil {
		log.Fatal(err)
	}

	player.wd = wd
	defer wd.Quit()

	err = player.PlayActions(player.playscript.Before)
	if err != nil {
		log.Println(err)
		return 1
	}

	err = player.PlayScenarios(player.playscript.Scenarios)
	if err != nil {
		log.Println(err)
		return 1
	}

	err = player.PlayActions(player.playscript.After)
	if err != nil {
		log.Println(err)
		return 1
	}

	fmt.Print("\n")
	if player.failCount == 0 {
		fmt.Println("Result: \033[32mSUCCESS\033[0m")
	} else {
		statusCode = 1
		fmt.Println("Result: \033[31mFAIL\033[0m")
	}

	fmt.Printf("Success: %d  Fail: %d\n", player.successCount, player.failCount)

	return statusCode
}

func (player *Player) LoadInclude(path string, scenarios *Scenarios) error {
	fullPath := filepath.Join(player.playscript.BaseDir, path)
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
			player.PlayScenarios(scenarios)
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
	fmt.Printf("\n## %s\n\n", scenario.Name)

	err := player.PlayActions(player.playscript.BeforeEach)
	if err != nil {
		return err
	}

	err = player.PlayActions(scenario.Actions)
	if err != nil {
		return err
	}

	err = player.PlayActions(player.playscript.AfterEach)
	if err != nil {
		return err
	}

	return nil
}

func (player *Player) PlayActions(actions Actions) error {
	for _, action := range actions {
		err := player.PlayAction(action)

		if err != nil {
			return err
		}
	}

	return nil
}

func (player *Player) PlayAction(action Action) error {
	_, isAssert := action["assert"]
	_, isCommand := action["command"]

	switch {
	case isAssert:
		return player.PlayAssertAction(action)
	case isCommand:
		return player.PlayCommandAction(action)
	default:
		log.Printf("Unknown action: %s", action["action"])
	}

	return nil
}

func (player *Player) PlayAssertAction(action Action) error {
	assert := action["assert"]
	methodName := fmt.Sprintf("Play%sAssert", ToCamelCase(assert))
	method := reflect.ValueOf(player).MethodByName(methodName)

	if !method.IsValid() {
		return fmt.Errorf("Unknown assert %s", assert)
	}

	result := method.Call([]reflect.Value{reflect.ValueOf(action)})
	err, _ := result[0].Interface().(error)

	return err
}

func (player *Player) PlayCommandAction(action Action) error {
	command := action["command"]
	methodName := fmt.Sprintf("Play%sCommand", ToCamelCase(command))
	method := reflect.ValueOf(player).MethodByName(methodName)

	if !method.IsValid() {
		return fmt.Errorf("Unknown command %s", command)
	}

	result := method.Call([]reflect.Value{reflect.ValueOf(action)})
	err, _ := result[0].Interface().(error)

	return err
}

func (player *Player) FindElement(selector string) (selenium.WebElement, error) {
	if selector == "" {
		return nil, fmt.Errorf("selector not defined")
	}

	return player.wd.FindElement(selenium.ByCSSSelector, selector)
}

func (player *Player) FindElements(selector string) ([]selenium.WebElement, error) {
	if selector == "" {
		return nil, fmt.Errorf("selector not defined")
	}

	return player.wd.FindElements(selenium.ByCSSSelector, selector)
}
