package action

import "fmt"

func init() {
	RegisterAction("assert_value", func() ActionRunner {
		return new(AssertValueAction)
	})
}

type AssertValueParams struct {
	Element string  `name:"element"`
	Value   string  `name:"value"`
	Equal   *string `name:"equal"`
	Contain *string `name:"contain"`
	Timeout int     `name:"timeout"`
}

type AssertValueAction struct {
	ActionBase
}

func (a *AssertValueAction) Run(params interface{}) error {
	p := &AssertValueParams{}
	err := a.parseParams(params, p)

	if err != nil {
		return err
	}

	if p.Equal == nil && p.Contain == nil {
		return fmt.Errorf(`invalid parameters: "equal" or "contain" is required`)
	}

	return a.assertUntil(p.Timeout, func() error {
		value, err := a.getAttribute(p.Element, "value")

		if err != nil {
			return err
		}

		subject := fmt.Sprintf("%s value", p.Element)

		if p.Equal != nil {
			a.assertEqual(subject, value, *p.Equal)
		}

		if p.Contain != nil {
			a.assertContain(subject, value, *p.Contain)
		}

		return nil
	})
}
