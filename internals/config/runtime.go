package config

import (
	"strings"
)

// RuntimeConfig represents a runtime configuration (rc) file.
type RuntimeConfig struct {
	*File
}

// GetCommands returns an ordered list of commands to execute.
func (r *RuntimeConfig) GetCommands() []string {
	return strings.Split(r.content, "\n")
}
