package executor

import (
	"executor/cli/internal"
)

type GolangExecutor struct {
	stdStr string
}

func (exec *GolangExecutor) ShellExecutor() {
	exec.stdStr = internal.ExecShellCommand("/bin/sh", "-c", "echo shell ; echo command;")
}

func (exec *GolangExecutor) FormatLog() string {
	return exec.stdStr
}