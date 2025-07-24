package actions

import (
	"context"
	"fmt"
	"slices"

	"github.com/danninx/tim/internal/conf"
	"github.com/danninx/tim/internal/plate"
	"github.com/danninx/tim/internal/system"
	"github.com/urfave/cli/v3"
)

func List(ctx context.Context, cmd *cli.Command) error {
	sys := system.GetSystem()

	config, err := conf.Load(sys)

	if err != nil {
		return err
	}

	var keys []string
	for k := range config.Plates {
		keys = append(keys, k)
	}
	slices.Sort(keys)

	fmt.Println("tim sources:")
	for _, k := range keys {
		printPlateLine(k, config.Plates[k])
	}

	return nil
}

func printPlateLine(name string, p plate.UnloadedPlate) {
	n := fmt.Sprintf("%v%-15s%v", ANSI_GREEN, name, ANSI_RESET)
	i := fmt.Sprintf("%-s\n", p.Origin)
	var t string
	switch p.Type {
	case "git":
		t = fmt.Sprintf("%s%-8s%s", ANSI_YELLOW, p.Type, ANSI_RESET)
	case "dir":
		t = fmt.Sprintf("%s%-8s%s", ANSI_BLUE, p.Type, ANSI_RESET)
	case "file":
		t = fmt.Sprintf("%s%-8s%s", ANSI_MAGENTA, p.Type, ANSI_RESET)
	}
	fmt.Printf("- %-s%-s%-s", n, t, i)
}
