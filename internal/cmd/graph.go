package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/w-gao/gowdl/internal"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "graph FILE",
		Short: "Generate the dependency graph of the input WDL document",
		Long:  "Generate the dependency graph of the input WDL document in JSON (by default)",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			url := args[0]

			builder, err := internal.NewWdlBuilder(url)
			if err != nil {
				fmt.Printf("%v\n", err)
				return
			}

			builder.ParseDocument()
		},
	})
}
