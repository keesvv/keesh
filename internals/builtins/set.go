package builtins

import "os"

// Set sets the value of an environment variable.
func Set(args ...string) error {
	return os.Setenv(args[0], args[1])
}
