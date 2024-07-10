package cmd

import (
	"github.com/otms61/my-cli-template/pkg/log"

	"github.com/spf13/cobra"
)

type helloOptions struct {
	name string
}

func hello(opts helloOptions) error {
	log.Logger.Debug("debug message")
	log.Logger.Infof("Hello %s!", opts.name)

	return nil
}

func addHello(parentCmd *cobra.Command) {
	opts := helloOptions{}
	cmd := &cobra.Command{
		Short: "hello",
		Long:  `hello`,
		Use:   "hello",
		RunE: func(_ *cobra.Command, args []string) error {
			return hello(opts)
		},
	}
	cmd.PersistentFlags().StringVar(
		&opts.name,
		"name",
		"bob",
		"name",
	)

	parentCmd.AddCommand(cmd)

}
