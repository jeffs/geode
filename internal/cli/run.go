package cli

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jeffs/geode/internal/docker"
)

func Run(args []string) error {
	var profile string
	var dryRun bool
	var noCache bool
	var command []string // what to run in the container
	for _, arg := range args {
		// Args after the profile path are for the command, not for us.
		if profile == "" && strings.HasPrefix(arg, "-") {
			switch arg {
			default:
				return errors.New("bad flag: " + arg + "\n\n    Did you mean -n, or --no-cache?")
			case "-n":
				dryRun = true
			case "--no-cache":
				noCache = true
			}
		} else if profile == "" {
			profile = arg
		} else {
			command = append(command, arg)
		}
	}

	if profile == "" {
		return errors.New("expected a profile path")
	}

	if dryRun {
		a, err := docker.RunCommand(profile, command)
		if err != nil {
			return err
		}

		fmt.Printf("%v\n", a)
		return nil
	}

	// We implement Run as Attach because, unlike Docker, Geode does not
	// directly support simultaneous containers running from the same
	// image.  Instead, Geode encourages simultaneous connections to a
	// single container, starting the container automatically if need be;
	// so Run is more like 'tmux attach' than it is like 'docker run' or
	// 'docker exec'.
	return docker.Attach(profile, noCache, command)
}
