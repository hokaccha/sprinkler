package action

import "fmt"

func init() {
	RegisterAction("assert_css", func() ActionRunner {
		return new(AssertCssAction)
	})
}

type AssertCssParams struct {
	Element  string  "element"
	Property string  "property"
	Equal    *string "equal"
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
		return nil
	}

	return fmt.Errorf("invalid parameters: \"equal\" is required")
}
