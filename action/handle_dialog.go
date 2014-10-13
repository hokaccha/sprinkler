package action

import "fmt"

func init() {
	RegisterAction("handle_dialog", func() ActionRunner { return new(HandleDialogAction) })
}

type HandleDialogParams struct {
	Type string "type"
}

type HandleDialogAction struct {
	ActionBase
}

func (a *HandleDialogAction) Run(params interface{}) error {
	p := &HandleDialogParams{}
	err := a.parseParams(params, p)

	if err != nil {
		return err
	}

	actionLog("handle_dialog", "type=%s", p.Type)

	switch p.Type {
	case "accept":
		return a.Wd.AcceptAlert()
	case "dismiss":
		return a.Wd.DismissAlert()
	default:
		return fmt.Errorf("invalid handle_dialog type: %s", p.Type)
	}
}
