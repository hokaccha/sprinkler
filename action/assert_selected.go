package action

import "fmt"

func init() {
	RegisterAction("assert_selected", func() ActionRunner {
		return new(AssertSelectedAction)
	})
}

type AssertSelectedParams struct {
	Element  string "element"
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

	el, err := a.findElement(p.Element)

	if err != nil {
		return err
	}

	selected, err := el.IsSelected()

	if err != nil {
		return err
	}

	a.assert(selected, fmt.Sprintf("%s is selected", p.Element))

	return nil
}
