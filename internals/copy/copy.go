// Timfile copy source code for copying files and directories
package copy;

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Copy file from src path to dest path
func CopyFile(src string, dest string) error {
	srcFile, err := os.OpenFile(src, os.O_RDONLY, 0777)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	destFile, err := os.OpenFile(dest, os.O_WRONLY | os.O_CREATE, 0777)
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

	return os.Chmod(dest, info.Mode())
}

// Copy dir from src path to dest path
func CopyDir(src string, dest string) error {
	srcInfo, err := os.Stat(src)	
	if err != nil {
		return err
	}

	if !srcInfo.IsDir() {
		return errors.New(fmt.Sprintf("source path \"%v\" was not a directory\n", src))
	}

	err = os.MkdirAll(dest, srcInfo.Mode())
	if err != nil {
		return err
	}

	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		destPath := filepath.Join(dest, entry.Name())

		if entry.IsDir() {
			err := CopyDir(srcPath, destPath)
			if err != nil {
				return err
			}
		} else {
			err := CopyFile(srcPath, destPath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
