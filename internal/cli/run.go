package cli

import (
	"errors"

	"github.com/jeffs/geode/internal/docker"
)

func Run(args []string) error {
	if len(args) != 1 {
		return errors.New("expected exactly one profile")
	}

	return docker.Run(args[0])
}
