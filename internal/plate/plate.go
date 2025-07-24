package plate

import (
	"fmt"

	"github.com/danninx/tim/internal/system"
)

type UnloadedPlate struct {
	Type   string
	Origin string
	Path   string
}

type Plate interface {
	Name() string
	Origin() string
	Path() string
	Type() string
	Sync() error
	Copy(destination string) error
	Delete() error
	Rename(newName string) error
}

func NewPlate(plateType string, name string, origin string, sys system.System) (Plate, error) {
	switch plateType {
	case "dir":
		return newDirPlate(name, origin, sys)
	case "file":
		return newFilePlate(name, origin, sys)
	case "git":
		return newGitPlate(name, origin, sys)
	}
	return nil, fmt.Errorf("the specified plate type does not exist")
}

func Load(name string, unloaded UnloadedPlate, sys system.System) (Plate, error) {
	switch unloaded.Type {
	case "dir":
		return &dirPlate{
			name,
			unloaded.Origin,
			unloaded.Path,
			sys,
		}, nil
	case "file":
		return &filePlate{
			name,
			unloaded.Origin,
			unloaded.Path,
			sys,
		}, nil
	case "git":
		return &gitPlate{
			name,
			unloaded.Origin,
			unloaded.Path,
			sys,
		}, nil
	}
	return nil, fmt.Errorf("failed to load unknown plate type")
}

func Unload(plate Plate) UnloadedPlate {
	return UnloadedPlate{
		plate.Type(),
		plate.Origin(),
		plate.Path(),
	}
}
