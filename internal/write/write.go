package write

import (
	"io/ioutil"
	"os"

	"github.com/Matt-Gleich/gh_fsync/internal/source"
	"github.com/Matt-Gleich/gh_fsync/pkg/utils"
	"github.com/Matt-Gleich/logoru"
)

// Prod writer
func WriteChanges(files map[string]source.File) {
	logoru.Info("Writing changes files")
	changedFiles := getChangedFiles(files)
	if len(changedFiles) == 0 {
		logoru.Info("No changes detected. Have a nice day! 👋")
		os.Exit(0)
	}
	rawWrite(changedFiles)
	logoru.Success("Wrote changes to files!")
}

// Write to all the files
func rawWrite(files map[string]string) {
	for fileName, fileContent := range files {
		err := ioutil.WriteFile(fileName, []byte(fileContent), 0700)
		utils.CheckErr("Failed to write to "+fileName, err)
	}
}

// Filter out all the files that haven't changed
func getChangedFiles(files map[string]source.File) map[string]string {
	changedFiles := map[string]string{}
	for fileName, file := range files {
		if file.Current != file.Updated {
			changedFiles[fileName] = file.Updated
		}
	}
	return changedFiles
}
