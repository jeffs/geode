package cli

import (
	"errors"
	"strings"

	"github.com/jeffs/geode/internal/docker"
)

func Build(args []string) error {
	var profile string
	var noCache bool
	for _, arg := range args {
		if strings.HasPrefix(arg, "-") {
			if arg != "--no-cache" {
				return errors.New("bad flag: " + arg + "\n\n    Did you mean --no-cache?")
			}

			noCache = true
		} else {
			if profile != "" {
				return errors.New(arg + ": unexpected path; already given " + profile)
			}

			profile = arg
		}
	}

	if profile == "" {
		return errors.New("expected a profile path")
	}

	return docker.Build(profile, noCache)
}
