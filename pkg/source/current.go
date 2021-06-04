package source

import (
	"os"

	"github.com/gleich/logoru"
)

// Prod version of rawGetChanges
func GetChanges(files map[string]File) map[string]string {
	logoru.Info("👀 Checking for changed files")
	changes := rawGetChanges(files)
	if len(changes) == 0 {
		logoru.Info("No changes detected! Have a good day 👋")
		os.Exit(0)
	}
	logoru.Success("✅ Got changed files")
	return changes
}

// Filter out all the files that haven't changed
func rawGetChanges(files map[string]File) map[string]string {
	changedFiles := map[string]string{}
	for fileName, file := range files {
		if file.Current != file.Updated {
			changedFiles[fileName] = file.Updated
		}
	}
	return changedFiles
}
