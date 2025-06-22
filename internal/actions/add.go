package actions

import (
	"context"
	"fmt"

	"github.com/danninx/tim/internal/conf"
	"github.com/danninx/tim/internal/plate"
	"github.com/urfave/cli/v3"
)

func Add(ctx context.Context, cmd *cli.Command) error {
	name := cmd.StringArg("name")
	config, err := conf.Load()
	if (err != nil) {
		return err
	}

	_, exists := config.Plates[name]
	if exists {
		msg := fmt.Sprintf("%vsource \"%v\" already exists, would you like to replace it? (y/N)%v", ANSI_YELLOW, name, ANSI_RESET)
		confirm, err := ConfirmAction(msg)
		if (err != nil) {
			return err
		}
		if !confirm {
			fmt.Printf("skipping...")	
			return nil
		}
	}

	config.Plates[name] = plate.Plate {
		Type: cmd.StringArg("type"),
		Path: cmd.StringArg("path"),
	}

	err = conf.Save(config)
	if (err != nil) {
		return err
	}

	fmt.Printf("%vadded \"%v\" to templates!%v\n", ANSI_GREEN, name, ANSI_RESET)

	return nil
}
