package timcli;

import (
	"fmt"
	"os"
	"strings"
)

type Command struct {
	Options		[]string
	Flags		map [string] string
}

func CommandString(command Command) string {
	return fmt.Sprintf("Command{%v, %v}", command.Options, command.Flags)
}

func isSilent(str string, silents map [string] bool) bool {
	_, found := silents[str]
	return found
}

func isFlag(str string, prefix string, valid map [string] string) (string, bool) {
	after, found := strings.CutPrefix(str, prefix)
	if found { 
		_, v := valid[after]
		if v {
			return after, true
		} else {
			fmt.Printf("tim - Invalid flag \"%v\"\n", str)
			return "", false
		}
	} else {
		return str, false
	}
}

func ParseArgs(args []string, flagPrefix string, validFlags map [string] string, silents map [string] bool) Command {
	flags := map [string] string {}
	options := []string{}

	for i := 1; i < len(args); i++ {
		val, validFlag := isFlag(args[i], flagPrefix, validFlags)
		if val == "" {os.Exit(0)}
		if validFlag {
			silent := isSilent(val, silents)	
			if silent {
				flags[validFlags[val]] = "true"
			} else {
				i++
				flags[validFlags[val]] = args[i]
			}
		} else {
			options = append(options, val)
		}
	}

	return Command{options, flags}
}
