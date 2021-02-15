package parser

import (
	"os"
	"os/exec"
	"strings"
)

func ParseCommand(input string) error {
	if input == "" {
		return nil
	}

	cmdSplit := strings.Split(input, " ")

	cmd := exec.Command(cmdSplit[0], cmdSplit[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	return cmd.Run()
}
