package timfile;

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)


const TIM_FILE_NAME = "/.timfile"
type Src struct {
	Type 	string
	Value 	string
}

func Read() map [string] Src {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Panic(err)
	}

	full := home + TIM_FILE_NAME
	file, err := os.OpenFile(full, os.O_RDONLY | os.O_CREATE, 0777)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sources := map [string] Src {}

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
		s := Src{
			Type: split[0],
			Value: split[1],
		}
		sources[parts[0]] = s
	}
	return sources
}

func Write(sources map [string] Src) {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Panic(err)
	}

	full := home + TIM_FILE_NAME
	file, err := os.OpenFile(full, os.O_WRONLY | os.O_TRUNC | os.O_CREATE, 0777)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	for k := range sources {
		val := srcToString(sources[k])
		fmt.Fprintf(file, "%v=%v\n", k, val)
	}
}

func srcToString(source Src) string {
	return fmt.Sprintf("%v,%v", source.Type, source.Value)
}
