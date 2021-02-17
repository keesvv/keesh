package prompt

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
	"unicode"

	"github.com/fatih/color"
	"github.com/keesvv/keesh/internals/util"
	"github.com/pkg/term"
)

const (
	arrowRune rune = '❯'
	gitRune   rune = ''
)

var (
	segArrow string = color.HiGreenString(string(arrowRune))
)

// Prompt represents a shell prompt.
type Prompt struct {
	reader  *bufio.Reader
	term    *term.Term
	lastDir string
}

// Show displays the prompt.
func (p *Prompt) Show() (input string) {
	currentDir, wdErr := os.Getwd()
	if wdErr != nil {
		panic(wdErr)
	}

	homeDir, hdErr := os.UserHomeDir()
	if hdErr != nil {
		panic(hdErr)
	}

	segDir := currentDir
	segDir = strings.Replace(segDir, homeDir, color.HiCyanString("~"), 1)
	segDir = strings.ReplaceAll(segDir, "/", color.HiBlackString("/"))

	icons := make([]string, 0, 1)

	if _, err := os.Stat(path.Join(currentDir, ".git")); err == nil {
		icons = append(icons, color.HiRedString(string(gitRune)))
	}

	segIcons := strings.Join(icons, " ")

	// Print all segments
	fmt.Printf("%s %s\n%s ", segDir, segIcons, segArrow)

	// Enter cbreak mode
	p.term.SetCbreak()

	// TODO: refactor
	for {
		b, err := p.reader.ReadByte()
		if err != nil {
			panic(err)
		}

		// Tab
		if b == 9 {
			cmpl, err := AutoComplete("")

			if err != nil {
				code := err.(*exec.ExitError).ExitCode()

				// Aborted or completion not found
				if code == 1 || code == 130 {
					continue
				}

				fmt.Printf("\n%s\n", err)
				break
			}

			fmt.Print(cmpl)
			input += cmpl
			continue
		}

		// CTRL + D
		if b == 4 {
			fmt.Println()
			util.Exit()
		}

		// Newline
		if b == 10 {
			fmt.Println()
			break
		}

		// Backspace
		if b == 127 {
			if len(input) > 0 {
				fmt.Print("\b \b")
				input = input[:len(input)-1]
			}
			continue
		}

		if !unicode.IsPrint(rune(b)) {
			continue
		}

		input += string(b)
		fmt.Print(string(b))

		if input == ".." {
			input = "cd .."
			p.lastDir = currentDir
			fmt.Println()
			break
		}

		if input == ",," {
			input = fmt.Sprintf("cd %s", p.lastDir)
			fmt.Println()
			break
		}
	}

	p.term.Restore()
	return strings.TrimSpace(input)
}

// NewPrompt instantiates a Prompt.
func NewPrompt() *Prompt {
	term, err := term.Open(os.Stdout.Name())

	if err != nil {
		panic(err)
	}

	return &Prompt{
		reader: bufio.NewReader(os.Stdin),
		term:   term,
	}
}
