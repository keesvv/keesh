package builtins

import "github.com/keesvv/keesh/internals/alias"

// Alias registers a new shell alias.
func Alias(args ...string) error {
	alias.RegisterAlias(args[1], args[0])
	return nil
}
