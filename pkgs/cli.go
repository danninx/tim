package cli

type Flag struct {
	name 	string
	value 	string
}

type Command struct {
	cmd		[]string
	flags	[]Flag
}

func parseCommand(input []string) Command {
	return Command{[], []}
}
