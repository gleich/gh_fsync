package config

import (
	"io/ioutil"
	"os"

	"github.com/Matt-Gleich/logoru"
	"gopkg.in/yaml.v3"
)

// Outline for the configuration file
type Outline struct {
	GlobalReplace map[string]string `yaml:"variables"`
	Files         []FileOutline     `yaml:"files"`
}

// Outline for a file in the config
type FileOutline struct {
	Name         string            `yaml:"name"`
	URL          string            `yaml:"url"`
	LocalReplace map[string]string `yaml:"variables"`
}

var validLocations = []string{
	".github/fsync.yml", ".github/fsync.yaml", ".fsync.yml", ".fsync.yaml", "fsync.yml", "fsync.yaml",
}

// Check existence and read the configuration fiAle
func Read() Outline {
	logoru.Info("Reading from configuration file")
	path := checkExistence()
	var configuration Outline
	rawRead(&configuration, path)
	logoru.Success("Read from configuration file")
	return configuration
}

// Read from the config file
func rawRead(c *Outline, path string) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		logoru.Error("Failed to read config file;", err)
		os.Exit(1)
	}
	err = yaml.Unmarshal(content, c)
	if err != nil {
		logoru.Error("Failed to unmarshal yaml config;", err)
		os.Exit(1)
	}
}

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
