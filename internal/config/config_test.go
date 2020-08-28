package config

import (
	"testing"

	"github.com/Matt-Gleich/gh_fsync/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestRawRead(t *testing.T) {
	utils.ProjectRoot(t, 2)
	var instance Outline
	rawRead(&instance, "examples/config.yml")
	assert.Equal(t, Outline{CommitMessage: "🔄 Update via sync", GlobalReplace: []ReplaceOutline{{Before: "project_name", After: "gh_fsync"}}, Files: []FileOutline{{Path: "CONTRIBUTING.md", Source: "https://github.com/Matt-Gleich/go_template/blob/master/CONTRIBUTING.md", LocalReplace: []ReplaceOutline{{Before: "project_name", After: "gh_fsync2"}}, IgnoreGlobalReplace: false}, {Path: "LICENSE.md", Source: "https://github.com/Matt-Gleich/go_template/blob/master/LICENSE.md", LocalReplace: []ReplaceOutline{{Before: "author_name", After: "Matthew Gleich"}}, IgnoreGlobalReplace: true}}}, instance)
}

func TestCheckExistence(t *testing.T) {
	for i := range validLocations {
		utils.CreateTempFile(t, validLocations[i])
		instance := checkExistence()
		assert.Equal(t, validLocations[i], instance)
		utils.RemoveTempFile(t, validLocations[i])
	}
}
