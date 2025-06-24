package plate

import (
	"fmt"
)

type PlateData struct {
	Type 		string
	Origin		string
	LocalPath 	string
}

type Plate interface {
	CopyTo(string) 	error
	LocalPath() 	string
	ToString() 		string
}

func ToString(plate PlateData) string {
	return fmt.Sprintf("%v,%v", plate.Type, plate.Origin)
}

/*
	Git plate type
*/

type GitPlate struct {
	Data PlateData
}

func (plate GitPlate) CopyTo(dir string) error {

	return nil
}

func (plate GitPlate) ToString() string {
	return ToString(plate.Data)
}

/*
	Dir plate type
*/

type DirPlate struct {
	Data PlateData
}

func (plate DirPlate) ToString() string {
	return ToString(plate.Data)
}

/*
	File plate type
*/

type FilePlate struct {
	Data PlateData
}


func (plate FilePlate) ToString() string {
	return ToString(plate.Data)
}
