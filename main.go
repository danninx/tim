package main

import (
	"fmt"
	"os"
	cli "github.com/danninx/tim/internals/cli"
	actions "github.com/danninx/tim/internals/timactions"
)


func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		actions.Help(cli.Command{})
		os.Exit(0)
	}

	flagPrefix := "-"
	flags := map [string] string { 
		"f": "file",
		"-file": "file",
		"d": "directory",
		"-dir": "directory",
		"-directory": "directory",
		"g": "git",
		"-git": "git",
		"-debug": "debug",
	}
	silents := map [string] bool {}

	cmd := cli.ParseArgs(
		arguments,
		flagPrefix,
		flags,
		silents,
	)
	_, debug := cmd.Flags["debug"]
	if debug { fmt.Println(cli.CommandString(cmd)) }

	subcommands := map [string] func(cli.Command) { 
		"add": actions.Add,

		"copy": actions.Copy,
		"plate": actions.Copy, // for the pun ...

		"edit": actions.Edit,
		"set": actions.Edit,

		"list": actions.List,
		"ls": actions.List,

		"rm": actions.Remove,
		"help": actions.Help,
		"testwrite": actions.TestWrite, // TODO remove this when not needed
	}

	if len(cmd.Options) == 0 {
		actions.Help(cmd)
		os.Exit(0)
	}

	commandAction := cmd.Options[0]
	f, exists := subcommands[commandAction]
	if exists {
		f(cmd)
	} else {
		panic(fmt.Sprintf("tim - unrecognized action %v\n", commandAction))
	}
}

