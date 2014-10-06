package main

import (
	"fmt"
	"strconv"
	"time"
)

func playLog(command string, message string, args ...interface{}) {
	Debug("Play command "+command+" - " + message, args...)
}

func (player *Player) PlayVisitCommand(action Action) error {
	url := NormalizeUrl(action["url"], player.scenarioFile.BaseDir)
	playLog("visit", "url=%s", url)

	if url == "" {
		return fmt.Errorf("visit URL is not defined")
	}

	err := player.wd.Get(url)

	if err != nil {
		return err
	}

	return nil
}

func (player *Player) PlayWaitCommand(action Action) error {
	interval, ok := action["interval"]
	playLog("wait", "interval=%s", interval)

	if !ok {
		return fmt.Errorf("wait interval is not defined")
	}

	ms, err := strconv.Atoi(interval)

	if err != nil {
		return fmt.Errorf("wait interval is invalid: %s", interval)
	}

	time.Sleep(time.Duration(ms) * time.Millisecond)

	return nil
}

func (player *Player) PlayInputCommand(action Action) error {
	selector := action["element"]
	value := action["value"]
	playLog("input", "element=%s, value=%s", selector, value)

	if value == "" {
		return fmt.Errorf("value is not defiend")
	}

	el, err := player.FindElement(selector)

	if err != nil {
		return err
	}

	return el.SendKeys(value)
}

func (player *Player) PlayClickCommand(action Action) error {
	selector := action["element"]
	playLog("click", "element=%s", selector)

	el, err := player.FindElement(selector)

	if err != nil {
		return err
	}

	return el.Click()
}

func (player *Player) PlaySubmitCommand(action Action) error {
	selector := action["element"]
	playLog("submit", "element=%s", selector)

	el, err := player.FindElement(selector)

	if err != nil {
		return err
	}

	return el.Submit()
}

func (player *Player) PlayClearCommand(action Action) error {
	selector := action["element"]
	playLog("clear", "element=%s", selector)

	el, err := player.FindElement(selector)

	if err != nil {
		return err
	}

	return el.Clear()
}

func (player *Player) PlayAcceptAlertCommand(action Action) error {
	playLog("accept alert", "")
	return player.wd.AcceptAlert()
}

func (player *Player) PlayDismissAlertCommand(action Action) error {
	playLog("dismiss alert", "")
	return player.wd.DismissAlert()
}
