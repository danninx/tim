package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"tim/pkgs/cli"
)

const ANSI_RED = "\x1b[31m"
const ANSI_GREEN = "\x1b[32m"
const ANSI_WHITE = "\x1b[37m"
const ANSI_RESET = "\x1b[0m"

type Flag struct {
	name 	string
	value 	string
}

func ensureTimPath() {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal( err )
	}
	timPath := path.Join(dirname, "/.config/tim")
	err = os.MkdirAll(timPath, os.ModePerm)
	if err != nil {
		log.Fatal( err )
	}
}

func timHelp() {
	fmt.Printf("Implement tim help!")	
}

func checkGitUrl(url string) bool {
	_, err := exec.Command("git", "ls-remote", url, "HEAD").CombinedOutput()
	return err == nil
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		timHelp()
		os.Exit(0)
	}
	
}
