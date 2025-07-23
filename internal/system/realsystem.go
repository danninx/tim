package system

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"path"
	"path/filepath"
)

type RealSystem struct{}

func (_ RealSystem) CopyDir(src string, dest string) error {
	fsys := os.DirFS(src)
	return fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		fpath, err := filepath.Localize(path)
		if err != nil {
			return err
		}
		newPath := filepath.Join(dest, fpath)
		if d.IsDir() {
			return os.MkdirAll(newPath, 0777)
		}

		if !d.Type().IsRegular() {
			return &os.PathError{Op: "CopyFS", Path: path, Err: os.ErrInvalid}
		}

		r, err := fsys.Open(path)
		if err != nil {
			return err
		}
		defer r.Close()
		info, err := r.Stat()
		if err != nil {
			return err
		}
		w, err := os.OpenFile(newPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666|info.Mode()&0777)
		if err != nil {
			return err
		}

		if _, err := io.Copy(w, r); err != nil {
			w.Close()
			return &os.PathError{Op: "Copy", Path: newPath, Err: err}
		}
		return w.Close()
	})
}

func (_ RealSystem) CopyFile(src string, dest string) error {
	srcFile, err := os.OpenFile(src, os.O_RDONLY, 0777)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// check if destination path is an existing directory, if so append source file name, otherwise copy as though destination is a file
	var destFilePath string
	destInfo, err := os.Stat(dest)
	if err != nil {
		if os.IsNotExist(err) {
			destFilePath = dest
		} else {
			return err
		}
	} else {
		if destInfo.IsDir() {
			destFilePath = filepath.Join(dest, filepath.Base(src))
		} else {
			destFilePath = dest
		}
	}

	destFile, err := os.OpenFile(destFilePath, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return err
	}

	info, err := os.Stat(src)
	if err != nil {
		return err
	}

	return os.Chmod(destFilePath, info.Mode())
}

func (_ RealSystem) GitClone(src string, dest string) error {
	fmt.Printf("using git to copy source \"%v\"\n", src)
	cmd := exec.Command("git", "clone", src, dest)
	cmd.Stdout = nil // make git silent unless error
	cmd.Stderr = os.Stderr
	fmt.Println(cmd)
	return cmd.Run()
}

func (_ RealSystem) OpenFile(path string, flag int, perm os.FileMode) (*os.File, error) {
	return os.OpenFile(path, flag, perm)
}

func (_ RealSystem) ReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func (_ RealSystem) RemoveAll(path string) error {
	return os.RemoveAll(path)
}

func (_ RealSystem) Stat(path string) (os.FileInfo, error) {
	return os.Stat(path)
}

func (_ RealSystem) TimDirectory() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return path.Join(home, ".config/tim"), nil
}


func (_ RealSystem) TouchDir(path string) error {
	file, err := os.Stat(path)
	if err == nil {
		if file.IsDir() {
			return nil
		}
		return fmt.Errorf("tim directory exists as a file, but is not a directory")
	}
	if !os.IsNotExist(err) {
		return err
	}

	return os.Mkdir(path, 0740)
}

func (_ RealSystem) TouchFile(path string) error {
	_, err := os.Stat(path)
	if err == nil {
		// file already exists, skip
		return nil
	}
	if !os.IsNotExist(err) {
		return err
	}

	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	fmt.Fprintln(file, "")
	fmt.Printf("created file %v\n", path)
	return nil
}

func (_ RealSystem) WriteFile(path string, data []byte, perm os.FileMode) error {
	return os.WriteFile(path, data, perm)
}
