package main

import (
	"fmt"
	"os"
)

var usage string = `Geode is a tool for managing config files.

Usage:

    geode <command> [arguments]

Commands:

    help        Print this message
    update      Download and install the latest version of Geode

Run 'geode help <command>' for more information about a command.
`

type userError struct {
	what string
}

func (e userError) Error() string {
	return e.what
}

func handleHelp(args []string) error {
	if len(args) > 1 {
		return userError{"help: expected a single topic"}
	}

	if len(args) == 0 {
		fmt.Println(usage)
		return nil
	}

	switch args[0] {
	case "help":
		fmt.Println("usage: geode help [topic]")
	case "update":
		fmt.Println(`usage: geode update [-n]

The update command installs the latest version of Geode.  With -n, it
prints the version that would be installed but does not install it.`)
	default:
		return userError{"help: " + args[0] + ": bad topic"}
	}

	return nil
}

func handleUpdate(args []string) error {
	return nil	// TODO
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println(usage)
		os.Exit(2)
	}

	cmd, args := os.Args[1], os.Args[2:]
	var err error
	switch cmd {
	case "help":
		err = handleHelp(args)
	case "update":
		err = handleUpdate(args)
	default:
		err = userError{what: cmd + ": bad command"}
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "geode: %v\n", err)
		if _, ok := err.(userError); ok {
			os.Exit(2)
		}

		os.Exit(1)
	}
}
