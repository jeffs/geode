package cli

import (
	"errors"
	"os"

	"github.com/jeffs/geode/internal/docker"
)

type config struct {
	Base string
	Name string
	User string

	Locale   string
	TimeZone string
}

func Dockerfile(args []string) error {
	if len(args) != 1 {
		return errors.New("expected exactly one profile")
	}

	return docker.ExpandFile(args[0], os.Stdout)
}
