package action

import "fmt"

func init() {
	RegisterAction("assert_attribute", func() ActionRunner {
		return new(AssertAttributeAction)
	})
}

type AssertAttributeParams struct {
	Element   string  "element"
	Attribute string  "attribute"
	Equal     *string "equal"
	Contain   *string "contain"
	Present   *string "present"
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

	value, err := a.getAttribute(p.Element, p.Attribute)

	if err != nil {
		return err
	}

	subject := fmt.Sprintf("%s[%s]", p.Element, p.Attribute)

	if p.Equal != nil {
		a.assertEqual(subject, value, *p.Equal)
		return nil
	}

	if p.Contain != nil {
		a.assertContain(subject, value, *p.Contain)
		return nil
	}

	if p.Present != nil {
		a.assertPresent(subject, value, *p.Present)
		return nil
	}

	return fmt.Errorf("invalid parameters: \"equal\" or \"contain\" or \"present\" is required")
}
