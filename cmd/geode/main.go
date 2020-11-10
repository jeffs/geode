// Geode is a program to manage your personal config files: shell, editor, etc.
package main

import (
	"os"

	"github.com/jeffs/geode/internal/cli"
)

func main() {
	os.Exit(cli.Main(os.Args, os.Stdout, os.Stderr))
}
