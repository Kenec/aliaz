package root

import (
	"context"
	"fmt"
	"os"

	"github.com/Kenec/aliaz/cmd/aliaz/actions"
	"github.com/Kenec/aliaz/cmd/aliaz/util"
	"github.com/urfave/cli/v3"
)

func Commands() *cli.Command {
	// Initialize the shell path
	shell := util.DetectShell()
	if shell == "" {
		fmt.Println("Unsupported shell. Please use bash, zsh, or fish.")
		os.Exit(1)
	}

	shellConfig, err := util.SetShell(shell)
	if err != nil {
		fmt.Println("Error setting shell:", err)
		os.Exit(1)
	}

	return &cli.Command{
		Name:  "aliaz",
		Version: "0.1.0",
		Usage: "A command line tool to manage shell aliases",
		Description: "Aliaz is a command line tool to manage shell aliases. It allows you to add, list, edit, and delete aliases in the shell.",
		ArgsUsage: "<alias> <command> eg: aliaz add gs \"git status\"",
		Action: func(context.Context, *cli.Command) error {
			fmt.Println("aliaz")
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    "add",
				Aliases: []string{"a", "-a"},
				Usage:   "aliaz add <alias> \"<command>\"",
				Action: func(ctx context.Context, cmd *cli.Command) error {

					util.ValidateArguments(cmd.Args().First(), cmd.Args().Get(1))

					// Check if the alias already exists
					aliasExists, err := util.ValidateAlias(cmd.Args().First(), shellConfig.ShellPath)
					if err != nil {
						fmt.Println("Error validating alias:", err)
						os.Exit(1)
					}
					if aliasExists {
						fmt.Printf("Alias '%s' already exists. Use 'aliaz edit' to modify it.\n", cmd.Args().First())
						os.Exit(1)
					}

					// Add alias
					action.AddAlias(cmd.Args().First(), cmd.Args().Get(1), shellConfig.ShellPath)
					return nil
				},
			},
			{
				Name:    "list",
				Aliases: []string{"l", "-l"},
				Usage:   "aliaz list",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					action.ListAliases(shellConfig.ShellPath)
					return nil
				},
			},
			{
				Name:    "edit",
				Aliases: []string{"e", "-e"},
				Usage:   "aliaz edit <alias> \"<command>\"",
				Action: func(ctx context.Context, cmd *cli.Command) error {

					// 1. Validate arguments
					util.ValidateArguments(cmd.Args().First(), cmd.Args().Get(1))

					// Check if the alias already exists
					aliasExists, err := util.ValidateAlias(cmd.Args().First(), shellConfig.ShellPath)
					if err != nil {
						fmt.Println("Error validating alias:", err)
						os.Exit(1)
					}

					if !aliasExists {
						fmt.Printf("Alias %s does not exists. Use 'aliaz add' to add it.\n", cmd.Args().First())
						os.Exit(1)
					}

					// 3. Remove alias
					action.RemoveAlias(cmd.Args().First(), shellConfig.ShellPath)
					// 4. Add new alias
					action.AddAlias(cmd.Args().First(), cmd.Args().Get(1), shellConfig.ShellPath)

					return nil
				},
			},
			{
				Name:    "remove",
				Aliases: []string{"rm", "-rm"},
				Usage:   "aliaz remove <alias>",
				Action: func(ctx context.Context, cmd *cli.Command) error {

					// 1. Validate arguments
					util.ValidateArguments(cmd.Args().First())

					// Check if the alias exists
					aliasExists, err := util.ValidateAlias(cmd.Args().First(), shellConfig.ShellPath)
					if err != nil {
						fmt.Println("Error validating alias:", err)
						os.Exit(1)
					}

					if !aliasExists {
						fmt.Printf("Alias %s does not exists. Use 'aliaz add' to add it.\n", cmd.Args().First())
						os.Exit(1)
					}

					// 3. Remove alias
					action.RemoveAlias(cmd.Args().First(), shellConfig.ShellPath)
					return nil
				},
			},
		},
	}
}
