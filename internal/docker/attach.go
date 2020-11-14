package docker

import (
	"os/exec"
)

func containerExists(name string) bool {
	return exec.Command("docker", "container", "inspect", name).Run() == nil
}

func volumeExists(flatName string) bool {
	return exec.Command("docker", "volume", "inspect", flatName).Run() == nil
}

func Attach(profile string, args []string) error {
	cfg, err := readConfig(profile)
	if err != nil {
		return err
	}

	if !volumeExists(cfg.FlatName()) {
		if err := RunFromConfig(profile, cfg, cfg.Init); err != nil {
			return err
		}
	}

	if !containerExists(cfg.FlatName()) {
		return RunFromConfig(profile, cfg, args)
	}

	if len(args) < 1 {
		args = cfg.Command
	}

	return Exec(profile, args)
}
