package action

import (
	"fmt"
	"os"
	"strings"
)

func ListAliases(shellPath string) {
	data, err := os.ReadFile(shellPath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	lines := strings.Split(string(data), "\n")
	fmt.Println("Aliases:")
	for _, line := range lines {
		if strings.HasPrefix(line, "alias ") {
			fmt.Println(line)
		}
	}
}
