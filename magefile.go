//+build mage

package main

import "github.com/magefile/mage/sh"

// Builds the binary
func Build() {
	sh.Run("go", "build", "-o", "build/keesh", "--ldflags", "-s -w", "cmd/keesh/keesh.go")
}

// Cleans up build files
func Clean() {
	sh.Rm("build")
}
