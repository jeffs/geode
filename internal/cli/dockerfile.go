package cli

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path"
)

func Dockerfile(args []string) error {
	if len(args) != 1 {
		return errors.New("expected exactly one profile")
	}

	profile := args[0]

	bytes, err := ioutil.ReadFile(path.Join(profile, "Dockerfile"))
	if err != nil {
		return err
	}

	fmt.Println(string(bytes))

	return nil
}
