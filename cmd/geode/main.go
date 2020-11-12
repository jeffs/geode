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
	case "attach":
		return cli.Attach(args)
	case "build":
		return cli.Build(args)
	case "dockerfile":
		return cli.Dockerfile(args)
	case "help":
		return cli.Help(args)
	case "run":
		// TODO: Implemnt Run as Attach.  Do not expose Run directly.
		return cli.Run(args)
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
	}

	if err := dispatch(c, os.Args[2:]); err != nil {
		fmt.Fprintln(os.Stderr, "geode:", c+":", err)
		os.Exit(1)
	}
}
