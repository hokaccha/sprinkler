package main

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/sourcegraph/go-selenium"
)

type Player struct {
	wd selenium.WebDriver
}

func NewPlayer() *Player {
	return &Player{}
}

func (player *Player) Play(scenarios []Scenario) {
	caps := selenium.Capabilities(map[string]interface{}{"browserName": "chrome"})
	url := "http://localhost:4444/wd/hub"
	wd, err := selenium.NewRemote(caps, url)

	if err != nil {
		log.Fatal(err)
	}

	player.wd = wd
	defer wd.Quit()

	for _, scenario := range scenarios {
		err := player.PlayScenario(scenario)

		if err != nil {
			log.Println(err)
		}
	}
}

func (player *Player) PlayScenario(scenario Scenario) error {
	Debug("Play: %s", scenario.Name)

	err := player.wd.Get(scenario.URL)

	if err != nil {
		log.Fatal(err)
	}

	for _, action := range scenario.Actions {
		_err := player.PlayAction(action)

		if _err != nil {
			log.Println(_err)
		}
	}

	return nil
}

func (player *Player) PlayAction(action map[string]string) error {
	_, isAssert := action["assert"]
	_, isTrigger := action["trigger"]
	_, isWait := action["wait"]

	switch {
	case isAssert:
		return player.PlayAssertAction(action)
	case isTrigger:
		return player.PlayTriggerAction(action)
	case isWait:
		return player.PlayWaitAction(action)
	default:
		log.Printf("Unknown action: %s", action["action"])
	}

	return nil
}

func (player *Player) PlayAssertAction(action map[string]string) error {
	assert := action["assert"]

	switch assert {
	case "equal_title":
		title, err := player.wd.Title()
		if err != nil {
			return err
		}
		if title == action["expected"] {
			OK("title text is '%s'", action["expected"])
		} else {
			NG("title text is not '%s'", action["expected"])
		}
	case "contain_text":
		el, err := player.wd.FindElement(selenium.ByCSSSelector, action["element"])
		if err != nil {
			return err
		}

		text, err := el.Text()

		if err != nil {
			return err
		}

		if strings.Contains(text, action["expected"]) {
			OK("%s text contains '%s'", action["element"], action["expected"])
		} else {
			NG("%s text doesn't contain '%s'", action["element"], action["expected"])
		}
	default:
		log.Printf("warn: Unknown assert %s", assert)
	}

	return nil
}

func (player *Player) PlayTriggerAction(action map[string]string) error {
	trigger := action["trigger"]
	el, err := player.wd.FindElement(selenium.ByCSSSelector, action["element"])

	if err != nil {
		return err
	}

	if trigger == "input" {
		el.SendKeys(action["value"])
	}

	if trigger == "click" {
		el.Click()
	}

	return nil
}

func (player *Player) PlayWaitAction(action map[string]string) error {
	ms, err := strconv.Atoi(action["wait"])

	if err != nil {
		log.Println("Warn: Wait time not found")
		return nil
	}

	Debug("Wait %dms", ms)

	time.Sleep(time.Duration(ms) * time.Millisecond)

	return nil
}
