package action

import (
	"errors"
	"fmt"
)

func init() {
	RegisterAction("assert_local_storage", func() ActionRunner {
		return new(AssertLocalStorageAction)
	})
	RegisterAction("assert_session_storage", func() ActionRunner {
		return new(AssertSessionStorageAction)
	})
}

type AssertWebStorageParams struct {
	Key     string  `name:"key"`
	Equal   *string `name:"equal"`
	Contain *string `name:"contain"`
	Timeout int     `name:"timeout"`
}

type AssertWebStorageAction struct {
	ActionBase
}

func (a *AssertWebStorageAction) run(storage string, params interface{}) error {
	p := &AssertWebStorageParams{}
	err := a.parseParams(params, p)

	if err != nil {
		return err
	}

	if p.Equal == nil && p.Contain == nil {
		return fmt.Errorf(`invalid parameters: "equal" or "contain" is required`)
	}

	script := fmt.Sprintf(`return window.%s.getItem('%s')`, storage, p.Key)

	return a.assertUntil(p.Timeout, func() error {
		result, err := a.Wd.ExecuteScript(script, nil)

		if err != nil {
			return err
		}

		var value string
		switch v := result.(type) {
		case string:
			value = v
		case nil:
			value = ""
		default:
			return errors.New("Invalid result type")
		}

		subject := fmt.Sprintf("%s['%s']", storage, p.Key)

		if p.Equal != nil {
			a.assertEqual(subject, value, *p.Equal)
		}

		if p.Contain != nil {
			a.assertContain(subject, value, *p.Contain)
		}

		return nil
	})
}

type AssertLocalStorageAction struct {
	AssertWebStorageAction
}

func (a *AssertLocalStorageAction) Run(params interface{}) error {
	return a.run("localStorage", params)
}

type AssertSessionStorageAction struct {
	AssertWebStorageAction
}

func (a *AssertSessionStorageAction) Run(params interface{}) error {
	return a.run("sessionStorage", params)
}
