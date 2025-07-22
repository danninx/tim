package conf

import (
	"fmt"
	"os"
	"path"

	"github.com/pelletier/go-toml/v2"
)

type TOMLConfig struct {
	path string
}

func (file TOMLConfig) Read() (TimConfig, error) {
	dir, err := getTimDirectory()
	if err != nil {
		return TimConfig{}, err
	}

	full := path.Join(dir, file.path)

	b, err := os.ReadFile(full)
	if err != nil {
		return TimConfig{}, err
	}

	var config TimConfig;
	err = toml.Unmarshal(b, &config)
	if err != nil {
		return TimConfig{}, err
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
	
	fmt.Println(string(b))

	full := path.Join(dir, file.path)
	err = os.WriteFile(full, b, 0644)
	if err != nil {
		return err
	}

	return nil
}
