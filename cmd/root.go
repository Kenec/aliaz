package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

const bashrc = "~/.bashrc"
const zshrc = "~/.zshrc"
const fishConfig = "~/.config/fish/config.fish"

func detectShell() string {
	shell := os.Getenv("SHELL")
	if strings.Contains(shell, "bash") {
		return "bash"
	} else if strings.Contains(shell, "zsh") {
		return "zsh"
	} else if strings.Contains(shell, "fish") {
		return "fish"
	}
	return ""
}

func setShell() (string, error) {
	shell := detectShell()

	var shellConfig string

	switch shell {
	case "bash":
		shellConfig = bashrc
	case "zsh":
		shellConfig = zshrc
	case "fish":
		shellConfig = fishConfig
	default:
		return "", errors.New("unsupported shell detected")
	}

	return shellConfig, nil
}

func expandPath(path string) string {
	if strings.HasPrefix(path, "~") {
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error finding user home directory:", err)
			os.Exit(1)
		}
		return strings.Replace(path, "~", home, 1)
	}
	return path
}
