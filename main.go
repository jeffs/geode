package main

import (
	"fmt"
	"os"
)

var usage string = `Geode is a tool for managing config files.

Usage:

	geode <command> [arguments]

The commands are:

	help		Print this message
	update		Download and install the latest version of Geode
`

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "geode: expected command")
		fmt.Fprintln(os.Stderr, "Run 'geode help' for usage.")
	}
}
