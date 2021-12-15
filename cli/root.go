package cli

import (
	"executor/cli/internal/constant"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

func newRootCommand() *cobra.Command {

	rootCmd := &cobra.Command{
		Use:              "pjhu COMMAND [ARG...]",
		Short:            "A self-sufficient runtime for pjhu platform",
		SilenceUsage:     true,
		SilenceErrors:    true,
		TraverseChildren: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				cmd.HelpFunc()(cmd, args)
				return nil
			}
			return fmt.Errorf("pjhu: '%s' is not a pjhu cli.\nSee 'pjhu --help'", args[0])
		},
		Version:               fmt.Sprintf("%s", "0.1"),
	}

	AddCommands(rootCmd)
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	return rootCmd
}

func initCmd(rootCmd *cobra.Command) {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringP("server", "s", constant.ServerDefaultUrl, "server URL.")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cmd.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cmd" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cmd")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd := newRootCommand()
	initCmd(rootCmd)
	cobra.CheckErr(rootCmd.Execute())
}
