package action

func init() {
	RegisterAction("submit", func() ActionRunner { return new(SubmitAction) })
}

type SubmitParams struct {
	Element string `name:"url"`
}

type SubmitAction struct {
	ActionBase
}

func (a *SubmitAction) Run(params interface{}) error {
	p := &SubmitParams{}
	err := a.parseParams(params, p)

	if err != nil {
		return err
	}

	actionLog("submit", "element=%s", p.Element)

	el, err := a.findElement(p.Element)

	if err != nil {
		return err
	}

	return el.Submit()
}
