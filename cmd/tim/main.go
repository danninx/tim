package main

import (
	"context"
	"log"
	"os"

	"github.com/danninx/tim/internal/actions"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "tim",
		Usage: "manage your template files locally",
		Version: "0.3.0",
		Commands: []*cli.Command{
			// TEMPLATE
			{
				Name:    "plate",
				Aliases: []string{"clone"},
				Arguments: []cli.Argument{
					&cli.StringArg{
						Name: "name",
					},
					&cli.StringArg{
						Name: "dest",
					},
				},
				ArgsUsage: "<name> <dest>",
				Usage:     "create files using a plate",
				Action:    actions.Clone,
				Category:  "management",
			},

			// MANAGE
			{
				Name:    "add",
				Aliases: []string{"link"},
				Arguments: []cli.Argument{
					&cli.StringArg{
						Name: "type",
					},
					&cli.StringArg{
						Name: "name",
					},
					&cli.StringArg{
						Name: "origin",
					},
				},
				ArgsUsage: "<type> <name> <origin>",
				Usage:     "add a plate into your templates",
				Action:    actions.Add,
				Category:  "management",
			},
			{
				Name:    "remove",
				Aliases: []string{"rm"},
				Arguments: []cli.Argument{
					&cli.StringArg{
						Name: "name",
					},
				},
				ArgsUsage: "<name>",
				Usage:     "remove a plate from your templates",
				Action:    actions.Remove,
				Category:  "management",
			},
			{
				Name: "rename",
				Aliases: []string{"re"},
				Arguments: []cli.Argument{
					&cli.StringArg{
						Name: "name",
					},
					&cli.StringArg{
						Name: "newName",
					},
				},
				ArgsUsage: "<old-name> <new-name>",
				Usage: "rename a plate from your templates",
				Action: actions.Rename,
				Category: "management",
			},

			// INFO
			{
				Name:     "list",
				Aliases:  []string{"ls"},
				Usage:    "list current plates and basic info",
				Action:   actions.List,
				Category: "info",
			},
			{
				Name:    "show",
				Aliases: []string{"get"},
				Arguments: []cli.Argument{
					&cli.StringArg{
						Name: "name",
					},
				},
				ArgsUsage: "<name>",
				Usage:     "show information about a specific source",
				Action:    actions.Show,
				Category:  "info",
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
