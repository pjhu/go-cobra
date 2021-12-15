package executor

type IExecutor interface {
	ShellExecutor()
	FormatLog() string
}