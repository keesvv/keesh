package prompt

import (
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

// AutoComplete interfaces with `fzf` to provide autocompletion.
func AutoComplete(relPath string) (string, error) {
	fzf := exec.Command("fzf")

	files, err := ioutil.ReadDir(".")
	if err != nil {
		return "", err
	}

	var dirs []string
	for _, file := range files {
		entry := file.Name()

		if file.IsDir() {
			entry += "/"
		}

		dirs = append(dirs, entry)
	}

	r := strings.NewReader(strings.Join(dirs, "\n"))

	fzf.Stdin = r
	fzf.Stderr = os.Stderr

	completed, err := fzf.Output()
	return strings.TrimSpace(string(completed)), err
}
