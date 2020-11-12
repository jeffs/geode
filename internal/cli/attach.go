package cli

import (
	"errors"

	"github.com/jeffs/geode/internal/docker"
)

func Attach(args []string) error {
	if len(args) < 1 {
		return errors.New("expected profile")
	}

	return docker.Attach(args[0], args[1:])
}
