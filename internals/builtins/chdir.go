package builtins

import "os"

// Chdir changes the current working directory.
func Chdir(args ...string) error {
	return os.Chdir(args[0])
}
