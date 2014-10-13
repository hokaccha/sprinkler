package action

import "fmt"

func init() {
	RegisterAction("assert_exist", func() ActionRunner {
		return new(AssertExistAction)
	})
}

type AssertExistParams struct {
	Element  string "element"
	Expected *bool  "expected"
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

	expected := true
	if p.Expected != nil && *p.Expected == false {
		expected = false
	}

	// TODO: Check status code
	_, err = a.findElement(p.Element)
	ok := err == nil

	if expected == true {
		a.assert(ok, fmt.Sprintf("%s exists", p.Element))
	} else {
		a.assert(!ok, fmt.Sprintf("%s doesn't exist", p.Element))
	}

	return nil
}
