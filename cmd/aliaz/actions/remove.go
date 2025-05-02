package action

import (
	"fmt"
	"os"
	"strings"
)

func RemoveAlias(alias, shellPath string) {
	data, err := os.ReadFile(shellPath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	lines := strings.Split(string(data), "\n")
	var newLines []string

	for _, line := range lines {
		if !strings.HasPrefix(line, "alias "+alias+"=") {
			newLines = append(newLines, line)
		}
	}

	os.WriteFile(shellPath, []byte(strings.Join(newLines, "\n")), 0644)
	fmt.Printf("Alias %s removed successfully.\n", alias)
}
