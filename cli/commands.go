package cli

import (
	"executor/cli/command/run"
	"github.com/spf13/cobra"
)

func AddCommands(cmd *cobra.Command) {
	cmd.AddCommand(
		run.NewRunCommand(),
	)
}