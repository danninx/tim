package actions

import (
	"context"
	"fmt"

	"github.com/danninx/tim/internal/timfile"
	"github.com/urfave/cli/v3"
)

func Remove(ctx context.Context, cmd *cli.Command) error {
	sources, err := timfile.Read()
	if (err != nil) {
		return err
	}

	name := cmd.StringArg("name")
	_, exists := sources[name]

	if !exists {
		return &NO_PLATE_EXISTS{
			Name: name,
		}
	}

	msg := fmt.Sprintf("%vare you sure you want to delete source \"%v\"? (y/N)%v", ANSI_YELLOW, name, ANSI_RESET)
	confirm, err := ConfirmAction(msg)

	if (err != nil) {
		return err
	} else if !confirm {
		fmt.Printf("skipping...")	
		return nil
	}

	delete(sources, name)

	timfile.Write(sources)

	return nil
}
