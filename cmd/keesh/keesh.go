package main

import (
	"fmt"

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
		parser.ParseCommand(input)
		fmt.Println()
	}
}
