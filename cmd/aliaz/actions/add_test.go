package action_test

import (
	"os"
	"testing"

	"github.com/Kenec/aliaz/cmd/aliaz/actions"
	"github.com/Kenec/aliaz/cmd/aliaz/util"
)

// Test setup
func CreateTestShellFile(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()

	return nil
}

func TestAddAlias(t *testing.T) {
	// set up the test environment
	testShellPath := "../../../assets/.bashrc"
	err := CreateTestShellFile(testShellPath)
	if err != nil {
		t.Fatalf("Failed to create test shell file: %v", err)
	}

	// Set up the shell configuration
	shellConfig := util.ShellConfig{
		Name:      "bash",
		ShellPath: testShellPath,
	}
	// Test cases for adding aliases
	tests := []struct {
		alias   string
		command string
	}{
		{"ll", "ls -la"},
		{"gs", "git status"},
	}

	for _, test := range tests {
		t.Run(test.alias, func(t *testing.T) {
			action.AddAlias(test.alias, test.command, shellConfig.ShellPath)
			// Here you would check if the alias was added correctly
			isValid, err := util.ValidateAlias(test.alias, shellConfig.ShellPath)
			if err != nil {
				t.Fatalf("Error validating alias: %v", err)
			}
			if !isValid {
				t.Errorf("Alias %s was not added correctly", test.alias)
			}
			// This is a placeholder for the actual test logic
			t.Logf("Alias %s added for command %s", test.alias, test.command)
		})
	}

	defer os.Remove(testShellPath) // Clean up after test
}
