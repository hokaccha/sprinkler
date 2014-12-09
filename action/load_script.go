package action

import (
	"fmt"
	"path/filepath"

	"github.com/hokaccha/sprinkler/utils"
)

func init() {
	RegisterAction("load_script", func() ActionRunner { return new(LoadScriptAction) })
}

type LoadScriptParams struct {
	Src string `name:"src"`
}

type LoadScriptAction struct {
	ActionBase
}

func (a *LoadScriptAction) Run(params interface{}) error {
	p := &LoadScriptParams{}
	err := a.parseParams(params, p)

	if err != nil {
		return err
	}

	if p.Src == "" {
		return fmt.Errorf("src is not defined")
	}

	actionLog("load_script", "src=%s", p.Src)
	srcPath := filepath.Join(a.BaseDir, p.Src)
	data, err := utils.ReadFile(srcPath)

	if err != nil {
		return err
	}

	_, err = a.Wd.ExecuteScript(string(data), nil)

	if err != nil {
		return err
	}

	return nil
}
