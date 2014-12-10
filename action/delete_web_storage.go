package action

import "fmt"

func init() {
	RegisterAction("delete_local_storage", func() ActionRunner {
		return new(DeleteLocalStorageAction)
	})
	RegisterAction("delete_session_storage", func() ActionRunner {
		return new(DeleteSessionStorageAction)
	})
}

type DeleteWebStorageParams struct {
	Key string `name:"key"`
}

type DeleteWebStorageAction struct {
	ActionBase
}

func (a *DeleteWebStorageAction) run(storage string, params interface{}) error {
	p := &DeleteWebStorageParams{}
	err := a.parseParams(params, p)

	if err != nil {
		return err
	}

	script := fmt.Sprintf(`window.%s.removeItem('%s')`, storage, p.Key)
	_, err = a.Wd.ExecuteScript(script, nil)

	if err != nil {
		return err
	}

	actionLog("delete_web_storage", "storage=%s, key=%s", storage, p.Key)

	return nil
}

type DeleteLocalStorageAction struct {
	DeleteWebStorageAction
}

func (a *DeleteLocalStorageAction) Run(params interface{}) error {
	return a.run("localStorage", params)
}

type DeleteSessionStorageAction struct {
	DeleteWebStorageAction
}

func (a *DeleteSessionStorageAction) Run(params interface{}) error {
	return a.run("sessionStorage", params)
}
