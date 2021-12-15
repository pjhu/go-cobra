package command

import (
	"fmt"
	"github.com/spf13/cobra"
)

// NewVersionCommand returns a cobra command for `version` command
func NewVersionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Show the pjhu version information",
		Long:  `All software has versions. This is pjhu's`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("0.1")
		},
	}
	return cmd
}