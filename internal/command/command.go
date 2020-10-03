// Package command implements command-line parsing and subcommand dispatch.
package command

import (
	"fmt"
	"os"

	"github.com/jeffs/geode/internal/command/errs"
	"github.com/jeffs/geode/internal/command/sub"
)

// dispatch finds the function implementing the named cmd, and passes that
// function the specified args.
func dispatch(cmd string, args []string) error {
	switch cmd {
	case "help":
		return sub.Help(args)
	default:
		return errs.User{cmd + ": bad command"}
	}
}

// Main implements the geode command-line interface.  Returns 2 on usage error,
// 1 on any other error, and 0 on success.
func Main(osArgs []string) int {
	if len(osArgs) < 2 {
		fmt.Println(sub.Usage) // TODO: Inject stdout.
		return 2
	}

	if err := dispatch(osArgs[1], osArgs[2:]); err != nil {
		fmt.Fprintf(os.Stderr, "geode: %v\n", err) // TODO: Inject stderr.
		if _, ok := err.(errs.User); ok {
			return 2
		}

		return 1
	}

	return 0
}
