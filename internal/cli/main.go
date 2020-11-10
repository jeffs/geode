// Package cli implements command-line parsing and subcommand dispatch.
package cli

import (
	"errors"
	"fmt"
	"io"
)

// Passes the specified args to the function implementing the specified cmd.
func dispatch(cmd string, args []string, wout, werr io.Writer) error {
	switch cmd {
	case "build":
		return nil // TODO
	case "dockerfile":
		return nil // TODO
	case "-h", "--help", "help":
		return Help(args, wout)
	default:
		return errors.New("bad command")
	}
}

// Main implements the geode command-line interface.
func Main(args []string, wout, werr io.Writer) int {
	if len(args) < 2 {
		fmt.Fprintln(werr, Usage)
		return 2
	}

	if err := dispatch(args[1], args[2:], wout, werr); err != nil {
		fmt.Fprintln(werr, "geode:", args[1] + ":", err)
		return 1
	}

	return 0
}
