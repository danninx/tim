package plate

import (
	"fmt"
	"os"
	"path/filepath"
)

type gitPlate struct {
	name   string
	origin string
	path   string
}

func (plate *gitPlate) Name() string {
	return plate.name
}

func (plate *gitPlate) Origin() string {
	return plate.origin
}

func (plate *gitPlate) Path() string {
	return plate.path
}

func (plate *gitPlate) Type() string {
	return "git"
}

func (plate *gitPlate) Sync() error {
	err := os.RemoveAll(plate.path)
	if err != nil {
		return err
	}

	err = GitClone(plate.origin, plate.path)
	if err != nil {
		return err
	}

	fmt.Printf("succesfully synchronized \"%v\" from repository", plate.name)
	return nil
}

func (plate *gitPlate) Copy(destination string) error {
	return CopyDir(plate.path, destination)
}

func (plate *gitPlate) Delete() error {
	return os.RemoveAll(plate.path)
}

func newGitPlate(name string, origin string) (Plate, error) {
	fmt.Println("cloning repository for offline use...")
	timDir, err := TimDirectory()
	if err != nil {
		return nil, err
	}
	clonePath := filepath.Join(timDir, "git", name)
	err = GitClone(origin, clonePath)
	if err != nil {
		return nil, err
	}
	fmt.Printf("offline version saved at %v\n", clonePath)
	return &gitPlate{
		name:   name,
		origin: origin,
		path:   clonePath,
	}, nil
}
