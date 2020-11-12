package docker

import (
	"os"
	"os/exec"
	"path"
	"strings"

	"golang.org/x/term"
)

func execArgs(name string) []string {
	a := []string{}
	if term.IsTerminal(int(os.Stdin.Fd())) {
		a = append(a, "--interactive", "--tty")
	}

	return append(a, name)
}

func Exec(profile string, args []string) error {
	_, name := path.Split(strings.TrimRight(profile, "/"))
	a := []string{"container", "exec"}
	a = append(a, execArgs(name)...)
	a = append(a, args...)
	c := exec.Command("docker", a...)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c.Run()
}
