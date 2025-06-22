package actions

import (
	"context"
	"fmt"
	"io/fs"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func PrintDir(ctx context.Context, cmd *cli.Command) error {
	return fs.WalkDir(os.DirFS(cmd.StringArg("dir")), ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(path)
		return nil
	})

	return nil
}
