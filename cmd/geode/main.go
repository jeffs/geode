// Geode is a program to manage your personal config files: shell, editor, etc.
package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/jeffs/geode/internal/cli"
)

// Passes the specified args to the function implementing the specified cmd.
func dispatch(cmd string, args []string) error {
	switch cmd {
	case "attach":
		return cli.Attach(args)
	case "build":
		return cli.Build(args)
	case "dockerfile":
		return cli.Dockerfile(args)
	case "help":
		return cli.Help(args)
	case "run":
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

	cmd := os.Args[1]
	if cmd == "-h" || cmd == "--help" {
		cmd = "help"
	}

	if err := dispatch(cmd, os.Args[2:]); err != nil {
		fmt.Fprintln(os.Stderr, "geode:", cmd+":", err)
		os.Exit(1)
	}
}
