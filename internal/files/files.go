// Timfile copy source code for copying files and directories
package files;

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
)

// Returns path /../tmp/tim
func timTmp() string {
	return filepath.Join(os.TempDir(), "tim")
}

// Copy file from src path to dest path
func CopyFile(src string, dest string) error {
	srcFile, err := os.OpenFile(src, os.O_RDONLY, 0777)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// check if destination path is an existing directory, if so append source file name, otherwise copy as though destination is a file
	var destFilePath string
	destInfo, err := os.Stat(dest)	
	if err != nil {
		if (os.IsNotExist(err)) {
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

	

	destFile, err := os.OpenFile(destFilePath, os.O_WRONLY | os.O_CREATE, 0777)
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

// Copy dir from src path to dest path
func CopyDir(src string, dest string) error {
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

func GitClone(src string, dest string) error {
	fmt.Printf("using git to copy source \"%v\"\n", src)	
	cmd := exec.Command("git", "clone", src, dest)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fmt.Println(cmd)
	return cmd.Run()	
}

// Makes a directory clone of `src`, and then removes any present .git directory
func TempCopy(src string) (string, error) { 
	CleanTmp()
	dest := timTmp()	
	err := CopyDir(src, dest)	
	if err != nil { return "", err }
	gitPath := filepath.Join(dest, ".git")	
	err = os.RemoveAll(gitPath)
	return dest, err
}

// Makes a git clone of `src`, and then removes the .git directory
func TempGit(src string) (string, error) {
	CleanTmp()
	dest := timTmp()
	err := GitClone(src, dest)
	if err != nil { return "", err }
	gitPath := filepath.Join(dest, ".git")
	err = os.RemoveAll(gitPath)
	return dest, err
}

func CleanTmp() error {
	return os.RemoveAll(timTmp())
} 
