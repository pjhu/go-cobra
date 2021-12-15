package executor

import (
	"fmt"
)

func Run() string {
	fmt.Println("RUN")
	golangExecutor := GolangExecutor{}
	golangExecutor.ShellExecutor()
	return golangExecutor.FormatLog()
}