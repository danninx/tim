package conf

import (
	"bufio"
	"fmt"
	"os"
	"path"
)

const DEFAULT_TYPE = "toml"

func newConfigFile() (ConfigFile, error) {
	dir, err := getTimDirectory()
	if err != nil {
		return nil, err
	}
	err = mkDirIfNotExists(dir)
	if err != nil {
		return nil, err
	}

	dot := path.Join(dir, ".tim")
	info, err := os.Stat(dot)

	var filetype string
	if !(err == nil && !info.IsDir()) { //
		fmt.Printf("no configuration file was found, creating a 'tim.toml' in %s\n", dir)
		err = touchFile(path.Join(dir, "tim.toml"))
		if err != nil {
			return nil, err
		}

		filetype = "toml"
		err = SetConfFileType(DEFAULT_TYPE)
		if err != nil {
			return nil, err
		}
	} else {
		file, err := os.OpenFile(dot, os.O_RDONLY|os.O_CREATE, 0777)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		if !scanner.Scan() {
			return nil, fmt.Errorf("no format detected, try using 'tim set-file [filetype]'. valid filetypes are:\n%s", validFiletypes())
		}
		filetype = scanner.Text()
	}

	switch filetype {
	case "toml":
		return TOMLConfig{"tim.toml"}, nil
	default:
		return nil, fmt.Errorf("an invalid filetype \"%s\" was found in %s, try using 'tim set-file [filetype]'. valid filetypes are: \n%s", filetype, dot, validFiletypes())
	}
}

func getTimDirectory() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return path.Join(home, ".config/tim"), nil
}

func validFiletypes() string {
	return "toml"
}

func mkDirIfNotExists(path string) error {
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

func touchFile(path string) error {
	_, err := os.Stat(path)
	if err == nil {
		fmt.Printf("file %v exists", path)
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
