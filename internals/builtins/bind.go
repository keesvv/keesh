package builtins

import (
	"strings"

	"github.com/keesvv/keesh/internals/keybinds"
)

func Bind(args ...string) error {
	keybinds.RegisterBinding(args[0][0], strings.Join(args[1:], " "))
	return nil
}
