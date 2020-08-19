package write

import (
	"testing"

	"github.com/Matt-Gleich/gh_fsync/internal/source"
	"github.com/Matt-Gleich/gh_fsync/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestCheckExistence(t *testing.T) {
	utils.ProjectRoot(t, 2)
	checkExistence(map[string]string{
		"LICENSE.md": "",
	})
	checkExistence(map[string]string{
		"README.md":       "",
		"CONTRIBUTING.md": "",
	})
}

func TestGetChangedFiles(t *testing.T) {
	instance1 := getChangedFiles(map[string]source.File{
		"file1.txt": {
			Original: "same",
			Updated:  "same",
		},
		"file2.txt": {
			Original: "xerox",
			Updated:  "xerox",
		},
	})
	assert.Equal(t, map[string]string{
		"file1.txt": "same",
		"file2.txt": "xerox",
	}, instance1)

	instance2 := getChangedFiles(map[string]source.File{
		"file1.txt": {
			Original: "same",
			Updated:  "same",
		},
		"file2.txt": {
			Original: "different",
			Updated:  "xerox",
		},
	})
	assert.Equal(t, map[string]string{
		"file1.txt": "same",
	}, instance2)

	instance3 := getChangedFiles(map[string]source.File{
		"file1.txt": {
			Original: "same",
			Updated:  "different",
		},
	})
	assert.Equal(t, map[string]string{}, instance3)
}
