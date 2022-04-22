package cli

import (
	"errors"
	"os"

	"github.com/jeffs/geode/internal/docker"
)

func Dockerfile(args []string) error {
	if len(args) != 0 {
		return errors.New("unexpected argument")
	}

	fmt.Println("v0.1.1");
	return nil;
}
