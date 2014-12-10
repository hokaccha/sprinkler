package action

import "fmt"

func init() {
	RegisterAction("set_local_storage", func() ActionRunner {
		return new(SetLocalStorageAction)
	})
	RegisterAction("set_session_storage", func() ActionRunner {
		return new(SetSessionStorageAction)
	})
}

type SetWebStorageParams struct {
	Key   string `name:"key"`
	Value string `name:"value"`
}

type SetWebStorageAction struct {
	ActionBase
}

func (a *SetWebStorageAction) run(storage string, params interface{}) error {
	p := &SetWebStorageParams{}
	err := a.parseParams(params, p)

	if err != nil {
		return err
	}

	script := fmt.Sprintf(`window.%s.setItem('%s', '%s')`, storage, p.Key, p.Value)
	_, err = a.Wd.ExecuteScript(script, nil)

	if err != nil {
		return err
	}

	actionLog("set_web_storage", "storage=%s, key=%s, value=%s", storage, p.Key, p.Value)

	return nil
}

type SetLocalStorageAction struct {
	SetWebStorageAction
}

func (a *SetLocalStorageAction) Run(params interface{}) error {
	return a.run("localStorage", params)
}

type SetSessionStorageAction struct {
	SetWebStorageAction
}

func (a *SetSessionStorageAction) Run(params interface{}) error {
	return a.run("sessionStorage", params)
}
