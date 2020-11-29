package docker

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

func BuildFromConfig(profile string, noCache bool, cfg *config) error {
	dir, err := ioutil.TempDir("", "geode-build")
	if err != nil {
		return fmt.Errorf("can't create temp dir %s: %w\n", dir, err)
	}

	defer os.RemoveAll(dir)

	file, err := os.Create(filepath.Join(dir, "Dockerfile"))
	if err != nil {
		return fmt.Errorf("can't create Dockerfile: %w\n", err)
	}

	if err := DockerfileFromConfig(profile, cfg, file); err != nil {
		return err
	}

	a := []string{"image", "build", "-t", cfg.Name}
	if noCache {
		a = append(a, "--no-cache")
	}

	c := exec.Command("docker", append(a, dir)...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	if err := c.Run(); err != nil {
		return err
	}

	return nil
}

// Build creates a Docker image per the specified Geode profile directory.  If
// noCache is true, Build passes the --no-cache flag to the underlying 'docker
// build' command.
func Build(profile string, noCache bool) error {
	cfg, err := readConfig(profile)
	if err != nil {
		return err
	}

	return BuildFromConfig(profile, noCache, cfg)
}
