package actions

import (
	"fmt"
)

type NO_PLATE_EXISTS struct {
	Name string
}

func (e *NO_PLATE_EXISTS) Error() string {
	return fmt.Sprintf("no plate found with name \"%v\"", e.Name)
}

type INVALID_PLATE_TYPE struct {
	Type string
}

func (e *INVALID_PLATE_TYPE) Error() string {
	return fmt.Sprintf("plate type \"%v\" does not exist", e.Type)
}

type INVALID_NUM_ARGS struct {
	Expected int
	Received int
}

func (e *INVALID_NUM_ARGS) Error() string {
	return fmt.Sprintf("invalid number of arguments - expected %d, received %d", e.Expected, e.Received)
}
