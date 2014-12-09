package action

import "fmt"

func init() {
	RegisterAction("assert_hidden", func() ActionRunner {
		return new(AssertHiddenAction)
	})
}

type AssertHiddenParams struct {
	Element string `name:"element"`
	Timeout int    `name:"timeout"`
}

type AssertHiddenAction struct {
	ActionBase
}

func (a *AssertHiddenAction) Run(params interface{}) error {
	p := &AssertHiddenParams{}
	err := a.parseParams(params, p)

	if err != nil {
		return err
	}

	return a.assertUntil(p.Timeout, func() error {
		el, err := a.findElement(p.Element)

		if err != nil {
			return err
		}

		visible, err := el.IsDisplayed()

		if err != nil {
			return err
		}

		a.assert(!visible, fmt.Sprintf("%s is hidden", p.Element))

		return nil
	})
}
