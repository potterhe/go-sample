package agent

import (
	"fmt"
	"os"
	"os/exec"
)

// NewAgent
func NewAgent(binaryPath string, args []string) *Agent {
	return &Agent{
		binaryPath: binaryPath,
		args:       args,
	}
}

type Agent struct {
	binaryPath string
	args       []string
}

func (a *Agent) Run() error {
	cmd := exec.Command(a.binaryPath, a.args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		return err
	}
	err := cmd.Wait()
	fmt.Println(err)
	return err
}
