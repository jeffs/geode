package cli

import (
	"errors"
	"fmt"

	"github.com/jeffs/geode/internal/docker"
)

func Run(args []string) error {
	if len(args) < 1 {
		return errors.New("expected profile")
	}

	if len(args) > 1 && args[0] == "-n" {
		if len(args) < 2 {
			return errors.New("-n: expected profile")
		}

		a, err := docker.RunCommand(args[1], args[2:])
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
	return docker.Attach(args[0], args[1:])
}
