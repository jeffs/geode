// Geode is a program to manage your personal config files: shell, editor, etc.
package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/jeffs/geode/internal/cli"
)

// Passes the specified args to the function implementing the specified command.
func dispatch(command string, args []string) error {
	switch command {
	case "build":
		return cli.Build(args)
	case "dockerfile":
		return cli.Dockerfile(args)
	case "help":
		return cli.Help(args)
	case "run":
		return cli.Run(args)
	case "version":
		return cli.Version(args)
	default:
		return errors.New("bad command")
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, cli.Usage)
		os.Exit(2)
	}

	c := os.Args[1]
	if c == "-h" || c == "--help" {
		c = "help"
	} else if c == "-V" || c == "--version" {
		c = "version"
	}

	if err := dispatch(c, os.Args[2:]); err != nil {
		fmt.Fprintln(os.Stderr, "geode:", c+":", err)
		os.Exit(1)
	}
}
