package action

import (
	"fmt"
	"os"
	"os/exec"
)

func init() {
	RegisterAction("exec_sh", func() ActionRunner { return new(ExecShAction) })
}

type ExecShParams struct {
	Command string "command"
}

type ExecShAction struct {
	ActionBase
}

func (a *ExecShAction) Run(params interface{}) error {
	p := &ExecShParams{}
	err := a.parseParams(params, p)

	if err != nil {
		return err
	}

	if p.Command == "" {
		return fmt.Errorf("command is not defined")
	}

	actionLog("exec_sh", "command=%s", p.Command)

	cmd := exec.Command("sh", "-c", p.Command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
