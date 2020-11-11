package cli

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"text/template"

	"github.com/BurntSushi/toml"
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

	var cfg config
	if _, err := toml.DecodeFile(path.Join(args[0], "dockerfile.toml"), &cfg); err != nil {
		return err
	}

	bytes, err := ioutil.ReadFile(path.Join(args[0], "Dockerfile"))
	if err != nil {
		return err
	}

	tpl := template.Must(template.New("Dockerfile").Parse(string(bytes)))
	if err := tpl.Execute(os.Stdout, &cfg); err != nil {
		return fmt.Errorf("can't generate Dockerfile: %w\n", err)
	}

	return nil
}
