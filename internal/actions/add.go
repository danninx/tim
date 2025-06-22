package actions
import (
	"context"
	"fmt"

	"github.com/danninx/tim/internal/plate"
	"github.com/danninx/tim/internal/timfile"
	"github.com/urfave/cli/v3"
)

func Add(ctx context.Context, cmd *cli.Command) error {
	name := cmd.StringArg("name")
	plates, err := timfile.Read()
	if (err != nil) {
		return err
	}

	_, exists := plates[name]
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

	plates[name] = plate.Plate {
		Type: cmd.StringArg("type"),
		Path: cmd.StringArg("path"),
	}

	err = timfile.Write(plates)
	if (err != nil) {
		return err
	}

	fmt.Printf("%vadded \"%v\" to templates!%v\n", ANSI_GREEN, name, ANSI_RESET)

	return nil
}
