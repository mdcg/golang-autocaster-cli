package cmd

import (
	"fmt"
	"time"

	"github.com/briandowns/spinner"
	"github.com/mdcg/golang-autocaster-cli/config"
	"github.com/mdcg/golang-autocaster-cli/core"
	"github.com/spf13/cobra"
)

const (
	CONFIG_FILE_PATH_INDEX = 0
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
			awaitTime := 10
			screenChangeAwaitingMsg := fmt.Sprintf("Waiting for %v seconds until you switch to the correct screen..", awaitTime)

			awaitingProcessing(screenChangeAwaitingMsg, awaitTime)

			macros := config.LoadMacroConfig(args[CONFIG_FILE_PATH_INDEX])
			core.LoadMacros(macros)
		},
	}
	cmdRun.Flags().BoolVarP(&verbose, "verbose", "v", true, "verbose mode")
	rootCmd.AddCommand(cmdRun)
	return rootCmd
}

func awaitingProcessing(msg string, sleepTime int) {
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Suffix = fmt.Sprintf(" %v", msg)
	s.FinalMSG = "\nDone!\n"

	s.Start()
	time.Sleep(time.Duration(sleepTime) * time.Second)
	s.Stop()
}
