package prompt

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const promptRune rune = '❯'

// Prompt represents a shell prompt.
type Prompt struct {
	reader *bufio.Reader
}

// Show displays the prompt.
func (p *Prompt) Show() (input string) {
	currentDir, wdErr := os.Getwd()
	if wdErr != nil {
		panic(wdErr)
	}

	fmt.Printf("%s\n%s ", currentDir, string(promptRune))

	input, err := p.reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(input)
}

// NewPrompt instantiates a Prompt.
func NewPrompt() *Prompt {
	return &Prompt{
		reader: bufio.NewReader(os.Stdin),
	}
}
