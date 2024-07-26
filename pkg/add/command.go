package add

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

// ccli add --file genpass.py --name 01-genpass

func NewCommandAdd() *cobra.Command {
	var command = &cobra.Command{
		Use:   "add",
		Short: "Add a command",
		Run:   runAddCommand,
	}

	command.Flags().StringP("file", "f", "", "File to add")
	command.Flags().StringP("name", "n", "", "Name of the command")

	return command
}

func runAddCommand(cmd *cobra.Command, args []string) {
	// validate flags
	file, _ := cmd.Flags().GetString("file")
	name, _ := cmd.Flags().GetString("name")

	if file == "" {
		cmd.Println("please provide a file to add")
		os.Exit(1)
		return
	}

	if name == "" {
		cmd.Println("please provide a name for the command")
		os.Exit(1)
		return
	}

	// validate file exists
	if _, err := os.Stat(file); os.IsNotExist(err) {
		cmd.Println("file does not exist")
		os.Exit(1)
		return
	}

	// Expand the tilde to the home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		cmd.Println("error getting home directory", err)
		os.Exit(1)
		return
	}
	configPathExpanded := homeDir + "/.ccli/"

	// Ensure the configPath directory exists
	err = os.MkdirAll(configPathExpanded, os.ModePerm)
	if err != nil {
		cmd.Println("error creating directory", err)
		os.Exit(1)
		return
	}

	commandPath := configPathExpanded + name
	// validate command does not exist in config
	if _, err := os.Stat(commandPath); !os.IsNotExist(err) {
		fmt.Println("command already exists")
		os.Exit(1)
		return
	}

	// copy file to configPath
	_, err = copyFile(file, commandPath)
	if err != nil {
		cmd.Println("error copying file", err)
		os.Exit(1)
		return
	}

	// make the file executable
	err = os.Chmod(commandPath, 0755)
	if err != nil {
		cmd.Println("error making file executable", err)
		os.Exit(1)
		return
	}
	// cmd.Println("added command", name)
	fmt.Println("added command", name)
}

func copyFile(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
