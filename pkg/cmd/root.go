package cmd

import (
	"github.com/otms61/my-cli-template/pkg/log"

	"github.com/spf13/cobra"
)

type rootOptions struct {
	debug bool
}

var rootOpts = rootOptions{}

var rootCmd = &cobra.Command{
	Short:             "A tool for Poc",
	Long:              `A tool for PoC`,
	Use:               "cmdtemplate",
	SilenceUsage:      false,
	PersistentPreRunE: initLogging,
}

func init() {
	rootCmd.PersistentFlags().BoolVar(
		&rootOpts.debug,
		"debug",
		false,
		"enable debug log",
	)

	addHello(rootCmd)
}

func initLogging(*cobra.Command, []string) error {
	err := log.InitLogger(rootOpts.debug)
	if err != nil {
		return err
	}
	return nil
}

// Execute builds the command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Logger.Fatalf("%+v\n", err)
	}
}
