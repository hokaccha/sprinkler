package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"
)

func playLog(command string, message string, args ...interface{}) {
	Debug("Play command "+command+" - "+message, args...)
}

func (player *Player) PlayVisitCommand(action Action) error {
	url := NormalizeUrl(action["url"], player.playscript.BaseDir)
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
	interval := action["interval"]
	playLog("wait", "interval=%s", interval)

	ms, err := strconv.Atoi(interval)

	if err != nil {
		return fmt.Errorf("wait interval is invalid: %s", interval)
	}

	time.Sleep(time.Duration(ms) * time.Millisecond)

	return nil
}

func (player *Player) PlayWaitForElementCommand(action Action) error {
	selector := action["element"]
	timeout, err := strconv.Atoi(action["timeout"])

	if err != nil {
		timeout = 5000
	}

	playLog("wait", "element=%s, timeout=%d", selector, timeout)

	// TODO: Use Channel
	interval := 200
	duration := time.Duration(interval) * time.Millisecond
	totalTime := 0
	wait := func() {
		totalTime += interval
		time.Sleep(duration)
	}

	for {
		if totalTime > timeout {
			return fmt.Errorf("Wait element timeout: %s", selector)
		}

		el, err := player.FindElement(selector)

		if err != nil {
			wait()
			continue
		}

		visible, err := el.IsDisplayed()

		if visible == false || err != nil {
			wait()
			continue
		}

		return nil
	}
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

func (player *Player) PlayExecScriptCommand(action Action) error {
	script := action["script"]
	src := action["src"]

	if src == "" {
		playLog("execute script", "script=%s", script)
	} else {
		playLog("execute script", "src=%s", src)
		srcPath := filepath.Join(player.playscript.BaseDir, src)
		data, err := ReadFile(srcPath)

		if err != nil {
			return err
		}

		script = string(data)
	}

	_, err := player.wd.ExecuteScript(script, nil)

	if err != nil {
		return err
	}

	return nil
}

func (player *Player) PlayExecShCommand(action Action) error {
	script := action["script"]
	playLog("execute sh", "script=%s", script)

	cmd := exec.Command("sh", "-c", script)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
