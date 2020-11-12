package cli

import (
	"errors"

	"github.com/jeffs/geode/internal/docker"
)

func Build(args []string) error {
	if len(args) != 1 {
		return errors.New("expected exactly one profile")
	}

	return docker.Build(args[0])
}
