package utils

import (
	"io/ioutil"
	"os"

	"github.com/Matt-Gleich/logoru"
)

// Read a file and check for an error beforehand
func SafeFileRead(fileName string) string {
	// Checking to see if the file exists
	info, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		logoru.Error("File", fileName, "doesn't exist")
		os.Exit(1)
	}
	if info.IsDir() {
		logoru.Error("File", fileName, "is a directory")
		os.Exit(1)
	}
	// Reading from the actual file
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		logoru.Error("Failed to read from file:", fileName, err)
		os.Exit(1)
	}
	return string(content)
}
