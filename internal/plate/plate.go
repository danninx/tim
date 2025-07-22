package plate

import (
	"fmt"
)

type Plate interface {
	Name()						(string)
	Origin() 					(string)
	Path()						(string)
	Type() 						(string)
	Sync(destination string)	(error)
}

type OldPlate struct {
	Type 	string
	Origin	string
	Path 	string
}

func ToString(plate Plate) string {
	return fmt.Sprintf("%v,%v", plate.Type(), plate.Origin())
}

func NewPlate(plateType string, name string, origin string, path string) (error, Plate) {
	switch plateType {
	case "dir":
		return newDirPlate(name, origin, path)
	case "file":
		return newFilePlate(name, origin, path)
	case "git":
		return newGitPlate(name, origin, path)
	}
	return fmt.Errorf("An invalid plate type was provided"), nil
}
