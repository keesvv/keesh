package util

import (
	"fmt"
	"os"
)

// Exit exits the shell.
func Exit() {
	fmt.Println("Goodbye!")
	os.Exit(0)
}
