package write

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/Matt-Gleich/gh_fsync/pkg/utils"
	"github.com/Matt-Gleich/logoru"
)

// Prod version for rawWrite
func WriteChanges(changes map[string]string) {
	logoru.Info("✍️ Writing changes to files")
	rawWrite(changes)
	logoru.Success("✅ Wrote changes to files")
}

// Write to all the files
func rawWrite(files map[string]string) {
	for fileName, fileContent := range files {
		createParentFolder(fileName)
		err := ioutil.WriteFile(fileName, []byte(fileContent), 0655)
		utils.CheckErr("Failed to write to "+fileName, err)
	}
}

// Create the parent folder for a file if it doesn't exist
func createParentFolder(fName string) {
	parentFolder := filepath.Dir(fName)
	err := os.MkdirAll(parentFolder, 0700)
	utils.CheckErr("Failed to create parent folder for file", err)
}
