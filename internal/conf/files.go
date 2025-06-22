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
	if (err != nil) {
		return nil, err
	}
	dot := path.Join(dir, ".tim")
	info, err := os.Stat(dot)

	var filetype string
	if !(err == nil && !info.IsDir()) { //
		fmt.Printf("no configuration file was found, creating a 'tim.toml' in %s\n", dir)
		err = SetConfFileType(DEFAULT_TYPE)
		if (err != nil) {
			return nil, err
		}
		filetype = DEFAULT_TYPE
	} else {
		file, err := os.OpenFile(dot, os.O_RDONLY | os.O_CREATE, 0777)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		if !scanner.Scan() {
			return nil, fmt.Errorf("no format detected, try using 'tim set-file [filetype]'. valid filetypes are:\n%s\n", validFiletypes())
		}
		filetype = scanner.Text()
	}

	switch filetype {
	case "legacy":
		return LegacyConfig{"tim.conf"}, nil
	case "toml":
		return TOMLConfig{"tim.toml"}, nil
	default:
		return nil, fmt.Errorf("an invalid filetype \"%s\" was found in %s, try using 'tim set-file [filetype]'. valid filetypes are: \n%s\n", filetype, dot, validFiletypes())
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
	return "legacy, toml"
}
