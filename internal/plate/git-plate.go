package plate

type gitPlate struct {
	name		string
	origin		string
	path 		string
}

func (plate *gitPlate) Name() (string) {

	return plate.name
}

func (plate *gitPlate) Origin() (string) {
	
	return plate.origin
}

func (plate *gitPlate) Path() (string) {
	return plate.path
}

func (plate *gitPlate) Type() (string) {

	return "git"
}

func (plate *gitPlate) Sync(destination string) error {
	return nil
}

func newGitPlate(name string, origin string, path string) (error, Plate) {
	return nil, &gitPlate{
		name: path,
		origin: origin,
		path: path,
	}
}
