package actions

import (
	"context"

	"github.com/danninx/tim/internal/conf"
	"github.com/urfave/cli/v3"
)

func Migrate(ctx context.Context, cmd *cli.Command) error {
	t := cmd.StringArg("type")
	
	config, err := conf.Load()
	if err != nil {
		return err
	}			

	if err = conf.SetConfFileType(t); err != nil {
		return err
	}
	
	return conf.Save(config)
}
