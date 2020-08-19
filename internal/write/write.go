package write

import (
	"github.com/Matt-Gleich/gh_fsync/internal/source"
)

// Filter out all the files that haven't changed
func getChangedFiles(files map[string]source.File) map[string]string {
	changedFiles := map[string]string{}
	for fileName, file := range files {
		if file.Current == file.Updated {
			changedFiles[fileName] = file.Updated
		}
	}
	return changedFiles
}
