package parser

import (
	"os"
	"os/exec"
	"strings"

	"github.com/keesvv/keesh/internals/builtins"
)

func preprocessInput(input string) (output string) {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	output = input
	output = strings.ReplaceAll(output, "~", home)
	output = os.ExpandEnv(output)
	return
}

// ParseCommand parses and executes the given input.
func ParseCommand(input string) error {
	if input == "" {
		return nil
	}

	input = preprocessInput(input)
	cmdSplit := strings.Split(input, " ")

	name := cmdSplit[0]
	args := cmdSplit[1:]

	if builtins.IsBuiltin(name) {
		return builtins.Execute(name, args)
	}

	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	cmd.Stdin = os.Stdin
	return cmd.Run()
}
