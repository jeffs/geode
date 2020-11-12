package docker

import (
	"os"
	"os/exec"
	"path"
	"strings"

	"golang.org/x/term"
)

func dockerArgs(name string) []string {
	dargs := []string{"container", "run", "--name", name, "--rm"}
	if term.IsTerminal(int(os.Stdin.Fd())) {
		dargs = append(dargs, "--interactive", "--tty")
	}

	return append(dargs, name)
}

func imageExists(name string) bool {
	return exec.Command("docker", "image", "inspect", name).Run() == nil
}

func Run(profile string) error {
	_, name := path.Split(strings.TrimRight(profile, "/"))
	if !imageExists(name) {
		if err := Build(profile); err != nil {
			return err
		}
	}

	cmd := exec.Command("docker", dockerArgs(name)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
