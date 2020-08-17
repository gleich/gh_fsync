package utils

import (
	"os"
	"testing"
)

// Check for a error in one line
func CheckError(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}

// Create a temporary environment for testing
func CreateTempEnv(t *testing.T, fName string) {
	f, err := os.Create(fName)
	CheckError(t, err)
	err = f.Close()
	CheckError(t, err)
}

// Remove the temporary environment for testing
func RemoveTempEnv(t *testing.T, fName string) {
	err := os.Remove(fName)
	CheckError(t, err)
}

// Chdir into project root
func ProjectRoot(t *testing.T, levels int) {
	var directories string
	for i := 1; i <= levels; i++ {
		directories = directories + "../"
	}
	err := os.Chdir(directories)
	CheckError(t, err)
}
