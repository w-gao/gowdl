package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/w-gao/gowdl/internal"
	"github.com/w-gao/gowdl/internal/domain"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "graph FILE",
		Short: "Generate the dependency graph of the input WDL document",
		Long:  "Generate the dependency graph of the input WDL document in JSON (by default)",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			uri := args[0]
			documents := []*domain.Document{}

			toVisit := []string{uri}
			visited := map[string]bool{}

			for len(toVisit) > 0 {
				uri = toVisit[0]
				toVisit = toVisit[1:]

				if _, ok := visited[uri]; ok {
					continue
				}
				visited[uri] = true

				builder, err := internal.NewWdlBuilder(uri)
				if err != nil {
					fmt.Printf("Error occurred when importing \"%v\": %v\n", uri, err)
					break
				}

				document, err := builder.ParseDocument()
				if err != nil {
					fmt.Printf("Failed to parse the WDL document. Reason: %v\n", err)
					break
				}

				documents = append(documents, document)

				// Recursively visit imports.
				for _, import_doc := range document.Imports {
					toVisit = append(toVisit, import_doc.AbsoluteUrl)
				}
			}

			// JSON output
			out, err := json.MarshalIndent(documents, "", "    ")
			if err != nil {
				fmt.Printf("Failed to marshal JSON: %v\n", err)
				return
			}

			fmt.Printf("%s\n", out)
		},
	})
}
