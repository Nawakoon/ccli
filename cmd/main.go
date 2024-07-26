package main

import (
	"ccli/pkg/add"
	"ccli/pkg/delete"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	// command factory
	addCommand := add.NewCommandAdd()
	deleteCommand := delete.NewCommandDelete()

	// command register
	var rootCommand = &cobra.Command{
		Use: "ccli",
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true, // Disable shell completion command
		},
	}
	rootCommand.AddCommand(
		addCommand,
		deleteCommand,
	)

	// command execution
	if err := rootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}
