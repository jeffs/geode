package cli

import (
	"errors"
	"fmt"
	"io"
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

Run 'geode help TOPIC' for information about that topic.`

func topic_help(topic string) (error, string) {
	switch topic {
	case "dockerfile":
		// TODO: Document TOML schema.
		return nil, `usage: geode dockerfile PROFILE.toml

Print a Docker file built from a TOML profile.`
	case "build":
		return nil, `usage: geode build PROFILE.toml

Build a Docker image from PROFILE.toml.  To see the Dockerfile contents
without actually building the image, use: geode dockerfile PROFILE.toml`
	case "help":
		return nil, "usage: geode help [TOPIC]"
	default:
		return errors.New(topic + ": bad topic"), ""
	}
}

// Help prints a usage message, or describes specified topics.
//
// TODO: Inject stdout.
func Help(args []string, wout io.Writer) error {
	if len(args) == 0 {
		fmt.Fprintln(wout, Usage)
		return nil
	}

	if len(args) > 1 {
		return errors.New("expected a single topic")
	}

	err, help := topic_help(args[0])
	if err != nil {
		return err
	}

	fmt.Fprintln(wout, help)
	return nil
}
