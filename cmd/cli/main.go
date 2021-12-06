package main

import (
	"os"

	"github.com/quadrosh/user-list/cli"
)

// CommandLine manages CLI
type CommandLine struct {
}

func main() {
	defer os.Exit(0)
	cmd := cli.CommandLine{}
	cmd.Run()

}
