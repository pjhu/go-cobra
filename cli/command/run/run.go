package run

import (
	"executor/cli/command/run/step"
	"executor/cli/internal"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// NewRuCommand returns a cobra command for `start` command
func NewRunCommand() *cobra.Command {
	newCmd := &cobra.Command{
		Use:   "start",
		Short: "start",
		Example: "pjhu start ID",
		Run: func(cmd *cobra.Command, args []string) {
			internal.CheckArgsAmount(cmd, args, 1)
			server, err := cmd.Flags().GetString("server")
			fmt.Printf("server: %s\n", server)

			if err != nil {
				fmt.Println(err.Error())
				os.Exit(4)
			}
			run(args)
		},
	}
	return newCmd
}

func run(args []string) {

	// start run
	startTemplate := &step.StartStep{
		Id:         args[0],
	}
	startStep := step.IStartStep{
		Template: startTemplate,
	}
	startStep.Start()
}