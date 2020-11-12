package docker

import (
	"os"
	"os/exec"

	"golang.org/x/term"
)

func execArgs(cfg *config) []string {
	a := []string{}
	if term.IsTerminal(int(os.Stdin.Fd())) {
		a = append(a, "--interactive", "--tty")
	}

	return append(a, cfg.FlatName())
}

func Exec(profile string, args []string) error {
	cfg, err := readConfig(profile)
	if err != nil {
		return err
	}

	a := []string{"container", "exec"}
	a = append(a, execArgs(cfg)...)
	a = append(a, args...)
	c := exec.Command("docker", a...)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c.Run()
}
