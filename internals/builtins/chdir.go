package builtins

import "os"

// Chdir changes the current working directory.
func Chdir(args ...string) error {
	if len(args) == 0 {
		home, _ := os.UserHomeDir()
		return os.Chdir(home)
	}

	return os.Chdir(args[0])
}
