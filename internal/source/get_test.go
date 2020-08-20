package source

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetChangedFiles(t *testing.T) {
	instance1 := GetChanges(map[string]File{
		"file1.txt": {
			Current: "same",
			Updated: "same",
		},
		"file2.txt": {
			Current: "xerox",
			Updated: "xerox",
		},
	})
	assert.Equal(t, map[string]string{}, instance1)

	instance2 := GetChanges(map[string]File{
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
		"file2.txt": "xerox",
	}, instance2)

	instance3 := GetChanges(map[string]File{
		"file1.txt": {
			Current: "same",
			Updated: "different",
		},
	})
	assert.Equal(t, map[string]string{
		"file1.txt": "different",
	}, instance3)
}
