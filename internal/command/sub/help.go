// Package sub implements geode subcommands.
package sub

import (
	"fmt"
	"io"

	"github.com/jeffs/geode/internal/command/errs"
)

const Usage string = `Geode is a tool for managing personal config files.

Usage:

    geode COMMAND [ARGUMENTS]

Commands:

    help                Print this message
    docker.image.build  Build a Docker image from a TOML profile
    docker.image.file   Print a Docker file from a TOML profile

Run 'geode help COMMAND' for information about that command.

Other help topics:

Run 'geode help TOPIC' for information about that topic.
`

const dockerImageBuildUsage string = `usage: geode docker.image.build PROFILE.toml

Build a Docker image from PROFILE.toml.  To see the Dockerfile contents without
actually building the image, : geode docker.file PROFILE.toml
`

// TODO: Document TOML schema.
const dockerFileUsage string = `usage: geode docker.image.file PROFILE.toml

Print a Docker file built from a TOML profile.
`

// Help prints a usage message, or describes specified topics.
//
// TODO: Inject stdout.
func Help(args []string, wout io.Writer) error {
	if len(args) == 0 {
		fmt.Fprintln(wout, Usage)
		return nil
	}

	if len(args) > 1 {
		return errs.User{"expected a single topic"}
	}

	switch arg := args[0]; arg {
	case "docker.image.build":
		fmt.Fprintln(wout, "usage: geode docker.image.build TOML")
	case "help":
		fmt.Fprintln(wout, "usage: geode help [TOPIC]")
	default:
		return errs.User{arg + ": bad topic"}
	}

	return nil
}
