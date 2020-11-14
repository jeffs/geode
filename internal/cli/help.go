// Package cli implements Geode subcommands.
//
// TODO: Document the TOML schema.
package cli

import (
	"errors"
	"fmt"
	"strings"
)

const Usage string = `Geode is a tool for managing personal config files.

Usage:

    geode COMMAND [ARGUMENTS]

Commands:

    build       Build a Docker image from a Geode profile
    dockerfile  Print a Dockerfile from a Geode profile
    help        Print this message
    run         Start a container, or attach to an existing one

Run 'geode help COMMAND' for information about that command.`

func topic_help(topic string) (error, string) {
	switch topic {
	case "build":
		return nil, `usage: geode build PROFILE

		Builds a Docker image from Geode profile.  To see the
		Dockerfile contents without actually building the image, use:
		geode dockerfile PROFILE`

	case "dockerfile":
		return nil, `usage: geode dockerfile PROFILE

		Prints a Dockerfile from a Geode profile.  The profile must be
		a directory containing a Dockerfile template and a docker.toml
		file.  Variables in the template are replaced by corresponding
		values set in the toml file.`

	case "help":
		return nil, "usage: geode help [TOPIC]"

	case "run":
		return nil, `usage: geode run [-n] PROFILE [ARGS...]

		Starts a new Docker container from the specified Geode profile,
		building the image first if it does not already exist. With -n,
		only prints the command that would be used to run the
		container, without actually running it.`

	default:
		return errors.New(topic + ": bad topic"), ""
	}
}

// Help prints a usage message, or describes specified topics.
func Help(args []string) error {
	if len(args) == 0 {
		fmt.Println(Usage)
		return nil
	}

	if len(args) > 1 {
		return errors.New("expected a single topic")
	}

	err, help := topic_help(args[0])
	if err != nil {
		return err
	}

	help = strings.ReplaceAll(help, "\t", "")
	fmt.Println(help)
	return nil
}
