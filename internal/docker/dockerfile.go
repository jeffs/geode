// Package dockerfile supports interaction with Docker images and containers.
package docker

import (
	"fmt"
	"io"
	"io/ioutil"
	"path"
	"text/template"

	"github.com/BurntSushi/toml"
)

func readConfig(file string) (*Config, error) {
	var cfg Config
	if _, err := toml.DecodeFile(file, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

// WriteFile reads a Dockerfile template and dockerfile.toml from the specified
// profile directory, and expands the template with the values from the TOML
// file to the specified writer.
func ExpandFile(profile string, w io.Writer) error {
	cfg, err := readConfig(path.Join(profile, "dockerfile.toml"))
	if err != nil {
		return err
	}

	bytes, err := ioutil.ReadFile(path.Join(profile, "Dockerfile"))
	if err != nil {
		return err
	}

	tpl := template.Must(template.New("Dockerfile").Parse(string(bytes)))
	if err := tpl.Execute(w, &cfg); err != nil {
		return fmt.Errorf("can't generate Dockerfile: %w\n", err)
	}

	return nil
}
