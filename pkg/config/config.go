package config

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/Matt-Gleich/gh_fsync/pkg/utils"
	"github.com/Matt-Gleich/logoru"
	"gopkg.in/yaml.v3"
)

// Outline for the configuration file
type Outline struct {
	CommitMessage string           `yaml:"commit_message"`
	GlobalReplace []ReplaceOutline `yaml:"replace"`
	Files         []FileOutline    `yaml:"files"`
}

// Outline for a file in the config
type FileOutline struct {
	Path                string           `yaml:"path"`
	Source              string           `yaml:"source"`
	LocalReplace        []ReplaceOutline `yaml:"replace"`
	IgnoreGlobalReplace bool             `yaml:"ignore_global_replace"`
}

// Outline for a replacement
type ReplaceOutline struct {
	Before string `yaml:"before"`
	After  string `yaml:"after"`
}

// Valid locations for where the config file can live
var validLocations = []string{
	".github/fsync.yml", ".github/fsync.yaml",
	".fsync.yml", ".fsync.yaml",
	"fsync.yml", "fsync.yaml",
}

// Check existence and read the configuration fiAle
func Read() Outline {
	logoru.Info("⚙️ Reading from configuration file")
	path := checkExistence()
	configuration := rawRead(path)
	logoru.Success("✅ Read from configuration file")
	return configuration
}

// Read from the config file
func rawRead(path string) Outline {
	content, err := ioutil.ReadFile(path)
	utils.CheckErr("Failed to read config file", err)
	var configuration Outline
	err = yaml.Unmarshal(content, &configuration)
	utils.CheckErr("Failed to unmarshal yaml config", err)

	// Removing files that come from the current repo
	url := strings.TrimSuffix(
		utils.RunCommand(
			"git",
			[]string{"config", "--get", "remote.origin.url"},
			"Failed to get remote URL",
		),
		".git",
	)
	cleaned_files := []FileOutline{}
	for _, file := range configuration.Files {
		if !strings.HasPrefix(file.Source, url) {
			cleaned_files = append(cleaned_files, file)
		}
	}
	configuration.Files = cleaned_files

	return configuration
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
		logoru.Critical("Configuration file not found")
		os.Exit(1)
	}
	return foundPath
}
