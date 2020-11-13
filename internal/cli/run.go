package cli

import (
	"errors"

	"github.com/jeffs/geode/internal/docker"
)

func Run(args []string) error {
	if len(args) < 1 {
		return errors.New("expected profile")
	}

	// We implement Run as Attach because, unlike Docker, Geode does not
	// directly support simultaneous containers running from the same
	// image.  Instead, Geode encourages simultaneous connections to a
	// single container, starting the container automatically if need be;
	// so Run is more like 'tmux attach' than it is like 'docker run' or
	// 'docker exec'.
	return docker.Attach(args[0], args[1:])
}
