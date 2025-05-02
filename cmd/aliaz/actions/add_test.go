package action_test

import (
	"testing"

	"github.com/Kenec/aliaz/cmd/aliaz/actions"
	"github.com/Kenec/aliaz/cmd/aliaz/util"
)

func TestAddAlias(t *testing.T) {
	// set up the test environment

	shellConfig := util.ShellConfig {
		Name: "bash",
		ShellPath: "../../../assets/.bashrc",
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
			// This is a placeholder for the actual test logic
			t.Logf("Alias %s added for command %s", test.alias, test.command)
		})
	}

}
