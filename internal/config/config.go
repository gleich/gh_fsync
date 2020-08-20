package config

import (
	"io/ioutil"
	"os"

	"github.com/Matt-Gleich/gh_fsync/pkg/utils"
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
	Source       string            `yaml:"source"`
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
	utils.CheckErr("Failed to read config file", err)
	err = yaml.Unmarshal(content, c)
	utils.CheckErr("Failed to unmarshal yaml config", err)
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
