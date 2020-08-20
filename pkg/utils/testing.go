package utils

import (
	"os"
	"testing"
)

// Create a temporary environment for testing
func CreateTempEnv(t *testing.T, fName string) {
	f, err := os.Create(fName)
	CheckTestingErr(t, err)
	err = f.Close()
	CheckTestingErr(t, err)
}

// Remove the temporary environment for testing
func RemoveTempEnv(t *testing.T, fName string) {
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
