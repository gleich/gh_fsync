package write

import (
	"os"
	"testing"

	"github.com/gleich/gh_fsync/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestRawWrite(t *testing.T) {
	utils.ProjectRoot(t, 2)
	const (
		fName1        = "test1.txt"
		fName2        = "test2.txt"
		test1Contents = "HEEEEEELLLLLLOOOOO"
		test2Contents = "WWWWOOOOORRRRLLLDD"
	)
	utils.CreateTempFile(t, fName1)
	utils.CreateTempFile(t, fName2)
	defer utils.RemoveTempFile(t, fName1)
	defer utils.RemoveTempFile(t, fName2)

	rawWrite(map[string]string{
		"test1.txt": test1Contents,
		"test2.txt": test2Contents,
	})
	assert.Equal(t, test1Contents, utils.SafeFileRead(fName1))
	assert.Equal(t, test2Contents, utils.SafeFileRead(fName2))
}

func TestCreateParentFolder(t *testing.T) {
	const folderName = "test-folder"
	completePath := "./" + folderName + "/text.txt"

	// Folder doesn't exist
	createParentFolder(completePath)
	assert.True(t, fExists(folderName))

	// Folder does exist
	createParentFolder(completePath)
	assert.True(t, fExists(folderName))
	assert.False(t, fExists(completePath))

	err := os.RemoveAll(folderName)
	utils.CheckTestingErr(t, err)
}

// Check if a file or folder exists
func fExists(folderName string) bool {
	_, err := os.Stat(folderName)
	return !os.IsNotExist(err)
}
