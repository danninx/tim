package conf

import (
	"os"
	"path"

	"github.com/danninx/tim/internal/plate"
	"github.com/pelletier/go-toml/v2"
)

type TOMLConfig struct {
	path string
}

func (file TOMLConfig) Read() (TimConfig, error) {
	dir, err := getTimDirectory()
	if err != nil {
		return emptyConfig(), err
	}

	full := path.Join(dir, file.path)

	b, err := os.ReadFile(full)
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

func (file TOMLConfig) Write(config TimConfig) error {
	dir, err := getTimDirectory()
	if err != nil {
		return err
	}

	b, err := toml.Marshal(config)
	if err != nil {
		return err
	}

	full := path.Join(dir, file.path)
	err = os.WriteFile(full, b, 0644)
	if err != nil {
		return err
	}

	return nil
}
