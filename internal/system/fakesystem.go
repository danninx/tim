package system

import "os"

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

func (_ FakeSystem) OpenFile(path string, flag int, perm os.FileMode) (*os.File, error) {
	return nil, nil
}

func (_ FakeSystem) ReadFile(path string) ([]byte, error) {
	return []byte{}, nil
}

func (_ FakeSystem) RemoveAll(path string) error {
	return nil
}

func (_ FakeSystem) Stat(path string) (os.FileInfo, error) {
	return nil, nil
}

func (_ FakeSystem) TimDirectory() (string, error) {
	return "timTest", nil
}

func (_ FakeSystem) TouchDir(path string) error {
	return nil
}

func (_ FakeSystem) TouchFile(path string) error {
	return nil
}

func (_ FakeSystem) WriteFile(path string, data []byte, perm os.FileMode) error {
	return nil
}
