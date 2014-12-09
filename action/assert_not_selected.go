package action

import "fmt"

func init() {
	RegisterAction("assert_not_selected", func() ActionRunner {
		return new(AssertNotSelectedAction)
	})
}

type AssertNotSelectedParams struct {
	Element  string "element"
}

type AssertNotSelectedAction struct {
	ActionBase
}

func (a *AssertNotSelectedAction) Run(params interface{}) error {
	p := &AssertNotSelectedParams{}
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

	a.assert(!selected, fmt.Sprintf("%s is not selected", p.Element))

	return nil
}

