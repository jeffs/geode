package docker

import (
	"os/exec"
	"path"
	"strings"
)

func containerExists(name string) bool {
	return exec.Command("docker", "container", "inspect", name).Run() == nil
}

func Attach(profile string, args []string) error {
	_, name := path.Split(strings.TrimRight(profile, "/"))
	if !containerExists(name) {
		return Run(name, args)
	}

	if len(args) < 1 {
		cfg, err := readConfig(profile)
		if err != nil {
			return err
		}

		args = cfg.Command
	}

	return Exec(profile, args)
}
