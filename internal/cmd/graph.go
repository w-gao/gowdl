package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "graph FILE",
		Short: "Generate the dependency graph of the input WDL document",
		Long:  "Generate the dependency graph of the input WDL document in JSON (by default)",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("running... %v \n", args)
		},
	})
}
