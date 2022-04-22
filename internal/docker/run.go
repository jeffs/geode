package docker

import (
	"fmt"
	"os"
	"os/exec"

	"golang.org/x/term"
)

func imageExists(name string) bool {
	return exec.Command("docker", "image", "inspect", name).Run() == nil
}

func runArgs(cfg *config) []string {
	fn := cfg.FlatName()
	a := []string{
		"--env=DISPLAY=host.docker.internal:0",
		"--hostname=" + fn,
		"--mount=type=volume,source=" + fn + ",target=/home/" + cfg.User,
		"--name=" + fn,
		"--rm",
		"--tmpfs=/tmp",
	}

	for k, v := range cfg.Bind {
		a = append(a, "--mount=type=bind,src="+k+",dst="+v)
	}

	for k, v := range cfg.Volumes {
		a = append(a, "--mount=type=volume,src="+k+",dst="+v)
	}

	for k, v := range cfg.Ports {
		a = append(a, fmt.Sprintf("--publish=%s:%d", k, v))
	}

	if term.IsTerminal(int(os.Stdin.Fd())) {
		a = append(a, "--interactive", "--tty")
	}

	return append(a, cfg.Name)
}

func RunCommandFromConfig(cfg *config, args []string) []string {
	a := []string{"docker", "container", "run"}
	a = append(a, runArgs(cfg)...)
	a = append(a, args...)

	return a
}

// Returns arguments suitable for exec.Command.
func RunCommand(profile string, args []string) ([]string, error) {
	cfg, err := readConfig(profile)
	if err != nil {
		return nil, err
	}

	return RunCommandFromConfig(cfg, args), nil
}

func RunFromConfig(profile string, noCache bool, cfg *config, args []string) error {
	if noCache || !imageExists(cfg.Name) {
		if err := BuildFromConfig(profile, noCache, cfg); err != nil {
			return err
		}
	}

	a := RunCommandFromConfig(cfg, args)
	c := exec.Command(a[0], a[1:]...)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	return c.Run()
}
