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
				Usage:    "create files using a plate",
				Action:   actions.Clone,
				Category: "management",
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
				Usage:    "add a plate into your templates",
				Action:   actions.Add,
				Category: "management",
			},
			{
				Name:    "remove",
				Aliases: []string{"rm"},
				Arguments: []cli.Argument{
					&cli.StringArg{
						Name: "name",
					},
				},
				Usage:    "remove a plate from your templates",
				Action:   actions.Remove,
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
				Usage:    "show information about a specific source",
				Action:   actions.Show,
				Category: "info",
			},

			// CONFIG
			{
				Name: "migrate",
				Arguments: []cli.Argument{
					&cli.StringArg{
						Name: "type",
					},
				},
				Usage:    "migrate configuration file to a specific format",
				Action:   actions.Migrate,
				Category: "config",
			},

			// DEBUG
			{
				Name: "print",
				Arguments: []cli.Argument{
					&cli.StringArg{
						Name: "dir",
					},
				},
				Action: actions.PrintDir,
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
