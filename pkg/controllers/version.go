package controllers

import (
	"io"
	"os"
	"strings"
)

type Versioner interface {
	Get() (string, error)
}

type Version struct{}

func NewVersion() *Version {
	return &Version{}
}

func (v Version) Get() (string, error) {
	var (
		version string
		err     error
	)

	if version, err = getVersion("VERSION"); err != nil {
		return "", err
	}

	return version, nil
}

// getVersion reads a file, skips empty lines, and returns the first non-empty line as a string.
func getVersion(filePath string) (string, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Read the contents of the file using io.ReadAll
	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	// Convert the byte slice to a string
	version := string(content)

	// Split the content into lines
	lines := strings.Split(version, "\n")

	// Find the first non-empty line
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			return line, nil
		}
	}

	// If all lines are empty, return an empty string
	return "", nil
}
