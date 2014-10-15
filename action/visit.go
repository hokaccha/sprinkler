package action

import (
	"fmt"

	"github.com/hokaccha/sprinkler/utils"
)

func init() {
	RegisterAction("visit", func() ActionRunner { return new(VisitAction) })
}

type VisitParams struct {
	Url string "url"
}

type VisitAction struct {
	ActionBase
}

func (a *VisitAction) Run(params interface{}) error {
	p := &VisitParams{}
	err := a.parseParams(params, p)

	if err != nil {
		return err
	}

	if p.Url == "" {
		return fmt.Errorf("url is not defined")
	}

	url := utils.NormalizeUrl(p.Url, a.BaseDir)

	actionLog("visit", "url=%s", url)

	err = a.Wd.Get(url)

	if err != nil {
		return fmt.Errorf("Failed to visit: %s", url)
	}

	return nil
}
