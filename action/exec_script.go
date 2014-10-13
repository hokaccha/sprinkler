package action

import "fmt"

func init() {
	RegisterAction("exec_script", func() ActionRunner { return new(ExecScriptAction) })
}

type ExecScriptParams struct {
	Script string "script"
}

type ExecScriptAction struct {
	ActionBase
}

func (a *ExecScriptAction) Run(params interface{}) error {
	p := &ExecScriptParams{}
	err := a.parseParams(params, p)

	if err != nil {
		return err
	}

	if p.Script == "" {
		return fmt.Errorf("script is not defined")
	}

	actionLog("exec_script", "script=%s", p.Script)

	_, err = a.Wd.ExecuteScript(p.Script, nil)

	if err != nil {
		return err
	}

	return nil
}
