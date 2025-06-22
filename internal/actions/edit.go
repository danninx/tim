package actions

/*
	Needs to be revamped; since plates will be cached locally in the future, edits should be made using a text editor
*/

/*
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
*/
