package cmd

import (
	"fmt"
	"os"
	"strings"
)

func RemoveAlias(alias string) {
	configFile, err := setShell()

	if err != nil {
		fmt.Println("error in adding alias ", err)
		os.Exit(1)
	}

	data, err := os.ReadFile(expandPath(configFile))
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

	os.WriteFile(expandPath(configFile), []byte(strings.Join(newLines, "\n")), 0644)
	fmt.Printf("Alias %s removed successfully.\n", alias)

}
