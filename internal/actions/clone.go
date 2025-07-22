package actions

import (
	"context"

	"github.com/danninx/tim/internal/plate"
	"github.com/urfave/cli/v3"
)

func Clone(ctx context.Context, cmd *cli.Command) error {
	name := cmd.StringArg("name")
	dest := cmd.StringArg("dest")

	template, err := GetPlate(name)
	if err != nil {
		return err
	}

	source, err := plate.Load(name, template)
	if err != nil {
		return err
	}
	return source.Copy(dest)
}
