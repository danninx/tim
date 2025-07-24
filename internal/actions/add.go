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

func Add(ctx context.Context, cmd *cli.Command) error {
	sys := system.GetSystem()

	plateType := cmd.StringArg("type")
	plateName := cmd.StringArg("name")
	plateOrigin := cmd.StringArg("origin")

	if plateType == "" || plateName == "" || plateOrigin == "" {
		cli.ShowSubcommandHelp(cmd)
		os.Exit(1)
	}

	// check if plate already exists
	config, err := conf.Load(sys)
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

	newPlate, err := plate.NewPlate(plateType, plateName, plateOrigin, sys)
	if err != nil {
		return err
	}

	config.Plates[plateName] = plate.Unload(newPlate)
	err = conf.Save(config, sys)
	if err != nil {
		return err
	}

	fmt.Printf("%vadded \"%v\" to templates!%v\n", ANSI_GREEN, plateName, ANSI_RESET)

	return nil
}
