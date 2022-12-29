package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "gowdl COMMAND [OPTIONS]",
	Short:   "A Workflow Description Language (WDL) parser and runner.",
	Version: "0.0.1a",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
