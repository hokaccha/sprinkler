package action

import "fmt"

func init() {
	RegisterAction("input", func() ActionRunner { return new(InputAction) })
}

type InputParams struct {
	Element string "element"
	Value   string "value"
}

type InputAction struct {
	ActionBase
}

func (a *InputAction) Run(params interface{}) error {
	p := &InputParams{}
	err := a.parseParams(params, p)

	if err != nil {
		return err
	}

	if p.Value == "" {
		return fmt.Errorf("value is not defined")
	}

	actionLog("input", "element=%s, value=%s", p.Element, p.Value)

	el, err := a.findElement(p.Element)

	if err != nil {
		return err
	}

	return el.SendKeys(p.Value)
}
