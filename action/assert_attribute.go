package action

import "fmt"

func init() {
	RegisterAction("assert_attribute", func() ActionRunner {
		return new(AssertAttributeAction)
	})
}

type AssertAttributeParams struct {
	Element   string  `name:"element"`
	Attribute string  `name:"attribute"`
	Equal     *string `name:"equal"`
	Contain   *string `name:"contain"`
	Present   *string `name:"present"`
	Timeout   int     `name:"timeout"`
}

type AssertAttributeAction struct {
	ActionBase
}

func (a *AssertAttributeAction) Run(params interface{}) error {
	p := &AssertAttributeParams{}
	err := a.parseParams(params, p)

	if err != nil {
		return err
	}

	if p.Equal == nil && p.Contain == nil && p.Present == nil {
		return fmt.Errorf(`invalid parameters: "equal" or "contain" or "present" is required`)
	}

	return a.assertUntil(p.Timeout, func() error {
		value, err := a.getAttribute(p.Element, p.Attribute)

		if err != nil {
			return err
		}

		subject := fmt.Sprintf("%s[%s]", p.Element, p.Attribute)

		if p.Equal != nil {
			a.assertEqual(subject, value, *p.Equal)
		}

		if p.Contain != nil {
			a.assertContain(subject, value, *p.Contain)
		}

		if p.Present != nil {
			a.assertPresent(subject, value, *p.Present)
		}

		return nil
	})
}
