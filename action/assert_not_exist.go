package action

import "fmt"

func init() {
	RegisterAction("assert_not_exist", func() ActionRunner {
		return new(AssertNotExistAction)
	})
}

type AssertNotExistParams struct {
	Element  string `name:"element"`
}

type AssertNotExistAction struct {
	ActionBase
}

func (a *AssertNotExistAction) Run(params interface{}) error {
	p := &AssertNotExistParams{}
	err := a.parseParams(params, p)

	if err != nil {
		return err
	}

	// TODO: Check status code
	_, err = a.findElement(p.Element)
	ok := err == nil

	a.assert(!ok, fmt.Sprintf("%s doesn't exist", p.Element))

	return nil
}
