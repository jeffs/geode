package cli

import (
	"errors"
	"os"

	"github.com/jeffs/geode/internal/docker"
)

func Dockerfile(args []string) error {
	if len(args) != 1 {
		return errors.New("expected exactly one profile")
	}

	return docker.Dockerfile(args[0], os.Stdout)
}
