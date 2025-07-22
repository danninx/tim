package actions

import (
	"context"
	"fmt"

	"github.com/danninx/tim/internal/conf"
	"github.com/danninx/tim/internal/plate"
	"github.com/urfave/cli/v3"
)

func Remove(ctx context.Context, cmd *cli.Command) error {
	config, err := conf.Load()
	if err != nil {
		return err
	}

	name := cmd.StringArg("name")
	template, exists := config.Plates[name]

	if !exists {
		return &NO_PLATE_EXISTS{
			Name: name,
		}
	}

	msg := fmt.Sprintf("%vare you sure you want to delete source \"%v\"? (y/N)%v", ANSI_YELLOW, name, ANSI_RESET)
	confirm, err := ConfirmAction(msg)

	if err != nil {
		return err
	} else if !confirm {
		fmt.Printf("skipping...")
		return nil
	}

	template, err = plate.Load(name, template)
	template.Delete()
	delete(config.Plates, name)

	conf.Save(config)

	return nil
}
