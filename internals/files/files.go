// Timfile copy source code for copying files and directories
package copy;

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

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
	dest := timTmp()	
	err := CopyDir(src, dest)	
	if err != nil { return "", err }
	gitPath := filepath.Join(dest, ".git")	
	err = os.RemoveAll(gitPath)
	return dest, err
}

// Makes a git clone of `src`, and then removes the .git directory
func TempGit(src string) (string, error) {
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
