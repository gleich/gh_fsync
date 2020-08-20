package source

import (
	"github.com/Matt-Gleich/logoru"
)

func GetChanges(files map[string]File) map[string]string {
	logoru.Info("Get changed files")
	changes := rawGetChanges(files)
	logoru.Success("Got changed files")
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
