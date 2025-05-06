package action_test

import (
	"os"
	"testing"

	"github.com/Kenec/aliaz/cmd/aliaz/actions"
	"github.com/Kenec/aliaz/cmd/aliaz/util"
)

func TestRemoveAlias(t *testing.T) {
	// set up the test environment
	testShellPath := "../../../assets/.bashrc"
	err := CreateTestShellFile(testShellPath)
	if err != nil {
		t.Fatalf("Failed to create test shell file: %v", err)
	}

	// Add alias
	action.AddAlias("mk", "mkdir", testShellPath)

	aliasToRemove := "mk"

	// Test case for removing an alias
	t.Run(aliasToRemove, func(t *testing.T) {
		action.RemoveAlias(aliasToRemove, testShellPath)
		// Here you would check if the alias was removed correctly
		isValid, err := util.ValidateAlias(aliasToRemove, testShellPath)
		if err != nil {
			t.Fatalf("Error validating alias: %v", err)
		}
		if isValid {
			t.Errorf("Alias %s was not removed correctly", aliasToRemove)
		}
	})
	// Clean up
	defer os.Remove(testShellPath)
}
