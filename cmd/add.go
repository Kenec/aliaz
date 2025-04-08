package cmd

import (
	"fmt"
	"os"
)

func AddAlias(alias, command string) {
	configFile, err := setShell()

	if err != nil {
		fmt.Println("error in adding alias ", err)
		os.Exit(1)
	}

	entry := fmt.Sprintf("alias %s='%s'\n", alias, command)
	file, err := os.OpenFile(expandPath(configFile), os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	if _, err := file.WriteString(entry); err != nil {
		fmt.Println("Error writing to file:", err)
		os.Exit(1)
	}
	fmt.Println("Alias " + alias + " added successfully. Restart your shell or run 'source " + configFile + "' to apply changes.")
}
