package actions

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli/v3"
)

func Show(ctx context.Context, cmd *cli.Command) error {
	name := cmd.StringArg("name")
	if name == "" {
		cli.ShowSubcommandHelp(cmd)
		os.Exit(1)
	}

	template, err := GetPlate(name)
	if err != nil {
		return err
	}

	// PREPARE STRINGS
	typeString, err := colorPlateType(template.Type)
	if err != nil {
		return err
	}
	plateTitle := fmt.Sprintf("Showing info for plate %v\"%v\"%v", ANSI_GREEN, name, ANSI_RESET)
	plateType := fmt.Sprintf("type:\t%v", typeString)
	plateValue := fmt.Sprintf("value: \t%v", template.Path)

	lineLength := max(25+len(name), len(plateType), len(plateValue))
	line := strings.Repeat("-", lineLength)

	fmt.Printf(
		"%v\n%v\n%v\n%v\n%v\n",
		line,
		plateTitle,
		line,
		plateType,
		plateValue,
	)

	return nil
}

func colorPlateType(t string) (string, error) {
	if t == "git" {
		return fmt.Sprintf("%v%v%v", ANSI_YELLOW, t, ANSI_RESET), nil
	} else if t == "dir" {
		return fmt.Sprintf("%v%v%v", ANSI_BLUE, t, ANSI_RESET), nil
	} else if t == "file" {
		return fmt.Sprintf("%v%v%v", ANSI_MAGENTA, t, ANSI_RESET), nil
	} else {
		return "", &INVALID_PLATE_TYPE{Type: t}
	}
}
