package conf

import (
	"path"

	"github.com/danninx/tim/internal/plate"
	"github.com/danninx/tim/internal/system"
	"github.com/pelletier/go-toml/v2"
)

type TOMLConfig struct {
	path string
}

func (file TOMLConfig) Read(sys system.System) (TimConfig, error) {
	dir, err := system.GetSystem().TimDirectory()
	if err != nil {
		return emptyConfig(), err
	}

	full := path.Join(dir, file.path)

	b, err := sys.ReadFile(full)
	if err != nil {
		return emptyConfig(), err
	}

	var config TimConfig
	err = toml.Unmarshal(b, &config)
	if err != nil {
		return emptyConfig(), err
	}

	if config.Plates == nil {
		config.Plates = make(map[string]plate.UnloadedPlate)
	}
	return config, nil
}

func (file TOMLConfig) Write(config TimConfig, sys system.System) error {
	dir, err := sys.TimDirectory()
	if err != nil {
		return err
	}

	b, err := toml.Marshal(config)
	if err != nil {
		return err
	}

	full := path.Join(dir, file.path)
	err = sys.WriteFile(full, b, 0644)
	if err != nil {
		return err
	}

	return nil
}
