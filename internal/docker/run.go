package docker

import (
	"os"
	"os/exec"
	"path"
	"strings"

	"golang.org/x/term"
)

// TODO: mount, publish, "--security-opt=seccomp:unconfined" if necessary
func runArgs(name string) []string {
	a := []string{
		"--env=DISPLAY=host.docker.internal:0",
		"--hostname=" + name,
		"--name=" + name,
		"--rm",
	}

	if term.IsTerminal(int(os.Stdin.Fd())) {
		a = append(a, "--interactive", "--tty")
	}

	return append(a, name)
}

func imageExists(name string) bool {
	return exec.Command("docker", "image", "inspect", name).Run() == nil
}

func Run(profile string, args []string) error {
	_, name := path.Split(strings.TrimRight(profile, "/"))
	if !imageExists(name) {
		if err := Build(profile); err != nil {
			return err
		}
	}

	a := []string{"container", "run"}
	a = append(a, runArgs(name)...)
	a = append(a, args...)
	c := exec.Command("docker", a...)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c.Run()
}
