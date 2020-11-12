// Package dockerfile supports interaction with Docker images and containers.
package docker

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"path"
	"text/template"
)

func jsonArray(elems []string) (string, error) {
	b, err := json.Marshal(elems)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func DockerfileFromConfig(profile string, cfg *config, w io.Writer) error {
	bytes, err := ioutil.ReadFile(path.Join(profile, "Dockerfile"))
	if err != nil {
		return err
	}

	tpl := template.Must(
		template.
			New("Dockerfile").
			Funcs(template.FuncMap{"ja": jsonArray}).
			Parse(string(bytes)))

	if err := tpl.Execute(w, &cfg); err != nil {
		return fmt.Errorf("can't generate Dockerfile: %w\n", err)
	}

	return nil
}

// WriteFile reads a Dockerfile template and docker.toml from the specified
// profile directory, and expands the template with the values from the TOML
// file to the specified writer.
func Dockerfile(profile string, w io.Writer) error {
	cfg, err := readConfig(profile)
	if err != nil {
		return err
	}

	return DockerfileFromConfig(profile, cfg, w)
}
