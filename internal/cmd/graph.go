package cmd

import (
	"encoding/json"
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

			document, err := builder.ParseDocument()

			if err != nil {
				fmt.Printf("Failed to parse the WDL document. Reason: %v\n", err)
				return
			}

			fmt.Printf("%v\n", document)
			out, err := json.MarshalIndent(document, "", "    ")
			if err == nil {
				fmt.Printf("%v\n", string(out))
			}

		},
	})
}
