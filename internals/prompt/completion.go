package prompt

import (
	"os"
	"os/exec"
	"strings"
)

// AutoComplete interfaces with `fzf` to provide autocompletion.
func AutoComplete(relPath string) (string, error) {
	ls := exec.Command("ls", relPath)
	fzf := exec.Command("fzf")

	pipe, err := ls.StdinPipe()
	defer pipe.Close()

	if err != nil {
		return "", err
	}

	fzf.Stdin = os.Stdin
	fzf.Stderr = os.Stderr

	ls.Run()

	completed, err := fzf.Output()
	return strings.TrimSpace(string(completed)), err
}
