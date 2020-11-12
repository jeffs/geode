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
	case "build":
		return cli.Build(args)
	case "dockerfile":
		return cli.Dockerfile(args)
	case "-h", "--help", "help":
		return cli.Help(args)
	default:
		return errors.New("bad command")
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, cli.Usage)
		os.Exit(2)
	}

	if err := dispatch(os.Args[1], os.Args[2:]); err != nil {
		fmt.Fprintln(os.Stderr, "geode:", os.Args[1]+":", err)
		os.Exit(1)
	}
}
