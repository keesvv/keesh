package parser

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/keesvv/keesh/internals/alias"
	"github.com/keesvv/keesh/internals/builtins"
)

func preprocessInput(input string) (output string) {
	if input == "" {
		return
	}

	// Ignore comments
	if input[0] == '#' {
		return
	}

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
func ParseCommand(input string) {
	input = preprocessInput(input)

	if input == "" {
		return
	}

	cmdSplit := strings.Split(input, " ")

	name := cmdSplit[0]
	args := cmdSplit[1:]

	if alias.IsAlias(name) {
		expandedArgs := strings.Split(alias.ExpandAlias(name), " ")
		name = expandedArgs[0]
		args = append(expandedArgs[1:], args...)
	}

	if builtins.IsBuiltin(name) {
		builtins.Execute(name, args)
		return
	}

	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if errors.Is(err, exec.ErrNotFound) {
		fmt.Printf("command '%s' not found\n", name)
	} else if err != nil {
		fmt.Println(err)
	}
}
