package action_test

import (
	"bytes"
	"io"
	"os"
	"testing"

	action "github.com/Kenec/aliaz/cmd/aliaz/actions"
)

func TestListAliases(t *testing.T) {
	zshrcPath := "../../../assets/.zshrc"

	// Save original stdout
	originalStdout := os.Stdout

	// Create a pipe to capture stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("Failed to create pipe: %v", err)
	}

	os.Stdout = w

	action.ListAliases(zshrcPath)

	w.Close()
	os.Stdout = originalStdout

	// Read captured output
	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	// Assert output
	expected := "Aliases:\nalias ll='ls -la'\nalias gs='git status'\n"
	if output != expected {
		t.Errorf("Expected %q but got %q", expected, output)
	}
}
