package write

import (
	"testing"

	"github.com/Matt-Gleich/gh_fsync/internal/source"
	"github.com/Matt-Gleich/gh_fsync/pkg/utils"
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
	utils.CreateTempEnv(t, fName1)
	utils.CreateTempEnv(t, fName2)
	defer utils.RemoveTempEnv(t, fName1)
	defer utils.RemoveTempEnv(t, fName2)

	rawWrite(map[string]string{
		"test1.txt": test1Contents,
		"test2.txt": test2Contents,
	})
	assert.Equal(t, test1Contents, utils.SafeFileRead(fName1))
	assert.Equal(t, test2Contents, utils.SafeFileRead(fName2))
}

func TestGetChangedFiles(t *testing.T) {
	instance1 := getChangedFiles(map[string]source.File{
		"file1.txt": {
			Current: "same",
			Updated: "same",
		},
		"file2.txt": {
			Current: "xerox",
			Updated: "xerox",
		},
	})
	assert.Equal(t, map[string]string{
		"file1.txt": "same",
		"file2.txt": "xerox",
	}, instance1)

	instance2 := getChangedFiles(map[string]source.File{
		"file1.txt": {
			Current: "same",
			Updated: "same",
		},
		"file2.txt": {
			Current: "different",
			Updated: "xerox",
		},
	})
	assert.Equal(t, map[string]string{
		"file1.txt": "same",
	}, instance2)

	instance3 := getChangedFiles(map[string]source.File{
		"file1.txt": {
			Current: "same",
			Updated: "different",
		},
	})
	assert.Equal(t, map[string]string{}, instance3)
}
