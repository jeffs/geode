package docker

import (
	"os"
	"os/exec"
	"path"
	"strings"

	"golang.org/x/term"
)

func dockerArgs(name string) []string {
	a := []string{"container", "run", "--name", name, "--rm"}
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

	args = append(dockerArgs(name), args...)
	cmd := exec.Command("docker", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
