// Geode is a program to manage your personal config files: shell, editor, etc.
package main

import (
	"os"

	"github.com/jeffs/geode/internal/command"
)

func main() {
	os.Exit(command.Main(os.Args))
}
