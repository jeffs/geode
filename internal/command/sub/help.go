// Package sub implements geode subcommands.
package sub

import (
	"fmt"

	"github.com/jeffs/geode/internal/command/errs"
)

const Usage string = `Geode is a tool for managing personal config files.

Usage:

    geode <command> [arguments]

Commands:

    help                Print this message
    docker.build        Build a Docker image from a TOML file
    docker.start        Run a command in a Docker container
    docker.file         Print the Dockerfile docker.build would use

Run 'geode help <command>' for more information about a command.
`

// Help prints a usage message, or describes specified topics.
//
// TODO: Inject stdout.
func Help(args []string) error {
	if len(args) == 0 {
		fmt.Println(Usage)
		return nil
	}

	if len(args) > 1 {
		return errs.User{"help: expected a single topic"}
	}

	arg := args[0]
	switch arg {
	case "help":
		fmt.Println("usage: geode help [topic]")
	default:
		return errs.User{"help: " + arg + ": bad topic"}
	}

	return nil
}
