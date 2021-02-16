package main

import (
	"errors"
	"fmt"
	"os/exec"

	"github.com/keesvv/keesh/internals/config"
	"github.com/keesvv/keesh/internals/parser"
	"github.com/keesvv/keesh/internals/prompt"
)

func main() {
	// Load the runtime config (rc)
	config.Runtime.Load()

	for _, cmd := range config.Runtime.GetCommands() {
		parser.ParseCommand(cmd)
	}

	p := prompt.NewPrompt()

	for {
		input := p.Show()

		if err := parser.ParseCommand(input); errors.Is(err, exec.ErrNotFound) {
			fmt.Println("command not found")
		} else if err != nil {
			fmt.Println(err)
		}

		fmt.Println()
	}
}
