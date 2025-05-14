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
		"l": "local",
		"-local": "local",
		"f": "file",
		"file": "file",
		"d": "directory",
		"dir": "directory",
		"g": "git",
		"git": "git",
	}
	silents := map [string] bool {
		"l": true,
		"-local": true,
	}

	cmd := cli.ParseArgs(
		arguments,
		flagPrefix,
		flags,
		silents,
	)
	fmt.Println(cli.CommandString(cmd))

	subcommands := map [string] func(cli.Command) { 
		"add": actions.Add,
		"edit": actions.Edit,
		"ls": actions.List,
		"rm": actions.Remove,
		"help": actions.Help,
		"testwrite": actions.TestWrite,
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

