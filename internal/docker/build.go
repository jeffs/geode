package docker

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"
)

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

	_, name := path.Split(strings.TrimRight(profile, "/"))
	c := exec.Command("docker", "image", "build", "-t", name, dir)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	if err := c.Run(); err != nil {
		return err
	}

	return nil
}

// Build creates a Docker image per the specified Geode profile directory.  The
// image tag is the directory's basename; i.e., final component of the path.
func Build(profile string) error {
	cfg, err := readConfig(profile)
	if err != nil {
		return err
	}

	return BuildFromConfig(profile, cfg)
}
