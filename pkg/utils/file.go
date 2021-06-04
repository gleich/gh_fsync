package utils

import (
	"io/ioutil"
	"os"

	"github.com/gleich/logoru"
)

// Read a file and check for an error beforehand
func SafeFileRead(fileName string) string {
	// Checking to see if the file exists
	info, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return ""
	}
	if info.IsDir() {
		logoru.Error("File", fileName, "is a directory")
		os.Exit(1)
	}
	// Reading from the actual file
	content, err := ioutil.ReadFile(fileName)
	CheckErr("Failed to read from file;", err)
	return string(content)
}
