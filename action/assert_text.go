package action

import "fmt"

func init() {
	RegisterAction("assert_text", func() ActionRunner {
		return new(AssertTextAction)
	})
}

type AssertTextParams struct {
	Element string  `name:"element"`
	Equal   *string `name:"equal"`
	Contain *string `name:"contain"`
}

type AssertTextAction struct {
	ActionBase
}

func (a *AssertTextAction) Run(params interface{}) error {
	p := &AssertTextParams{}
	err := a.parseParams(params, p)

	if err != nil {
		return err
	}

	el, err := a.findElement(p.Element)

	if err != nil {
		return err
	}

	text, err := el.Text()

	if err != nil {
		return err
	}

	subject := fmt.Sprintf("%s text", p.Element)

	if p.Equal != nil {
		a.assertEqual(subject, text, *p.Equal)
		return nil
	}

	if p.Contain != nil {
		a.assertContain(subject, text, *p.Contain)
		return nil
	}

	return fmt.Errorf("invalid parameters: \"equal\" or \"contain\" is required")
}
