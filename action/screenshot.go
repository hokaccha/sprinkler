package action

import (
	"fmt"
	"io/ioutil"

	"github.com/hokaccha/sprinkler/utils"
)

func init() {
	RegisterAction("screenshot", func() ActionRunner { return new(ScreenshotAction) })
}

type ScreenshotParams struct {
	SavePath string `name:"save_path"`
}

type ScreenshotAction struct {
	ActionBase
}

func (a *ScreenshotAction) Run(params interface{}) error {
	p := &ScreenshotParams{}
	err := a.parseParams(params, p)

	if err != nil {
		return err
	}

	if p.SavePath == "" {
		return fmt.Errorf("save_path is not defined")
	}

	savePath := utils.AbsPath(p.SavePath, a.BaseDir)

	actionLog("screenshot", "save_path=%s", savePath)

	png, err := a.Wd.Screenshot()

	if err != nil {
		return err
	}

	return ioutil.WriteFile(savePath, png, 0644)
}
