package action

func init() {
	RegisterAction("click", func() ActionRunner { return new(ClickAction) })
}

type ClickParams struct {
	Element string "url"
}

type ClickAction struct {
	ActionBase
}

func (a *ClickAction) Run(params interface{}) error {
	p := &ClickParams{}
	err := a.parseParams(params, p)

	if err != nil {
		return err
	}

	actionLog("click", "element=%s", p.Element)

	el, err := a.findElement(p.Element)

	if err != nil {
		return err
	}

	return el.Click()
}
