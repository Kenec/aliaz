package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ValidateArguments(arguments ...string) {
	for i, str := range arguments {
		if str == "" {
			fmt.Printf("argument %d is empty or unset\n", i+1)
			os.Exit(1)
		}
	}
}

func ValidateAlias(alias, shellPath string) (bool, error) {
	file, err := os.Open(shellPath)

	if err != nil {
		return false, fmt.Errorf("error opening file: %w", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	aliasPrefix := fmt.Sprintf("alias %s=", alias)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Skip commented lines
		if strings.HasPrefix(line, "#") {
			continue
		}

		// Check if the line starts with the alias prefix
		if strings.HasPrefix(line, aliasPrefix) {
			return true, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return false, fmt.Errorf("error reading file: %w", err)
	}

	return false, nil
}
