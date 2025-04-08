package main

import (
	"context"
	"fmt"
	"log"
	"os"

	action "github.com/Kenec/aliaz/cmd"
	"github.com/Kenec/aliaz/util"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "aliaz",
		Usage: "add, list, edit, and delete alias in the shell",
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

					// TODO: Check if alias already exist in shell.

					action.AddAlias(cmd.Args().First(), cmd.Args().Get(1))
					return nil
				},
			},
			{
				Name:    "list",
				Aliases: []string{"l", "-l"},
				Usage:   "aliaz list",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					action.ListAliases()
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

					// 2. TODO: Check if the alias exists

					// 3. Remove alias
					action.RemoveAlias(cmd.Args().First())
					// 4. Add new alias
					action.AddAlias(cmd.Args().First(), cmd.Args().Get(1))

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

					// 2. TODO" Check if alias exist

					// 3. Remove alias
					action.RemoveAlias(cmd.Args().First())
					return nil
				},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
