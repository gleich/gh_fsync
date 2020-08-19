package write

import (
	"os"

	"github.com/Matt-Gleich/gh_fsync/internal/source"
	"github.com/Matt-Gleich/logoru"
)

// Check to make sure all the files exist
func checkExistence(files map[string]string) {
	for fileName := range files {
		info, err := os.Stat(fileName)
		if os.IsNotExist(err) {
			logoru.Error("File", fileName, "doesn't exist")
			os.Exit(1)
		}
		if info.IsDir() {
			logoru.Error("File", fileName, "is a directory")
			os.Exit(1)
		}
	}
}

// Filter out all the files that haven't changed
func getChangedFiles(files map[string]source.File) map[string]string {
	changedFiles := map[string]string{}
	for fileName, file := range files {
		if file.Original == file.Updated {
			changedFiles[fileName] = file.Updated
		}
	}
	return changedFiles
}
