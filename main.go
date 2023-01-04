package main

import (
	"github.com/w-gao/gowdl/internal/cmd"
)

func main() {
	// version, err := internal.GetVersion(os.Args[1])
	// if err != nil {
	// 	fmt.Printf("error: %s\n", err)
	// 	return
	// }

	// fmt.Printf("WDL version: '%s'\n", version)

	cmd.Execute()
}
