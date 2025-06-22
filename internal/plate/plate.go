package plate

import (
	"fmt"
)

type Plate struct {
	Name 	string
	Type 	string
	Path 	string
}

func ToString(plate Plate) string {
	return fmt.Sprintf("%v,%v", plate.Type, plate.Path)
}

