package parser

import (
	"os"
	"strings"
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
