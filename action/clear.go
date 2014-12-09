package action

func init() {
	RegisterAction("clear", func() ActionRunner { return new(ClearAction) })
}

type ClearParams struct {
	Element string `name:"url"`
}

type ClearAction struct {
	ActionBase
}

func (a *ClearAction) Run(params interface{}) error {
	p := &ClearParams{}
	err := a.parseParams(params, p)

	if err != nil {
		return err
	}

	actionLog("clear", "element=%s", p.Element)

	el, err := a.findElement(p.Element)

	if err != nil {
		return err
	}

	return el.Clear()
}
