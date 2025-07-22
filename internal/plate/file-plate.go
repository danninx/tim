package plate

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/danninx/tim/internal/system"
)

type filePlate struct {
	name   	string
	origin 	string
	path   	string
	sys 	system.System
}

func (plate *filePlate) Name() string {
	return plate.name
}

func (plate *filePlate) Origin() string {
	return plate.origin
}

func (plate *filePlate) Path() string {
	return plate.path
}

func (plate *filePlate) Type() string {
	return "file"
}

func (plate *filePlate) Sync() error {
	os.RemoveAll(plate.path)
	return plate.sys.CopyFile(plate.origin, plate.path)
}

func (plate *filePlate) Copy(destination string) error {
	return plate.sys.CopyDir(plate.path, destination)
}

func (plate *filePlate) Delete() error {
	return plate.sys.RemoveAll(plate.path)
}

func newFilePlate(name string, origin string, sys system.System) (Plate, error) {
	fmt.Println("saving a copy of the file to tim directory...")
	timDir, err := sys.TimDirectory()
	if err != nil {
		return nil, err
	}

	clonePath := filepath.Join(timDir, "files", name)
	err = sys.CopyFile(origin, clonePath)
	if err != nil {
		return nil, err
	}

	fmt.Printf("copy saved to %v\n", clonePath)
	return &filePlate{
		name:   name,
		origin: origin,
		path:   clonePath,
		sys:	sys,
	}, nil
}
