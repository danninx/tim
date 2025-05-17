package actions;

import ( 
	"bufio"
	"fmt"
	"os"
	"path"
	"slices"
	"strings"

	cli "github.com/danninx/tim/internals/cli"
	timfile "github.com/danninx/tim/internals/timfile"
	files "github.com/danninx/tim/internals/files"
)

const ANSI_BLUE = "\x1b[34m"
const ANSI_BOLD = "\x1b[22m"
const ANSI_GREEN = "\x1b[32m"
const ANSI_MAGENTA = "\x1b[35m"
const ANSI_RESET = "\x1b[0m"
const ANSI_WHITE = "\x1b[37m"
const ANSI_YELLOW = "\x1b[33m"

const GIT_WARNING = "\x1b[31mwarning: tim cannot verify the integrity of git urls, make sure you have the correct url and proper read access\x1b[0m"

func confirmAction(msg string) bool {
	fmt.Print(msg)
	reader := bufio.NewReader(os.Stdin)	
	response, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	response = strings.Replace(response, "\n", "", -1)
	return response == "y" || response == "Y"
}

func Add(command cli.Command) {
	stype, source := getSource(command)

	if len(command.Options) < 2 {
		fmt.Printf("tim - invalid number of arguments\n")	
		return
	}

	sources := timfile.Read()
	_, exists := sources[command.Options[1]]
	if exists {
		msg := fmt.Sprintf("%vsource \"%v\" already exists, would you like to replace it? (y/N)%v", ANSI_YELLOW, command.Options[1], ANSI_RESET)
		confirm := confirmAction(msg)
		if !confirm {
			fmt.Printf("skipping...")	
			return
		}
	}

	sources[command.Options[1]] = timfile.Src {
		Type: stype,
		Value: source,
	}

	timfile.Write(sources)

	fmt.Printf("%vadded \"%v\" to templates!%v\n", ANSI_GREEN, source, ANSI_RESET)
}

func Copy(command cli.Command) {
	if len(command.Options) != 3 {
		fmt.Printf("tim - invalid number of arguments\n")
		return
	}

	_, filterGit := command.Flags["filter-git"]

	sources := timfile.Read()
	source, exists := sources[command.Options[1]]
	if !exists {
		fmt.Printf("%vcould not find source \"%v\"%v\n", ANSI_YELLOW, command.Options[1], ANSI_RESET)	
		return 
	}

	if source.Type == "git" {
		if filterGit {
			tmp, err := files.TempGit(source.Value)
			if err != nil {
				fmt.Printf("%verror making temporary git clone:\n%v\n%v", ANSI_YELLOW, err, ANSI_RESET)
				return
			}
			fmt.Println("copying cleaned source...")
			err = files.CopyDir(tmp, command.Options[2])	
			if err != nil {
				fmt.Printf("%verror copying files from temporary clone:\n%v\n%v", ANSI_YELLOW, err, ANSI_RESET)
				return
			}
			fmt.Println("cleaning temporary directory...")
			err = files.CleanTmp()
			if err != nil {
				fmt.Printf("%vthere was an issue removing the temporary directory \"%v\":\n%v%v\n", ANSI_YELLOW, tmp, ANSI_RESET, err)
				return
			}
		} else {
			err := files.GitClone(source.Value, command.Options[2])
			if err != nil {
				fmt.Printf("%vgit encountered an error while copying source \"%v\"%v\n", ANSI_YELLOW, err, ANSI_RESET)
				return
			}
		}
	} else if source.Type == "file" {
		valid, src := pathExists(source.Value)
		if !valid {
			fmt.Printf("%vsource had invalid path \"%v\"%v\n", ANSI_YELLOW, src, ANSI_RESET)	
			return
		}
		err := files.CopyFile(src, command.Options[2])			
		if err != nil {
			panic(err)
		}
	} else if source.Type == "dir" {
		valid, src := pathExists(source.Value)
		if !valid {
			fmt.Printf("%vsource had invalid path \"%v\"%v\n", ANSI_YELLOW, src, ANSI_RESET)	
			return
		}

		if filterGit {
			tmp, err := files.TempCopy(src)	
			if err != nil {
				panic(err)
			}
			fmt.Println("copying cleaned source...")
			err = files.CopyDir(tmp, command.Options[2])
			if err != nil {
				panic(err)
			}
			fmt.Println("cleaning temporary directory...")
			err = files.CleanTmp()
			if err != nil {
				fmt.Printf("%vthere was an issue removing the temporary directory \"%v\":\n%v%v\n", ANSI_YELLOW, tmp, ANSI_RESET, err)
				return
			}
		} else {
			err := files.CopyDir(src, command.Options[2])			
			if err != nil {
				panic(err)
			}
		}
	} else {
		fmt.Printf("tim - found an unexpected source type %v for source \"%v\"", source.Type, command.Options[1])
		return
	}
}

func Edit(command cli.Command) {
	stype, source := getSource(command)

	if len(command.Options) < 2 {
		fmt.Printf("tim - invalid number of arguments\n")
		return
	}

	sources := timfile.Read()
	_, exists := sources[command.Options[1]]
	if exists {
		msg := fmt.Sprintf("%vare you sure you want to replace source \"%v\"? (y/N)%v", ANSI_YELLOW, command.Options[1], ANSI_RESET)
		confirm := confirmAction(msg)
		if !confirm {
			fmt.Printf("skipping...")
			return
		}
	} else {
		msg := fmt.Sprintf("%vsource \"%v\" does not yet exist, would you like to replace it? (y/N)%v", ANSI_YELLOW, command.Options[1], ANSI_RESET)
		confirm := confirmAction(msg)
		if !confirm {
			fmt.Printf("skipping...")
			return
		}
	}

	sources[command.Options[1]] = timfile.Src {
		Type: stype,
		Value: source,
	}

	timfile.Write(sources)
	if exists {
		fmt.Printf("%vmodified source \"%v\"!%v\n", ANSI_GREEN, source, ANSI_RESET)
	} else {
		fmt.Printf("%vadded \"%v\" to templates!%v\n", ANSI_GREEN, source, ANSI_RESET)
	}
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
	sources := timfile.Read()

	if len(command.Options) != 2 {
		fmt.Printf("not enough arguments, expected\n\t- tim add <src>\n")
		return
	}

	_, exists := sources[command.Options[1]]

	if !exists {
		fmt.Printf("%vcould not find source \"%v\"%v", ANSI_YELLOW, command.Options[1], ANSI_RESET)	
		return
	}

	msg := fmt.Sprintf("%vare you sure you want to delete source \"%v\"? (y/N)%v", ANSI_YELLOW, command.Options[1], ANSI_RESET)
	confirm := confirmAction(msg)
	if !confirm {
		fmt.Printf("skipping...")	
		return
	}

	delete(sources, command.Options[1])

	timfile.Write(sources)
}

func Help(command cli.Command) {
	// OVERVIEW
	fmt.Println("usage: tim <cmd> [-f | --file] [-d | --dir | --directory]\n\t\t [-h | --help] [-g | --git]")

	// SOURCES
	fmt.Println("\nmodifying sources:")
	fmt.Println("\tadd\t\tadd a template to tim")
	fmt.Println("\tedit | set\tchange the source for a template")
	fmt.Println("\tlist | ls\tls: list templates and their sources")
	fmt.Println("\tremove | rm\tremove a template from tim")
	fmt.Println("\tcopy | plate\tcopy a source to a given path")
	fmt.Println("\thelp\t\tshow this list")

	// EXAMPLES
	fmt.Println("\nexample usage:")
	fmt.Println("\ttim add <NAME> --dir <PATH>")
	fmt.Println("\ttim set <NAME> <URL>")
	fmt.Println("\ttim plate <NAME> <PATH>")
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

	//TODO verify source integrity	

	if !(hasDir || hasFile || hasGit) {
		if len(command.Options) != 3 {
			fmt.Printf("tim - invalid number of arguments\n")	
			os.Exit(0)
		}
		return "git", command.Options[2]
	} else if hasDir {
		valid, path := pathExists(dir)
		if !valid {
			os.Exit(0)
		}
		return "dir", path
	} else if hasFile {
		valid, path := pathExists(file)
		if !valid {
			os.Exit(0)
		}
		return "file", path
	} else if hasGit {
		fmt.Println(GIT_WARNING)
		return "git", git
	}

	panic("[tim/actions] - getSource() did not return a source")
}

func pathExists(p string) (bool, string) {
	clean := path.Clean(p)
	var full string 
	if strings.HasPrefix(clean, "/") {
		full = clean
	} else {
		wd, err := os.Getwd()
		if err != nil {
			panic("Could not get current working directory")
		}
		full = path.Join(wd, clean)
	}

	// check if os.Stat returns an error
	_, err := os.Stat(full)

	if err != nil {
		fmt.Printf("Invalid path found: %v\n", p)
	}

	return err == nil, full
}
