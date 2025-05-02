package action

import (
	"fmt"
	"os"
)

func AddAlias(alias, command, shellPath string) {
	entry := fmt.Sprintf("alias %s='%s'\n", alias, command)
	file, err := os.OpenFile(shellPath, os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	if _, err := file.WriteString(entry); err != nil {
		fmt.Println("Error writing to file:", err)
		os.Exit(1)
	}
	fmt.Println("Alias " + alias + " added successfully. Restart your shell or run 'source " + shellPath + "' to apply changes.")
}
