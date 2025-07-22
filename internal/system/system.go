package system

import "testing"

type System interface {
	CopyDir(src string, dst string) error
	CopyFile(src string, dst string) error
	GitClone(src string, dest string) error
	RemoveAll(path string) error
	TimDirectory() (string, error)
}

func GetSystem() System {
	if testing.Testing() {
		return FakeSystem{}
	} else {
		return RealSystem{}
	}
}
