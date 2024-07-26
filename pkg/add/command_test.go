package add

import (
	"bytes"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

// integration test
func TestAddCommand(t *testing.T) {

	const mainPath = "../../"
	const testFile = "/tmp/hello-world-test.sh"

	var configPath string = ""

	beforeAll := func() {
		// create tmp file for test
		testData := []byte("echo 'hello'")
		err := os.WriteFile(testFile, testData, 0644)
		assert.NoError(t, err)
		homeDir, err := os.UserHomeDir()
		assert.NoError(t, err)
		configPath = homeDir + "/.ccli/"
	}

	afterAll := func() {
		// delete the tmp file for test
		err := os.Remove(testFile)
		assert.NoError(t, err)
	}

	beforeAll()
	defer afterAll()

	// ccli add --name 01-hello-world-test
	t.Run("should require file flag", func(t *testing.T) {
		command := exec.Command("go", "run", "cmd/main.go", "add", "--name", "01-hello-world-test")
		var stdout, stderr bytes.Buffer
		command.Stdout = &stdout
		command.Stderr = &stderr
		command.Dir = mainPath

		err := command.Run()

		assert.Error(t, err)
		assert.Equal(t, 1, command.ProcessState.ExitCode())
		assert.Contains(t, stderr.String(), "please provide a file to add")
	})

	// ccli add --file hello-world-test.py
	t.Run("should require name flag", func(t *testing.T) {
		command := exec.Command("go", "run", "cmd/main.go", "add", "--file", "hello-world-test.py")
		var stdout, stderr bytes.Buffer
		command.Stdout = &stdout
		command.Stderr = &stderr
		command.Dir = mainPath

		err := command.Run()

		assert.Error(t, err)
		assert.Equal(t, 1, command.ProcessState.ExitCode())
		assert.Contains(t, stderr.String(), "please provide a name for the command")
	})

	// ccli add --name 01-hello-world-test --file not-exist--file.rs
	t.Run("should return error if file does not exist", func(t *testing.T) {
		command := exec.Command("go", "run", "cmd/main.go", "add", "--name", "01-hello-world-test", "--file", "not-exist-file.rs")
		var stdout, stderr bytes.Buffer
		command.Stdout = &stdout
		command.Stderr = &stderr
		command.Dir = mainPath

		err := command.Run()

		assert.Error(t, err)
		assert.Equal(t, 1, command.ProcessState.ExitCode())
		assert.Contains(t, stderr.String(), "file does not exist")
	})

	// ccli add --name 01-hello-world-test --file hello-world-test.sh
	t.Run("should add command to config correctly", func(t *testing.T) {
		defer func() {
			err := os.Remove(configPath + "01-hello-world-test")
			assert.NoError(t, err)
		}()
		commandName := "01-hello-world-test"
		command := exec.Command("go", "run", "cmd/main.go", "add", "--name", commandName, "--file", testFile)
		var stdout, stderr bytes.Buffer
		command.Stdout = &stdout
		command.Stderr = &stderr
		command.Dir = mainPath

		err := command.Run()

		assert.NoError(t, err)
		assert.Equal(t, 0, command.ProcessState.ExitCode())
		assert.Contains(t, stdout.String(), "added command 01-hello-world-test")
		configCommand := configPath + commandName
		_, err = os.Stat(configCommand)
		assert.False(t, os.IsNotExist(err))
		checkCommand := exec.Command("which", commandName)
		err = checkCommand.Run()
		assert.NoError(t, err)
		assert.Equal(t, 0, checkCommand.ProcessState.ExitCode())
	})

	t.Run("should return error if command already exists", func(t *testing.T) {
		defer func() {
			err := os.Remove(configPath + "01-hello-world-test")
			assert.NoError(t, err)
		}()
		command1 := exec.Command("go", "run", "cmd/main.go", "add", "--name", "01-hello-world-test", "--file", testFile)
		command2 := exec.Command("go", "run", "cmd/main.go", "add", "--name", "01-hello-world-test", "--file", testFile)

		command1.Dir = mainPath
		command2.Dir = mainPath

		var stdout, stderr bytes.Buffer
		command2.Stdout = &stdout
		command2.Stderr = &stderr

		command1.Run()
		err := command2.Run()

		assert.Error(t, err)
		assert.Equal(t, 1, command2.ProcessState.ExitCode())
		assert.Contains(t, stdout.String(), "command already exists")
	})
}
