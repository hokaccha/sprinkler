package action

import (
	"fmt"
)

func init() {
	RegisterAction("navigate", func() ActionRunner { return new(NavigateAction) })
}

type NavigateParams struct {
	Type string `name:"type"`
}

type NavigateAction struct {
	ActionBase
}

func (a *NavigateAction) Run(params interface{}) error {
	p := &NavigateParams{}
	err := a.parseParams(params, p)

	if err != nil {
		return err
	}

	switch p.Type {
	case "refresh":
		err = a.Wd.Refresh()
	case "back":
		err = a.Wd.Back()
	case "forword":
		err = a.Wd.Forward()
	default:
		return fmt.Errorf(`navigate parameter is invalid, you can use refresh, back or forword`)
	}

	actionLog("navigate", "type=%s", p.Type)

	if err != nil {
		return fmt.Errorf("Failed %s", p.Type)
	}

	return nil
}
