package plate

import (
	"fmt"
	"os"
	"path"
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
}

func ToString(plate Plate) string {
	return fmt.Sprintf("%v,%v", plate.Type(), plate.Origin())
}

func NewPlate(plateType string, name string, origin string) (Plate, error) {
	switch plateType {
	case "dir":
		return newDirPlate(name, origin)
	case "file":
		return newFilePlate(name, origin)
	case "git":
		return newGitPlate(name, origin)
	}
	return nil, fmt.Errorf("the specified plate type does not exist")
}

func Load(name string, unloaded UnloadedPlate) (Plate, error) {
	switch unloaded.Type {
	case "dir":
		return &dirPlate{
			name,
			unloaded.Origin,
			unloaded.Path,
		}, nil
	case "file":
		return &filePlate{
			name,
			unloaded.Origin,
			unloaded.Path,
		}, nil
	case "git":
		return &gitPlate{
			name,
			unloaded.Origin,
			unloaded.Path,
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

// utility

func TimDirectory() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return path.Join(home, ".config/tim"), nil
}
