package action

import (
	"fmt"
	"strconv"
)

func init() {
	RegisterAction("assert_length", func() ActionRunner {
		return new(AssertLengthAction)
	})
}

type AssertLengthParams struct {
	Element string "element"
	Equal   *int   "equal"
}

type AssertLengthAction struct {
	ActionBase
}

func (a *AssertLengthAction) Run(params interface{}) error {
	p := &AssertLengthParams{}
	err := a.parseParams(params, p)

	if err != nil {
		return err
	}

	els, err := a.findElements(p.Element)

	if err != nil {
		return err
	}

	subject := fmt.Sprintf("%s length", p.Element)

	if p.Equal != nil {
		a.assertEqual(subject, strconv.Itoa(len(els)), strconv.Itoa(*p.Equal))
		return nil
	}

	return fmt.Errorf("invalid parameters: \"equal\" is required")
}
