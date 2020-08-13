package config

import (
	"os"

	"github.com/Matt-Gleich/logoru"
)

var validLocations = []string{".github/fsync.yml", ".github/fsync.yaml"}

// Check to make sure the config file exists. Returns the path
func checkExistence() string {
	var foundPath string
	for _, location := range validLocations {
		info, err := os.Stat(location)
		if !os.IsNotExist(err) && !info.IsDir() {
			foundPath = location
		}
	}
	if foundPath == "" {
		logoru.Error("Configuration file not found")
		os.Exit(1)
	}
	return foundPath
}
