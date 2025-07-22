package conf

import (
	"fmt"
	"os"
	"path"

	"github.com/danninx/tim/internal/plate"
)

type TimConfig struct {
	Options TimOptions                     `toml:"options"`
	Plates  map[string]plate.UnloadedPlate `toml:"plates"`
}

type TimOptions struct {
	//TODO
}

type ConfigFile interface {
	Read() (TimConfig, error)
	Write(TimConfig) error
}

func Load() (TimConfig, error) {
	file, err := newConfigFile()
	if err != nil {
		return TimConfig{}, err
	}

	return file.Read()
}

func Save(config TimConfig) error {
	file, err := newConfigFile()
	if err != nil {
		return err
	}

	return file.Write(config)
}

func SaveWithType(config TimConfig, t string) error {
	var file ConfigFile

	switch t {
	case "toml":
		file = TOMLConfig{"tim.toml"}
	default:
		return fmt.Errorf("invalid filetype \"%s\" provided. valid filetypes are: \n%s", t, validFiletypes())
	}

	err := file.Write(config)
	if err != nil {
		return err
	}

	return nil
}

func SetConfFileType(t string) error {
	dir, err := getTimDirectory()
	if err != nil {
		return err
	}

	if !(t == "legacy" || t == "toml") {
		return fmt.Errorf("invalid filetype \"%s\" provided. valid filetypes are: \n%s", t, validFiletypes())
	}

	full := path.Join(dir, ".tim")
	file, err := os.OpenFile(full, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Fprintln(file, t)
	file.Close()

	return nil
}

func emptyConfig() TimConfig {
	return TimConfig{
		Options: TimOptions{},
		Plates:  make(map[string]plate.UnloadedPlate),
	}
}
