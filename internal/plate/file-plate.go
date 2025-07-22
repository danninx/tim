package plate

type filePlate struct {
	name		string
	origin		string
	path 		string
}

func (plate *filePlate) Name() (string) {

	return plate.name
}

func (plate *filePlate) Origin() (string) {
	
	return plate.origin
}

func (plate *filePlate) Path() (string) {
	return plate.path
}

func (plate *filePlate) Type() (string) {

	return "file"
}

func (plate *filePlate) Sync(destination string) error {
	return nil
}

func newFilePlate(name string, origin string, path string) (error, Plate) {
	return nil, &filePlate{
		name: path,
		origin: origin,
		path: path,
	}
}
