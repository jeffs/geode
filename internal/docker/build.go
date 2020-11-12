package docker

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"
)

func volumeExists(name string) bool {
	return exec.Command("docker", "volume", "inspect", name).Run() == nil
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

	_, name := path.Split(strings.TrimRight(profile, "/"))
	c := exec.Command("docker", "image", "build", "-t", name, dir)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	if err := c.Run(); err != nil {
		return err
	}

	// TODO: If the user's home volume doesn't exist, create and initialize it.
	// We'll do this by running a container mounting the geode-init script
	// and host ssh creds, and specifying the script as the run command.
	//
	// TODO: Support alternative ssh creds per toml.
	if !volumeExists(name) {
		fmt.Fprintln(os.Stderr, "TODO:", name + ":", "initialize volume")
	}

	return nil
}

// Build creates a Docker image according to the specified Geode profile
// directory.  The image tag is the directory's basename; i.e., the final
// component of the path.
func Build(profile string) error {
	cfg, err := readConfig(profile)
	if err != nil {
		return err
	}

	return BuildFromConfig(profile, cfg)
}
