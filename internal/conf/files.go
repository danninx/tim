package conf

import (
	"bufio"
	"fmt"
	"os"
	"path"

	"github.com/danninx/tim/internal/system"
)

const DEFAULT_TYPE = "toml"
const DEFAULT_NAME = "tim.toml"

func newConfigFile(sys system.System) (ConfigFile, error) {
	dir, err := sys.TimDirectory()
	if err != nil {
		return nil, err
	}

	err = sys.TouchDir(dir)
	if err != nil {
		return nil, err
	}

	dot := path.Join(dir, ".tim")
	info, err := sys.Stat(dot)

	var filetype string
	if (err != nil && !info.IsDir()) { //
		fmt.Printf("no configuration file was found, creating a '%s' in %s\n", DEFAULT_NAME, dir)
		err = createDefault(sys)
		if err != nil {
			return nil, err
		}

		filetype = DEFAULT_TYPE
	} else {
		filetype, err = readDotfile(dot, sys)
		if err != nil {
			return nil, err
		}
	}

	if filetype == "toml" {
		return TOMLConfig{"tim.toml"}, nil
	} else {
		return nil, fmt.Errorf("an invalid filetype\"%s\" was found in %s, try renaming or removing any existing configuration files", filetype, dot)
	}
}

func createDefault(sys system.System) error {
	dir, err := sys.TimDirectory()
	if err != nil {
		return err
	}

	err = sys.TouchDir(dir)
	if err != nil {
		return err
	}

	defaultFilePath := path.Join(dir, DEFAULT_NAME)
	err = sys.TouchFile(defaultFilePath)
	if err != nil {
		return err
	}

	return SetConfFileType(DEFAULT_TYPE, sys)
}

func readDotfile(dotFilePath string, sys system.System) (string, error) {
	file, err := sys.OpenFile(dotFilePath, os.O_RDONLY|os.O_CREATE, 0777)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		fmt.Printf("no configuration file was found, using an empty '%s'\n", DEFAULT_NAME)
		err = createDefault(sys)
		if err != nil {
			return "", err
		}
		return DEFAULT_TYPE, nil
	} else {
		return scanner.Text(), nil
	}	
}
