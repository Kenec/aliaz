package cmd

import (
	"fmt"
	"os"
	"strings"
)

func ListAliases() {
	configFile, err := setShell()

	if err != nil {
		fmt.Println("error in listing alias ", err)
		os.Exit(1)
	}

	data, err := os.ReadFile(expandPath(configFile))
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
