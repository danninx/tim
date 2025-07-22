package plate

import (
	"fmt"
	"os"
	"path/filepath"
)

type filePlate struct {
	name   string
	origin string
	path   string
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
	return CopyFile(plate.origin, plate.path)
}

func (plate *filePlate) Copy(destination string) error {
	return CopyDir(plate.path, destination)
}

func (plate *filePlate) Delete() error {
	return os.RemoveAll(plate.path)
}

func newFilePlate(name string, origin string) (Plate, error) {
	fmt.Println("saving a copy of the file to tim directory...")
	timDir, err := TimDirectory()
	if err != nil {
		return nil, err
	}

	clonePath := filepath.Join(timDir, "files", name)
	err = CopyFile(origin, clonePath)
	if err != nil {
		return nil, err
	}

	fmt.Printf("copy saved to %v\n", clonePath)
	return &filePlate{
		name:   name,
		origin: origin,
		path:   clonePath,
	}, nil
}
