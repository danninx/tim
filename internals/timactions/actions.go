package timactions;

import ( "fmt"
	"os"
	"slices"
	cli "github.com/danninx/tim/pkgs/cliparse"
	timfile "github.com/danninx/tim/internals/timfile"
)

const ANSI_BLUE = "\x1b[34m"
const ANSI_BOLD = "\x1b[22m"
const ANSI_GREEN = "\x1b[32m"
const ANSI_MAGENTA = "\x1b[35m"
const ANSI_RESET = "\x1b[0m"
const ANSI_WHITE = "\x1b[37m"
const ANSI_YELLOW = "\x1b[33m"

const GIT_WARNING = "\x1b[31mwarning: tim cannot verify the integrity of git urls, make sure you have the correct url and proper read access\x1b[0m"

func Add(command cli.Command) {
	stype, source := getSource(command)

	if len(command.Options) < 2 {
		fmt.Printf("tim - invalid number of arguments\n")	
	}

	sources := timfile.Read()
	_, exists := sources[command.Options[1]]
	if exists {
		var response string;
		fmt.Printf("%vsource \"%v\" already exists, would you like to replace it? (y/n)%v", ANSI_YELLOW, command.Options[1], ANSI_RESET)
		_, err := fmt.Scanln(&response)
		if err != nil {
			panic(err)
		}

		if response != "y" {
			fmt.Printf("skipping...")	
			return
		}
	}

	sources[command.Options[1]] = timfile.Src {
		Type: stype,
		Value: source,
	}

	timfile.Write(sources)

	fmt.Printf("%vadded %v to templates!%v", ANSI_GREEN, source, ANSI_RESET)
}

func addErr() {
	fmt.Println("IMPLEMENT TIM ADD ERR MSG")
}

func Edit(command cli.Command) {
	fmt.Println("IMPLEMENT EDIT")	
}

func List(command cli.Command) {
	fmt.Println("tim sources:")
	sources := timfile.Read()

	var keys[] string
	for k := range sources {
		keys = append(keys, k)
	}
	slices.Sort(keys)

	for _, k := range keys {
		printSrc(k, sources[k])
	}
}

func Remove(command cli.Command) {
	fmt.Println("IMPLEMENT REMOVE")
}

func Help(command cli.Command) {
	fmt.Printf("%vusage: tim [cmd] [options]\n%v", ANSI_WHITE, ANSI_RESET)
	fmt.Println("\tcommands:")
	fmt.Println("\t\t- add <file|src>:add a template to tim")
	fmt.Println("\t\t- edit <name>: change the source for a template")
	fmt.Println("\t\t- ls: list templates and their sources")
	fmt.Println("\t\t- rm <name>: remove a template from tim")
	fmt.Println("\t\t- help: show this list")
	fmt.Println("\toptions:")
	fmt.Println("\t--local\t\tuse a file or directory path as source (source defaults to git)")
}

func TestWrite(command cli.Command) {
	sources := map [string] timfile.Src {}
	sources["hello"] = timfile.Src {
		Type: "git",
		Value: "world",
	}

	sources["test"] = timfile.Src {
		Type: "file",
		Value: "/home/danninx/booglydooglydoo",
	}

	sources["dadadadir"] = timfile.Src {
		Type: "dir",
		Value: "/home/danninx/awooga/",
	}

	timfile.Write(sources)	
}

func printSrc(name string, source timfile.Src) {
	n := fmt.Sprintf("%v%v%-15s%v", ANSI_BOLD, ANSI_GREEN, name, ANSI_RESET)
	p := fmt.Sprintf("%-s\n", source.Value)
	var t string
	switch source.Type {
	case "git":
		t = fmt.Sprintf("%s%-8s%s", ANSI_YELLOW, source.Type, ANSI_RESET)
	case "dir":
		t = fmt.Sprintf("%s%-8s%s", ANSI_BLUE, source.Type, ANSI_RESET)
	case "file":
		t = fmt.Sprintf("%s%-8s%s", ANSI_MAGENTA, source.Type, ANSI_RESET)
	}
	fmt.Printf("\t- %-s%-s%-s", n, t, p)
}

func getSource(command cli.Command) (string, string) {
	dir, hasDir	:= command.Flags["directory"]
	file, hasFile := command.Flags["file"]
	git, hasGit := command.Flags["git"]

	validFlags := (!hasDir && !hasFile) || (!hasDir && !hasGit) || (!hasFile && !hasGit)
	if !validFlags {
		fmt.Printf("tim - you cannot specify more than one type of source\n")
		os.Exit(0)
	}

	if !(hasDir || hasFile || hasGit) {
		if len(command.Options) != 3 {
			fmt.Printf("tim - invalid number of arguments\n")	
			os.Exit(0)
		}
		return "git", command.Options[2]
	} else if hasDir {
		return "dir", dir
	} else if hasFile {
		return "file", file
	} else if hasGit {
		fmt.Println(GIT_WARNING)
		return "git", git
	}

	panic("[tim/actions] - getSource() did not return a source")
	return "", ""
}
