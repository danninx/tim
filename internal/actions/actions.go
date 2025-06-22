package actions

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/danninx/tim/internal/plate"
	"github.com/danninx/tim/internal/timfile"
	"github.com/urfave/cli/v3"
)

/*
ANSI escape codes for output
*/

const ANSI_BLUE = "\x1b[34m"
const ANSI_GREEN = "\x1b[32m"
const ANSI_MAGENTA = "\x1b[35m"
const ANSI_RESET = "\x1b[0m"
const ANSI_WHITE = "\x1b[37m"
const ANSI_YELLOW = "\x1b[33m"

const GIT_WARNING = "\x1b[31mwarning: tim cannot verify the integrity of git urls, make sure you have the correct url and proper read access\x1b[0m"

/*
Generally helpful functions
*/

func CheckPathExists(p string) (string, error) {
	clean := path.Clean(p)
	var full string 
	if strings.HasPrefix(clean, "/") {
		full = clean
	} else {
		wd, err := os.Getwd()
		if err != nil {
			return "", err
		}
		full = path.Join(wd, clean)
	}

	_, err := os.Stat(full)

	if err != nil {
		return "", err
	}

	return full, nil
}

func ConfirmAction(msg string) (bool, error) {
	fmt.Print(msg)
	reader := bufio.NewReader(os.Stdin)	
	response, err := reader.ReadString('\n')
	if err != nil {
		return false, err
	}
	response = strings.Replace(response, "\n", "", -1)
	return response == "y" || response == "Y", nil
}

func EnforceNumArgs(expected int, cmd *cli.Command) error {
	if cmd.Args().Len() != expected {
		return &INVALID_NUM_ARGS{ Expected: expected, Received: cmd.Args().Len() }
	}

	return nil
}

func GetPlate(name string) (plate.Plate, error) {
	plates, err := timfile.Read()
	if (err != nil) {
		return plate.Plate{}, err
	}

	plate, exists := plates[name]	
	if !exists {
		return plate, &NO_PLATE_EXISTS{ Name: name }
	}

	return plate, nil
}


