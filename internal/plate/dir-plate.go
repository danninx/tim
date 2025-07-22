package plate

import (
	"fmt"
	"os"
	"path/filepath"
)

type dirPlate struct {
	name   string
	origin string
	path   string
}

func (plate *dirPlate) Name() string {
	return plate.name
}

func (plate *dirPlate) Origin() string {
	return plate.origin
}

func (plate *dirPlate) Path() string {
	return plate.path
}

func (plate *dirPlate) Type() string {
	return "dir"
}

func (plate *dirPlate) Sync() error {
	err := os.RemoveAll(plate.path)
	if err != nil {
		return err
	}
	return CopyDir(plate.origin, plate.path)
}

func (plate *dirPlate) Copy(destination string) error {
	return CopyDir(plate.path, destination)
}

func (plate *dirPlate) Delete() error {
	return os.RemoveAll(plate.path)
}

func newDirPlate(name string, origin string) (Plate, error) {
	fmt.Println("saving a copy of the directory to tim directory...")
	timDir, err := TimDirectory()
	if err != nil {
		return nil, err
	}

	clonePath := filepath.Join(timDir, "dir", name)
	err = CopyDir(origin, clonePath)
	if err != nil {
		return nil, err
	}

	fmt.Printf("copy saved to %v\n", clonePath)
	return &dirPlate{
		name:   name,
		origin: origin,
		path:   clonePath,
	}, nil
}
