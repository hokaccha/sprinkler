package action

import "fmt"

func init() {
	RegisterAction("assert_visible", func() ActionRunner {
		return new(AssertVisibleAction)
	})
}

type AssertVisibleParams struct {
	Element  string "element"
}

type AssertVisibleAction struct {
	ActionBase
}

func (a *AssertVisibleAction) Run(params interface{}) error {
	p := &AssertVisibleParams{}
	err := a.parseParams(params, p)

	if err != nil {
		return err
	}

	el, err := a.findElement(p.Element)

	if err != nil {
		return err
	}

	visible, err := el.IsDisplayed()

	if err != nil {
		return err
	}

	a.assert(visible, fmt.Sprintf("%s is visible", p.Element))

	return nil
}
