package action

import "time"

func init() {
	RegisterAction("wait", func() ActionRunner { return new(WaitAction) })
}

type WaitParams struct {
	Delay int "delay"
}

type WaitAction struct {
	ActionBase
}

func (a *WaitAction) Run(params interface{}) error {
	p := &WaitParams{}
	err := a.parseParams(params, p)

	if err != nil {
		return err
	}

	actionLog("wait", "delay=%d", p.Delay)

	time.Sleep(time.Duration(p.Delay) * time.Millisecond)

	return nil
}
