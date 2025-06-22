package timfile

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/danninx/tim/internal/plate"
)

const TIM_FILE_NAME = "/.timfile"
type Src struct {
	Type 	string
	Value 	string
}

func Read() (map [string] plate.Plate, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Panic(err)
	}

	full := home + TIM_FILE_NAME
	file, err := os.OpenFile(full, os.O_RDONLY | os.O_CREATE, 0777)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sources := map [string] plate.Plate {}

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
		s := plate.Plate {
			Type: split[0],
			Path: split[1],
		}
		sources[parts[0]] = s
	}
	return sources, nil
}

func Write(sources map [string] plate.Plate) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	full := home + TIM_FILE_NAME
	file, err := os.OpenFile(full, os.O_WRONLY | os.O_TRUNC | os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	defer file.Close()

	for k := range sources {
		val := plate.ToString(sources[k])
		fmt.Fprintf(file, "%v=%v\n", k, val)
	}

	return nil
}

