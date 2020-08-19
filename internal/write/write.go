package write

import (
	"io/ioutil"
	"os"

	"github.com/Matt-Gleich/gh_fsync/internal/source"
	"github.com/Matt-Gleich/logoru"
)

// Write to all the files
func rawWrite(files map[string]string) {
	for fileName, fileContent := range files {
		err := ioutil.WriteFile(fileName, []byte(fileContent), 0700)
		if err != nil {
			logoru.Error("Failed to write to file:", fileName, err)
			os.Exit(1)
		}
	}
}

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
