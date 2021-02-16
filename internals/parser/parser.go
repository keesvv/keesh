package parser

import (
	"os"
	"os/exec"
	"strings"
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

func ParseCommand(input string) error {
	if input == "" {
		return nil
	}

	input = preprocessInput(input)
	cmdSplit := strings.Split(input, " ")

	cmd := exec.Command(cmdSplit[0], cmdSplit[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	cmd.Stdin = os.Stdin
	return cmd.Run()
}
