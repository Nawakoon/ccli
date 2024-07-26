package delete

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// ccli delete 01-genpass

func NewCommandDelete() *cobra.Command {
	var command = &cobra.Command{
		Use:   "delete",
		Short: "Delete a command",
		Run:   runDeleteCommand,
	}

	return command
}

func runDeleteCommand(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		cmd.PrintErr("please provide a command to delete")
		os.Exit(1)
		return
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		cmd.PrintErr(err)
		os.Exit(1)
		return
	}

	configPath := homeDir + "/.ccli/"

	// check if command exists in config
	if _, err := os.Stat(configPath + args[0]); os.IsNotExist(err) {
		cmd.PrintErr("command does not exist")
		os.Exit(1)
		return
	}

	// delete command
	err = os.Remove(configPath + args[0])
	if err != nil {
		cmd.PrintErr(err)
		os.Exit(1)
		return
	}

	fmt.Println("command deleted successfully")
	os.Exit(0)
}
