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
	Element string `name:"element"`
	Equal   *int   `name:"equal"`
	Timeout int    `name:"timeout"`
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

	if p.Equal == nil {
		return fmt.Errorf(`invalid parameters: "equal" is required`)
	}

	return a.assertUntil(p.Timeout, func() error {
		els, err := a.findElements(p.Element)

		if err != nil {
			return err
		}

		subject := fmt.Sprintf("%s length", p.Element)

		if p.Equal != nil {
			a.assertEqual(subject, strconv.Itoa(len(els)), strconv.Itoa(*p.Equal))
		}

		return nil
	})
}
