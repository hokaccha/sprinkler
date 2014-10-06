package main

import (
	"fmt"
	"log"
	"reflect"

	"github.com/sourcegraph/go-selenium"
)

type Player struct {
	wd selenium.WebDriver
	scenarioFile *ScenarioFile
}

func NewPlayer(scenarioFile *ScenarioFile) *Player {
	return &Player{scenarioFile: scenarioFile}
}

func (player *Player) Play() {
	caps := selenium.Capabilities(map[string]interface{}{"browserName": "chrome"})
	url := "http://localhost:4444/wd/hub"
	wd, err := selenium.NewRemote(caps, url)

	if err != nil {
		log.Fatal(err)
	}

	player.wd = wd
	defer wd.Quit()

	for _, scenario := range player.scenarioFile.Scenarios {
		err := player.PlayScenario(scenario)

		if err != nil {
			log.Println(err)
		}
	}
}

func (player *Player) PlayScenario(scenario Scenario) error {
	Debug("Play - %s", scenario.Name)

	for _, action := range scenario.Actions {
		_err := player.PlayAction(action)

		if _err != nil {
			log.Println(_err)
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
