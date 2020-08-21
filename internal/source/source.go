package source

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/Matt-Gleich/gh_fsync/internal/config"
	"github.com/Matt-Gleich/gh_fsync/pkg/utils"
	"github.com/Matt-Gleich/logoru"
)

// A file entry
type File struct {
	Current string
	Updated string
}

// Get the content for all the files
func GetFromSource(configuration config.Outline) map[string]File {
	logoru.Info("📡 Getting files from source")
	files := map[string]File{}
	for _, file := range configuration.Files {
		currentFile := utils.SafeFileRead(file.Path)
		sourceURL := file.Source
		if sourceURL == "" {
			logoru.Error("Source URL not defined for " + file.Path)
			os.Exit(1)
		}
		sourceFile := getSourceContent(rawURL(sourceURL))
		updateFile := replace(sourceFile, configuration.GlobalReplace, file.LocalReplace)
		files[file.Path] = File{
			Current: currentFile,
			Updated: updateFile,
		}
	}
	logoru.Success("✅ Got files from source")
	return files
}

// Replace everything defined for that specific file (localReplace) first and then any global replacements (globalReplace)
func replace(raw string, globalReplace []config.ReplaceOutline, localReplace []config.ReplaceOutline) string {
	// Local replace
	for _, lReplace := range localReplace {
		raw = strings.ReplaceAll(raw, lReplace.Before, lReplace.After)
	}
	// Global replace
	for _, gReplace := range globalReplace {
		var alreadyUsed bool
		for _, lReplace := range localReplace {
			if gReplace.Before == lReplace.Before {
				alreadyUsed = true
				break
			}
		}
		if !alreadyUsed {
			raw = strings.ReplaceAll(raw, gReplace.Before, gReplace.After)
		}
	}
	return raw
}

// Get the content of a file from the raw url
// Returns the contents of that file
func getSourceContent(url string) string {
	// Making get request
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		logoru.Error("Failed to get contents of", url, ";", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	// Reading the actual response
	data, err := ioutil.ReadAll(resp.Body)
	utils.CheckErr("Failed to parse get request for file "+url, err)
	return string(data)
}

// Replace the github view of a url to a raw text view
// From: https://github.com/Matt-Gleich/go_template/blob/master/CONTRIBUTING.md
// To: https://raw.githubusercontent.com/Matt-Gleich/go_template/master/CONTRIBUTING.md
func rawURL(url string) string {
	chunks := strings.Split(url, "/")
	blobRemoved := strings.Join(append(chunks[:5], chunks[6:]...), "/")
	return strings.Replace(blobRemoved, "https://github.com", "https://raw.githubusercontent.com", 1)
}
