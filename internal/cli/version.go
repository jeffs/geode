package cli

import (
	"errors"
	"fmt"
)

func Version(args []string) error {
	if len(args) != 0 {
		return errors.New("unexpected argument")
	}

	fmt.Println("v0.1.2")
	return nil
}
