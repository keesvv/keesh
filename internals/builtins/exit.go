package builtins

import "github.com/keesvv/keesh/internals/util"

// Exit exits the shell.
func Exit(args ...string) error {
	util.Exit()
	return nil
}
