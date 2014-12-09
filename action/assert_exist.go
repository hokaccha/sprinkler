package action

import "fmt"

func init() {
	RegisterAction("assert_exist", func() ActionRunner {
		return new(AssertExistAction)
	})
}

type AssertExistParams struct {
	Element  string `name:"element"`
}

type AssertExistAction struct {
	ActionBase
}

func (a *AssertExistAction) Run(params interface{}) error {
	p := &AssertExistParams{}
	err := a.parseParams(params, p)

	if err != nil {
		return err
	}

	// TODO: Check status code
	_, err = a.findElement(p.Element)
	ok := err == nil

	a.assert(ok, fmt.Sprintf("%s exists", p.Element))

	return nil
}
