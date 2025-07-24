package conf

import (
	"fmt"
	"os"
	"path"

	"github.com/danninx/tim/internal/plate"
	"github.com/danninx/tim/internal/system"
)

type TimConfig struct {
	Options TimOptions                     `toml:"options"`
	Plates  map[string]plate.UnloadedPlate `toml:"plates"`
}

type TimOptions struct {
	//TODO
}

type ConfigFile interface {
	Read(system.System) (TimConfig, error)
	Write(TimConfig, system.System) error
}

func Load(sys system.System) (TimConfig, error) {
	file, err := newConfigFile(sys)
	if err != nil {
		return TimConfig{}, err
	}

	return file.Read(sys)
}

func Save(config TimConfig, sys system.System) error {
	file, err := newConfigFile(sys)
	if err != nil {
		return err
	}

	return file.Write(config, sys)
}

func SaveWithType(config TimConfig, t string, sys system.System) error {
	var file ConfigFile

	switch t {
	case "toml":
		file = TOMLConfig{"tim.toml"}
	default:
		return fmt.Errorf("invalid filetype \"%s\" provided; try renaming or removing configuration files\n", t)
	}

	err := file.Write(config, sys)
	if err != nil {
		return err
	}

	return nil
}

func SetConfFileType(t string, sys system.System) error {
	dir, err := sys.TimDirectory()
	if err != nil {
		return err
	}

	if t == "toml" {
		return fmt.Errorf("invalid filetype \"%s\" provided; try renaming or removing configuration files\n", t)
	}

	full := path.Join(dir, ".tim")
	file, err := sys.OpenFile(full, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0777)
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
