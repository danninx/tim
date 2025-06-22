package actions

import (
	"context"
	"fmt"

	"github.com/danninx/tim/internal/files"
	"github.com/urfave/cli/v3"
)

func Clone(ctx context.Context, cmd *cli.Command) error {
	name := cmd.StringArg("name")

	plate, err := GetPlate(name)
	if (err != nil) {
		return err
	}

	dest := cmd.StringArg("dest")

	if (plate.Type == "git" || plate.Type == "dir") {
		tmp, err := files.TempCopy(plate.Path)
		if err != nil {
			return err
		}
		fmt.Println("copying cleaned source...")
		err = files.CopyDir(tmp, dest)	
		if err != nil {
			return err
		}
		fmt.Println("cleaning temporary directory...")
		err = files.CleanTmp()
		if err != nil {
			return err
		}
	} else if plate.Type == "file" {
		src, err := CheckPathExists(plate.Path)
		if (err != nil) {
			return err
		}
		err = files.CopyFile(src, dest)			
		if err != nil {
			return err
		}
	}  else {
		return &INVALID_PLATE_TYPE{
			Type: plate.Type, 
		}
	}

	return nil
}
