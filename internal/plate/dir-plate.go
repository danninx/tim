package plate

type dirPlate struct {
	name		string
	origin		string
	path 		string
}

func (plate *dirPlate) Name() (string) {

	return plate.name
}

func (plate *dirPlate) Origin() (string) {
	
	return plate.origin
}

func (plate *dirPlate) Path() (string) {
	return plate.path
}

func (plate *dirPlate) Type() (string) {

	return "dir"
}

func (plate *dirPlate) Sync(destination string) error {
	return nil
}

func newDirPlate(name string, origin string, path string) (error, Plate) {
	return nil, &dirPlate{
		name: path,
		origin: origin,
		path: path,
	}
}
