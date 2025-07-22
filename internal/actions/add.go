package actions

import (
	"context"
	"fmt"

	"github.com/danninx/tim/internal/conf"
	"github.com/danninx/tim/internal/plate"
	"github.com/urfave/cli/v3"
)

func Add(ctx context.Context, cmd *cli.Command) error {
	plateType := cmd.StringArg("type")
	plateName := cmd.StringArg("name")
	plateOrigin := cmd.StringArg("origin")

	// check if plate already exists
	config, err := conf.Load()
	if err != nil {
		return err
	}

	_, exists := config.Plates[plateName]
	if exists {
		msg := fmt.Sprintf("%vsource \"%v\" already exists, would you like to replace it? (y/N)%v", ANSI_YELLOW, plateName, ANSI_RESET)
		confirm, err := ConfirmAction(msg)
		if err != nil {
			return err
		}
		if !confirm {
			fmt.Printf("skipping...")
			return nil
		}
	}

	newPlate, err := plate.NewPlate(plateType, plateName, plateOrigin)
	if err != nil {
		return err
	}

	config.Plates[plateName] = plate.Unload(newPlate)
	err = conf.Save(config)
	if err != nil {
		return err
	}

	fmt.Printf("%vadded \"%v\" to templates!%v\n", ANSI_GREEN, plateName, ANSI_RESET)

	return nil
}
