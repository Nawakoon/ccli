package delete

import (
	"bytes"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddCommand(t *testing.T) {

	const mainPath = "../../"
	var configPath string = ""

	beforeAll := func() {
		homeDir, err := os.UserHomeDir()
		assert.NoError(t, err)
		configPath = homeDir + "/.ccli/"
	}

	beforeAll()

	// ccli delete
	t.Run("should require command name", func(t *testing.T) {
		command := exec.Command("go", "run", "cmd/main.go", "delete")
		var stdout, stderr bytes.Buffer
		command.Stdout = &stdout
		command.Stderr = &stderr
		command.Dir = mainPath

		err := command.Run()

		assert.Error(t, err)
		assert.Equal(t, 1, command.ProcessState.ExitCode())
		assert.Contains(t, stderr.String(), "please provide a command to delete")
	})

	// ccli delete non-existent-command-ec37eb9b7164509ceeaea20e5a2f974e2fe9b4ce
	t.Run("should return error when command does not exist", func(t *testing.T) {
		notExistCommand := "non-existent-command-ec37eb9b7164509ceeaea20e5a2f974e2fe9b4ce"
		command := exec.Command("go", "run", "cmd/main.go", "delete", notExistCommand)
		var stdout, stderr bytes.Buffer
		command.Stdout = &stdout
		command.Stderr = &stderr
		command.Dir = mainPath

		err := command.Run()

		assert.Error(t, err)
		assert.Equal(t, 1, command.ProcessState.ExitCode())
		assert.Contains(t, stderr.String(), "command does not exist")
	})

	// ccli delete test-command-f987c1bb6cbb4ff2af39730843c67bd1ec0cf6b9
	t.Run("should delete command", func(t *testing.T) {
		addedCommand := "test-command-f987c1bb6cbb4ff2af39730843c67bd1ec0cf6b9"
		// add command to config for deletion
		commandPath := configPath + addedCommand
		_, err := os.Create(commandPath)
		defer os.Remove(commandPath)

		command := exec.Command("go", "run", "cmd/main.go", "delete", addedCommand)
		var stdout, stderr bytes.Buffer
		command.Stdout = &stdout
		command.Stderr = &stderr
		command.Dir = mainPath

		err = command.Run()

		assert.NoError(t, err)
		assert.Equal(t, 0, command.ProcessState.ExitCode())
		assert.Contains(t, stdout.String(), "command deleted successfully")
	})
}
