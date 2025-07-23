package system

type FakeSystem struct {}

func (_ FakeSystem) CopyDir(src string, dst string) error {
	return nil
}

func (_ FakeSystem) CopyFile(src string, dst string) error {
	return nil
}

func (_ FakeSystem) GitClone(src string, dst string) error {
	return nil
}

func (_ FakeSystem) RemoveAll(path string) error {
	return nil
}

func (_ FakeSystem) TimDirectory() (string, error) {
	return "timTest", nil
}
