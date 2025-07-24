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

func Rename(ctx context.Context, cmd *cli.Command) error {
	sys := system.GetSystem()

	name := cmd.StringArg("name")
	newName := cmd.StringArg("newName")

	if name == "" || newName == "" {
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

	_, exists = config.Plates[newName]
	if exists {
		fmt.Printf("a plate named \"%v\" already exists, skipping", newName)
		return nil
	}

	msg := fmt.Sprintf("%vare you sure you want to rename source \"%v\" to \"%v\"? (y/N)%v", ANSI_YELLOW, name, newName, ANSI_RESET)
	confirm, err := ConfirmAction(msg)

	if err != nil {
		return err
	} else if !confirm {
		fmt.Printf("skipping...")
		return nil
	}

	p, err := plate.Load(name, template, sys)
	if err != nil {
		return err
	}
	if err = p.Rename(newName); err != nil {
		return err
	}
	delete(config.Plates, name)
	config.Plates[newName] = plate.Unload(p)

	if err = conf.Save(config, sys); err != nil {
		return err
	}

	fmt.Printf("%vrenamed \"%v\" to \"%v\" succesfully%v\n", ANSI_GREEN, name, newName, ANSI_RESET)

	return nil
}
