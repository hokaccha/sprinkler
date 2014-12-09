package action

import "fmt"

func init() {
	RegisterAction("assert_dialog_text", func() ActionRunner {
		return new(AssertDialogTextAction)
	})
}

type AssertDialogTextParams struct {
	Equal   *string `name:"equal"`
	Contain *string `name:"contain"`
}

type AssertDialogTextAction struct {
	ActionBase
}

func (a *AssertDialogTextAction) Run(params interface{}) error {
	p := &AssertDialogTextParams{}
	err := a.parseParams(params, p)

	if err != nil {
		return err
	}

	text, err := a.Wd.AlertText()

	if err != nil {
		return err
	}

	subject := fmt.Sprintf("dialog text")

	if p.Equal != nil {
		a.assertEqual(subject, text, *p.Equal)
		return nil
	}

	if p.Contain != nil {
		a.assertContain(subject, text, *p.Contain)
		return nil
	}

	return fmt.Errorf("invalid parameters: \"equal\" or \"contain\" is required")
}
