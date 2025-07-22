package plate

import (
	"fmt"
	"path/filepath"

	"github.com/danninx/tim/internal/system"
)

type gitPlate struct {
	name   	string
	origin 	string
	path   	string
	sys		system.System
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
	err := plate.sys.RemoveAll(plate.path)
	if err != nil {
		return err
	}

	err = plate.sys.GitClone(plate.origin, plate.path)
	if err != nil {
		return err
	}

	fmt.Printf("succesfully synchronized \"%v\" from repository", plate.name)
	return nil
}

func (plate *gitPlate) Copy(destination string) error {
	return plate.sys.CopyDir(plate.path, destination)
}

func (plate *gitPlate) Delete() error {
	return plate.sys.RemoveAll(plate.path)
}

func newGitPlate(name string, origin string, sys system.System) (Plate, error) {
	fmt.Println("cloning repository for offline use...")
	timDir, err := sys.TimDirectory()
	if err != nil {
		return nil, err
	}
	clonePath := filepath.Join(timDir, "git", name)
	err = sys.GitClone(origin, clonePath)
	if err != nil {
		return nil, err
	}
	fmt.Printf("offline version saved at %v\n", clonePath)
	return &gitPlate{
		name:   name,
		origin: origin,
		path:   clonePath,
		sys:	sys,
	}, nil
}
