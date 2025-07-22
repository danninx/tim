package plate

import (
	"fmt"
	"path/filepath"

	"github.com/danninx/tim/internal/system"
)

type dirPlate struct {
	name   	string
	origin 	string
	path   	string
	sys 	system.System
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
	err := plate.sys.RemoveAll(plate.path)
	if err != nil {
		return err
	}
	return plate.sys.CopyDir(plate.origin, plate.path)
}

func (plate *dirPlate) Copy(destination string) error {
	return plate.sys.CopyDir(plate.path, destination)
}

func (plate *dirPlate) Delete() error {
	return plate.sys.RemoveAll(plate.path)
}

func newDirPlate(name string, origin string, sys system.System) (Plate, error) {
	fmt.Println("saving a copy of the directory to tim directory...")
	timDir, err := sys.TimDirectory()
	if err != nil {
		return nil, err
	}

	clonePath := filepath.Join(timDir, "dir", name)
	err = sys.CopyDir(origin, clonePath)
	if err != nil {
		return nil, err
	}

	fmt.Printf("copy saved to %v\n", clonePath)
	return &dirPlate{
		name:   name,
		origin: origin,
		path:   clonePath,
		sys:	sys,
	}, nil
}
