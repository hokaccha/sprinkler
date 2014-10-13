package action

import "fmt"

func init() {
	RegisterAction("assert_visible", func() ActionRunner {
		return new(AssertVisibleAction)
	})
}

type AssertVisibleParams struct {
	Element  string "element"
	Expected *bool  "expected"
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

	expected := true
	if p.Expected != nil && *p.Expected == false {
		expected = false
	}

	el, err := a.findElement(p.Element)

	if err != nil {
		return err
	}

	visible, err := el.IsDisplayed()

	if err != nil {
		return err
	}

	if expected == true {
		a.assert(visible, fmt.Sprintf("%s is visible", p.Element))
	} else {
		a.assert(!visible, fmt.Sprintf("%s is not visible", p.Element))
	}

	return nil
}
