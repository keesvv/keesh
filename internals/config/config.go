package config

import (
	"bufio"
	"os"
)

// File represents a config file.
type File struct {
	Location string
	file     *os.File
	content  string
}

// Load opens the config file for reading and loads its contents
// into memory.
func (f *File) Load() error {
	file, err := os.Open(f.Location)
	if err != nil {
		return err
	}

	f.file = file
	s := bufio.NewScanner(file)

	// Load the file into memory
	for s.Scan() {
		f.content += s.Text() + "\n"
	}

	return nil
}

// Close closes the config file.
func (f *File) Close() {
	f.Close()
}
