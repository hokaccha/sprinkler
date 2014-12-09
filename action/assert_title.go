package action

import "fmt"

func init() {
	RegisterAction("assert_title", func() ActionRunner {
		return new(AssertTitleAction)
	})
}

type AssertTitleParams struct {
	Equal   *string `name:"equal"`
	Contain *string `name:"contain"`
	Timeout int     `name:"timeout"`
}

type AssertTitleAction struct {
	ActionBase
}

func (a *AssertTitleAction) Run(params interface{}) error {
	p := &AssertTitleParams{}
	err := a.parseParams(params, p)

	if err != nil {
		return err
	}

	if p.Equal == nil && p.Contain == nil {
		return fmt.Errorf(`invalid parameters: "equal" or "contain" is required`)
	}

	return a.assertUntil(p.Timeout, func() error {
		url, err := a.Wd.Title()

		if err != nil {
			return err
		}

		subject := fmt.Sprintf("title")

		if p.Equal != nil {
			a.assertEqual(subject, url, *p.Equal)
		}

		if p.Contain != nil {
			a.assertContain(subject, url, *p.Contain)
		}

		return nil
	})
}
