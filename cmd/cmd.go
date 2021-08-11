package cmd

import (
	"github.com/mdcg/golang-autocaster-cli/config"
	"github.com/mdcg/golang-autocaster-cli/core"
	"github.com/spf13/cobra"
)

func AutocasterCMD() *cobra.Command {
	var verbose bool
	var rootCmd = &cobra.Command{Use: "autocaster"}

	var cmdRun = &cobra.Command{
		Use:   "run [CONFIG FILE PATH]",
		Short: "Start autocaster",
		Long:  `Starts autocaster processing from a given configuration file.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			macros := config.LoadConfig(args[0])
			core.LoadMacros(macros)
		},
	}
	cmdRun.Flags().BoolVarP(&verbose, "verbose", "v", true, "verbose mode")

	rootCmd.AddCommand(cmdRun)

	return rootCmd
}
