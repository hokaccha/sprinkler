package action

import "fmt"

func init() {
	RegisterAction("assert_css", func() ActionRunner {
		return new(AssertCssAction)
	})
}

type AssertCssParams struct {
	Element  string  `name:"element"`
	Property string  `name:"property"`
	Equal    *string `name:"equal"`
	Timeout  int     `name:"timeout"`
}

type AssertCssAction struct {
	ActionBase
}

func (a *AssertCssAction) Run(params interface{}) error {
	p := &AssertCssParams{}
	err := a.parseParams(params, p)

	if err != nil {
		return err
	}

	if p.Equal == nil {
		return fmt.Errorf(`invalid parameters: "equal" is required`)
	}

	return a.assertUntil(p.Timeout, func() error {
		el, err := a.findElement(p.Element)

		if err != nil {
			return err
		}

		value, err := el.CSSProperty(p.Property)

		if err != nil {
			return err
		}

		subject := fmt.Sprintf("%s %s", p.Element, p.Property)

		if p.Equal != nil {
			a.assertEqual(subject, value, *p.Equal)
		}

		return nil
	})
}
