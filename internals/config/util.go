package config

import (
	"os"
	"path"
)

// GetConfigRoot returns the root directory path where all
// configuration files reside.
func GetConfigRoot() string {
	confDir, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}

	return path.Join(confDir, "keesh")
}
