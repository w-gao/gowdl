package cmd

import (
	"encoding/json"
	"fmt"
	"net/url"
	"path"
	"strings"

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
			// url := args[0]

			// builder, err := internal.NewWdlBuilder(url)
			// if err != nil {
			// 	fmt.Printf("%v\n", err)
			// 	return
			// }

			// document, err := builder.ParseDocument()
			// if err != nil {
			// 	fmt.Printf("Failed to parse the WDL document. Reason: %v\n", err)
			// 	return
			// }

			uri := args[0]
			uris := []string{uri}
			documents := []domain.Document{}
			visited := map[string]bool{}

			for len(uris) > 0 {
				uri = uris[0]
				uris = uris[1:]

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

				documents = append(documents, *document)

				// Recursively visit imports.
				for _, import_doc := range document.Imports {
					if strings.Contains(import_doc.Url, "://") {
						uris = append(uris, import_doc.Url)
						continue
					}

					u, err := url.Parse(uri)
					if err != nil {
						fmt.Printf("%v\n", err)
						continue
					}

					u.Path = path.Join(path.Dir(u.Path), import_doc.Url)
					uris = append(uris, u.String())
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
