package util

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

const bashrc = "~/.bashrc"
const zshrc = "~/.zshrc"
const fishConfig = "~/.config/fish/config.fish"

type ShellConfig struct {
	Name      string
	ShellPath string
}

func DetectShell() string {
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

func SetShell(shell string) (*ShellConfig, error) {

	var shellConfig *ShellConfig

	switch shell {
	case "bash":
		shellConfig = &ShellConfig{
			Name:      "bash",
			ShellPath: ExpandPath(bashrc),
		}
	case "zsh":
		shellConfig = &ShellConfig{
			Name:      "zsh",
			ShellPath: ExpandPath(zshrc),
		}
	case "fish":
		shellConfig = &ShellConfig{
			Name:      "fish",
			ShellPath: ExpandPath(fishConfig),
		}
	default:
		return nil, errors.New("unsupported shell detected")
	}

	return shellConfig, nil
}

func ExpandPath(path string) string {
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
