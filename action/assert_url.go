package action

import "fmt"

func init() {
	RegisterAction("assert_url", func() ActionRunner {
		return new(AssertUrlAction)
	})
}

type AssertUrlParams struct {
	Equal   *string `name:"equal"`
	Contain *string `name:"contain"`
	Timeout int     `name:"timeout"`
}

type AssertUrlAction struct {
	ActionBase
}

func (a *AssertUrlAction) Run(params interface{}) error {
	p := &AssertUrlParams{}
	err := a.parseParams(params, p)

	if err != nil {
		return err
	}

	if p.Equal == nil && p.Contain == nil {
		return fmt.Errorf(`invalid parameters: "equal" or "contain" is required`)
	}

	return a.assertUntil(p.Timeout, func() error {
		url, err := a.Wd.CurrentURL()

		if err != nil {
			return err
		}

		subject := fmt.Sprintf("URL")

		if p.Equal != nil {
			a.assertEqual(subject, url, *p.Equal)
		}

		if p.Contain != nil {
			a.assertContain(subject, url, *p.Contain)
		}

		return nil
	})
}
