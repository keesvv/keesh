package builtins

var builtinsMap = map[string]func(args ...string) error{
	"cd":    Chdir,
	"exit":  Exit,
	"alias": Alias,
	"set":   Set,
	"unset": Unset,
}

// IsBuiltin returns whether the given command is a shell built-in or not.
func IsBuiltin(name string) bool {
	_, ok := builtinsMap[name]
	return ok
}

// Execute executes a shell built-in.
func Execute(name string, args []string) error {
	return builtinsMap[name](args...)
}
