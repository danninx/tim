package actions

import (
	"context"
	"fmt"
	"slices"

	"github.com/danninx/tim/internal/plate"
	"github.com/danninx/tim/internal/timfile"
	"github.com/urfave/cli/v3"
)


func List(ctx context.Context, cmd *cli.Command) error {
	fmt.Println("tim sources:")
	sources, err := timfile.Read()

	if (err != nil) {
		return err
	}

	var keys[] string
	for k := range sources {
		keys = append(keys, k)
	}
	slices.Sort(keys)

	for _, k := range keys {
		printPlateLine(k, sources[k])
	}

	return nil
}

func printPlateLine(name string, p plate.Plate) {
	n := fmt.Sprintf("%v%-15s%v", ANSI_GREEN, name, ANSI_RESET)
	i := fmt.Sprintf("%-s\n", p.Path)
	var t string
	switch p.Type {
	case "git":
		t = fmt.Sprintf("%s%-8s%s", ANSI_YELLOW, p.Type, ANSI_RESET)
	case "dir":
		t = fmt.Sprintf("%s%-8s%s", ANSI_BLUE, p.Type, ANSI_RESET)
	case "file":
		t = fmt.Sprintf("%s%-8s%s", ANSI_MAGENTA, p.Type, ANSI_RESET)
	}
	fmt.Printf("\t- %-s%-s%-s", n, t, i)
}

