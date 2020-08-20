package write

import (
	"io/ioutil"

	"github.com/Matt-Gleich/gh_fsync/pkg/utils"
	"github.com/Matt-Gleich/logoru"
)

// Prod writer
func WriteChanges(changes map[string]string) {
	logoru.Info("Writing changes files")
	rawWrite(changes)
	logoru.Success("Wrote changes to files!")
}

// Write to all the files
func rawWrite(files map[string]string) {
	for fileName, fileContent := range files {
		err := ioutil.WriteFile(fileName, []byte(fileContent), 0700)
		utils.CheckErr("Failed to write to "+fileName, err)
	}
}
