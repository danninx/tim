package actions

import (
	"context"
	"fmt"
	"os"

	"github.com/danninx/tim/internal/conf"
	"github.com/danninx/tim/internal/plate"
	"github.com/danninx/tim/internal/system"
	"github.com/urfave/cli/v3"
)

func Remove(ctx context.Context, cmd *cli.Command) error {
	sys := system.GetSystem()

	name := cmd.StringArg("name")

	if name == "" {
		cli.ShowSubcommandHelp(cmd)
		os.Exit(1)
	}

	config, err := conf.Load(sys)
	if err != nil {
		return err
	}

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

	source, err := plate.Load(name, template, sys)
	if err != nil {
		return err
	}
	source.Delete()
	delete(config.Plates, name)

	conf.Save(config, sys)

	return nil
}
