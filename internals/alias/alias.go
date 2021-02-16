package alias

var aliasMap = map[string]string{}

// IsAlias returns whether the given command is an alias or not.
func IsAlias(name string) bool {
	_, ok := aliasMap[name]
	return ok
}

// ExpandAlias returns the command associated with the alias.
func ExpandAlias(name string) string {
	return aliasMap[name]
}

// RegisterAlias registers a new alias.
func RegisterAlias(cmd string, alias string) {
	aliasMap[alias] = cmd
}
