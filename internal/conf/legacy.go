package conf

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/danninx/tim/internal/plate"
)

type LegacyConfig struct {
	path string
}

func (legacy LegacyConfig) Read() (TimConfig, error) {
	dir, err := getTimDirectory()
	if err != nil {
		return TimConfig{}, err
	}

	full := path.Join(dir, legacy.path)
	file, err := os.OpenFile(full, os.O_RDONLY|os.O_CREATE, 0777)
	if err != nil {
		return TimConfig{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sources := map[string]plate.Plate{}

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "=")
		if len(parts) != 2 {
			fmt.Printf("improperly formatted source was found:\n\t\"%v\"\n", line)
			continue
		}
		split := strings.Split(parts[1], ",")
		if len(split) != 2 {
			fmt.Printf("improperly formatted source was found:\n\t\"%v\"\n", line)
			continue
		}
		s := plate.Plate{
			Type: split[0],
			Path: split[1],
		}
		sources[parts[0]] = s
	}

	config := TimConfig{
		Options: TimOptions{},
		Plates:  sources,
	}
	return config, nil
}

func (legacy LegacyConfig) Write(config TimConfig) error {
	dir, err := getTimDirectory()
	if err != nil {
		return err
	}

	full := path.Join(dir, legacy.path)
	file, err := os.OpenFile(full, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	defer file.Close()

	for k := range config.Plates {
		val := plate.ToString(config.Plates[k])
		fmt.Fprintf(file, "%v=%v\n", k, val)
	}

	return nil
}
