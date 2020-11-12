package docker

import (
	"os"
	"os/exec"
	"path"
	"strings"

	"golang.org/x/term"
)

func imageExists(name string) bool {
	return exec.Command("docker", "image", "inspect", name).Run() == nil
}

// TODO: Mount, publish, --security-opt=seccomp:unconfined, etc. per config.
func runArgs(name string, cfg *config) []string {
	a := []string{
		"--env=DISPLAY=host.docker.internal:0",
		"--hostname=" + name,
		"--mount=type=volume,source=" + name + ",target=/home/" + cfg.User,
		"--name=" + name,
		"--rm",
	}

	if term.IsTerminal(int(os.Stdin.Fd())) {
		a = append(a, "--interactive", "--tty")
	}

	return append(a, name)
}

func Run(profile string, args []string) error {
	cfg, err := readConfig(profile)
	if err != nil {
		return err
	}
	_, name := path.Split(strings.TrimRight(profile, "/"))
	if !imageExists(name) {
		if err := BuildFromConfig(profile, cfg); err != nil {
			return err
		}
	}

	a := []string{"container", "run"}
	a = append(a, runArgs(name, cfg)...)
	a = append(a, args...)
	c := exec.Command("docker", a...)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c.Run()
}
