package utils

import (
	"os"
	"testing"
)

// Create a temporary environment variable
func CreateTempEnv(t *testing.T, key string, value string) {
	err := os.Setenv(key, value)
	CheckTestingErr(t, err)
}

// Remove a temporary environment variable
func RemoveTempEnv(t *testing.T, key string) {
	err := os.Setenv(key, "")
	CheckTestingErr(t, err)
}

// Create a temporary environment for testing
func CreateTempFile(t *testing.T, fName string) {
	f, err := os.Create(fName)
	CheckTestingErr(t, err)
	err = f.Close()
	CheckTestingErr(t, err)
}

// Remove the temporary environment for testing
func RemoveTempFile(t *testing.T, fName string) {
	err := os.Remove(fName)
	CheckTestingErr(t, err)
}

// Chdir into project root
func ProjectRoot(t *testing.T, levels int) {
	var directories string
	for i := 1; i <= levels; i++ {
		directories = directories + "../"
	}
	err := os.Chdir(directories)
	CheckTestingErr(t, err)
}
