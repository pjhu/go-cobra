package step

import (
	"executor/cli/command/run/step/executor"
	"executor/cli/internal"
	"fmt"
)

type StartStep struct {
	Id         string
	Logs       string
}

func (o *StartStep) step1() {
	fmt.Println("STEP 1")
	fmt.Printf("Id: %s\n", o.Id)
	stdStr := internal.ExecShellCommand("echo", "step1")
	o.Logs += stdStr
}

func (o *StartStep) step2() {
	fmt.Println("STEP 2")
	stdStr := internal.ExecShellCommand("echo", "step2")
	o.Logs += stdStr
}

func (o *StartStep) step3()  {
	fmt.Println("STEP 3")
	rst := executor.Run()
	o.Logs += rst
}