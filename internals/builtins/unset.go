package builtins

import "os"

// Unset unsets an environment variable.
func Unset(args ...string) error {
	return os.Unsetenv(args[0])
}
