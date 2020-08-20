package utils

import (
	"os"
	"testing"

	"github.com/Matt-Gleich/logoru"
)

// Check an error in one line
func CheckErr(msg string, err error) {
	if err != nil {
		logoru.Error(msg, ";", err)
		os.Exit(1)
	}
}

// Check for a error in one line
func CheckTestingErr(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}
