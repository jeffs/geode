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

func Attach(profile string, noCache bool, command []string) error {
	cfg, err := readConfig(profile)
	if err != nil {
		return err
	}

	if !volumeExists(cfg.FlatName()) {
		if err := RunFromConfig(profile, noCache, cfg, cfg.Init); err != nil {
			return err
		}

		// If we've just rebuilt the image, don't rebuild it yet again
		// when (if) we create the new container.
		noCache = false
	}

	if !containerExists(cfg.FlatName()) {
		return RunFromConfig(profile, noCache, cfg, command)
	}

	if len(command) < 1 {
		command = cfg.Command
	}

	return Exec(profile, command)
}
