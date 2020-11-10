// Package command implements command-line parsing and subcommand dispatch.
package command

import (
	"fmt"
	"io"

	"github.com/jeffs/geode/internal/command/errs"
	"github.com/jeffs/geode/internal/command/sub"
)

// dispatch finds the function implementing the named cmd, and passes that
// function the specified args.
func dispatch(cmd string, args []string, wout, werr io.Writer) error {
	switch cmd {
	case "docker.image.build":
		return nil // TODO
	case "docker.file":
		return nil // TODO
	case "help":
		return sub.Help(args, wout)
	default:
		return errs.User{cmd + ": bad command"}
	}
}

// Main implements the geode command-line interface.  Returns 2 on usage error,
// 1 on any other error, and 0 on success.
func Main(osArgs []string, wout, werr io.Writer) int {
	if len(osArgs) < 2 {
		fmt.Fprint(werr, sub.Usage)
		return 2
	}

	cmd := osArgs[1]
	if err := dispatch(cmd, osArgs[2:], wout, werr); err != nil {
		fmt.Fprintf(werr, "geode: %s: %v\n", cmd, err)
		if _, ok := err.(errs.User); ok {
			return 2
		}

		return 1
	}

	return 0
}
