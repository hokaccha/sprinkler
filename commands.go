package main

import (
	"fmt"
	"strconv"
	"time"
)

func playLog(command string, message string, args ...interface{}) {
	Debug("Play command "+command+" - "+message, args...)
}

func (player *Player) PlayVisitCommand(action Action) error {
	url := NormalizeUrl(action["url"], player.scenarioFile.BaseDir)

	if url == "" {
		return fmt.Errorf("visit URL is not defined")
	}

	err := player.wd.Get(url)

	if err != nil {
		return err
	}

	playLog("visit", "url=%s", url)

	return nil
}

func (player *Player) PlayWaitCommand(action Action) error {
	interval, ok := action["interval"]

	if !ok {
		return fmt.Errorf("wait interval is not defined")
	}

	ms, err := strconv.Atoi(interval)

	if err != nil {
		return fmt.Errorf("wait interval is invalid: %s", interval)
	}

	playLog("wait", "interval=%dms", ms)

	time.Sleep(time.Duration(ms) * time.Millisecond)

	return nil
}

func (player *Player) PlayInputCommand(action Action) error {
	selector := action["element"]
	value := action["value"]

	if value == "" {
		return fmt.Errorf("value is not defiend")
	}

	el, err := player.FindElement(selector)

	if err != nil {
		return err
	}

	playLog("input", "element=%s, value=%s", selector, value)

	return el.SendKeys(value)
}

func (player *Player) PlayClickCommand(action Action) error {
	selector := action["element"]

	el, err := player.FindElement(selector)

	if err != nil {
		return err
	}

	playLog("click", "element=%s", selector)

	return el.Click()
}

func (player *Player) PlaySubmitCommand(action Action) error {
	selector := action["element"]

	el, err := player.FindElement(selector)

	if err != nil {
		return err
	}

	playLog("submit", "element=%s", selector)

	return el.Submit()
}

func (player *Player) PlayClearCommand(action Action) error {
	selector := action["element"]

	el, err := player.FindElement(selector)

	if err != nil {
		return err
	}

	playLog("clear", "element=%s", selector)

	return el.Clear()
}
