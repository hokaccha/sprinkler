package action

import "fmt"

func init() {
	RegisterAction("assert_selected", func() ActionRunner {
		return new(AssertSelectedAction)
	})
}

type AssertSelectedParams struct {
	Element  string "element"
	Expected *bool  "expected"
}

type AssertSelectedAction struct {
	ActionBase
}

func (a *AssertSelectedAction) Run(params interface{}) error {
	p := &AssertSelectedParams{}
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

	selected, err := el.IsSelected()

	if err != nil {
		return err
	}

	if expected == true {
		a.assert(selected, fmt.Sprintf("%s is selected", p.Element))
	} else {
		a.assert(!selected, fmt.Sprintf("%s is not selected", p.Element))
	}

	return nil
}
