package action

import (
	"fmt"
	"time"
)

func init() {
	RegisterAction("wait_for", func() ActionRunner { return new(WaitForAction) })
}

type WaitForParams struct {
	Element string "element"
	Timeout int    "delay"
}

type WaitForAction struct {
	ActionBase
}

func (a *WaitForAction) Run(params interface{}) error {
	p := &WaitForParams{}
	err := a.parseParams(params, p)

	if err != nil {
		return err
	}

	timeout := p.Timeout
	if timeout == 0 {
		timeout = 3000
	}

	actionLog("wait_for", "element=%s, timeout=%d", p.Element, timeout)

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
			return fmt.Errorf("Wait timeout: %s", p.Element)
		}

		el, err := a.findElement(p.Element)

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
