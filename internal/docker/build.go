package docker

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
)

func volumeExists(flatName string) bool {
	return exec.Command("docker", "volume", "inspect", flatName).Run() == nil
}

func BuildFromConfig(profile string, cfg *config) error {
	dir, err := ioutil.TempDir("", "geode-build")
	if err != nil {
		return fmt.Errorf("can't create temp dir %s: %w\n", dir, err)
	}

	defer os.RemoveAll(dir)

	file, err := os.Create(path.Join(dir, "Dockerfile"))
	if err != nil {
		return fmt.Errorf("can't create Dockerfile: %w\n", err)
	}

	if err := DockerfileFromConfig(profile, cfg, file); err != nil {
		return err
	}

	c := exec.Command("docker", "image", "build", "-t", cfg.Name, dir)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	if err := c.Run(); err != nil {
		return err
	}

	if !volumeExists(cfg.Name) {
		nc := *cfg
		abs, err := filepath.Abs(profile)
		if err != nil {
			return err
		}

		nc.Bind[abs] = "/mnt/profile"
		return RunFromConfig(profile, &nc, []string{"/mnt/profile/init"})
	}

	return nil
}

// Build creates a Docker image per the specified Geode profile directory.
func Build(profile string) error {
	cfg, err := readConfig(profile)
	if err != nil {
		return err
	}

	return BuildFromConfig(profile, cfg)
}
