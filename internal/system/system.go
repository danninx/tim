package system

import (
	"os"
	"testing"
)

type System interface {
	CopyDir(src string, dst string) error
	CopyFile(src string, dst string) error
	GitClone(src string, dest string) error
	OpenFile(path string, flag int, perm os.FileMode) (*os.File, error)
	ReadFile(path string) ([]byte, error)
	RemoveAll(path string) error
	TimDirectory() (string, error)
	TouchDir(path string) error
	TouchFile(path string) error
	Stat(path string) (os.FileInfo, error)
	WriteFile(path string, data []byte, perm os.FileMode) error
}

func GetSystem() System {
	if testing.Testing() {
		return FakeSystem{}
	} else {
		return RealSystem{}
	}
}

func Real() System {
	return RealSystem{}
}

func Fake() System {
	return FakeSystem{}
}

// func Mock() System {
// 	return MockSystem{}
// }
