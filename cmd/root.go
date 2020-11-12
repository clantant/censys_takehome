package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/clantant/mys_sweepsy/app/scan"
)

func Execute() {
	var rootCmd = &cobra.Command{
		Use:   "mys_sweepsy",
		Short: "mys_sweepsy is a mysql port scanner",
		Run: func(cmd *cobra.Command, args []string) {
			if err := scan.Run(cmd.Flags().Args()); err != nil {
				log.Printf("%+v", err)
				os.Exit(1)
			}
		},
	}

	if err := rootCmd.Execute(); err != nil {
		if err == pflag.ErrHelp {
			log.Printf("%+v", err)
		}

		os.Exit(1)
	}
}
