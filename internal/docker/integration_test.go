//go:build integration
// +build integration

package docker_test

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/jeffs/geode/internal/docker"
)

func TestBuild(t *testing.T) {
	name := "geode/testdata-groovy"
	fmt.Fprintln(os.Stderr, "checking for existing image")
	if err := exec.Command("docker", "image", "inspect", name).Run(); err == nil {
		t.Fatalf("won't replace existing image %s", name)
	}

	fmt.Fprintln(os.Stderr, "building image")
	noCache := false
	if err := docker.Build("../../testdata/groovy", noCache); err != nil {
		t.Fatal(err)
	}

	// If Build failed to create the image, docker rmi will return non-zero.
	fmt.Fprintln(os.Stderr, "removing image")
	if err := exec.Command("docker", "rmi", "--no-prune", name).Run(); err != nil {
		t.Fatalf("can't remove test image %s: %v", name, err)
	}
}
