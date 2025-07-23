package actions

import (
	"context"
	"os"

	"github.com/danninx/tim/internal/plate"
	"github.com/danninx/tim/internal/system"
	"github.com/urfave/cli/v3"
)

func Clone(ctx context.Context, cmd *cli.Command) error {
	name := cmd.StringArg("name")
	dest := cmd.StringArg("dest")

	if name == "" {
		cli.ShowSubcommandHelp(cmd)
		os.Exit(1)
	}

	if dest == "" {
		dest = "."
	}

	template, err := GetPlate(name)
	if err != nil {
		return err
	}

	source, err := plate.Load(name, template, system.GetSystem())
	if err != nil {
		return err
	}
	return source.Copy(dest)
}
