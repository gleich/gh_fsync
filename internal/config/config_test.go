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
	assert.Equal(t, Outline{GlobalReplace: []ReplaceOutline{{Before: "project_name", After: "gh_fsync"}}, Files: []FileOutline{{Name: "CONTRIBUTING.md", Source: "https://github.com/Matt-Gleich/go_template/blob/master/CONTRIBUTING.md", LocalReplace: []ReplaceOutline{{Before: "project_name", After: "testing"}}}, {Name: "LICENSE.md", Source: "https://github.com/Matt-Gleich/go_template/blob/master/LICENSE.md", LocalReplace: []ReplaceOutline{{Before: "docker_username", After: "TESTING TESTING"}}}}}, instance)
}

func TestCheckExistence(t *testing.T) {
	for i := range validLocations {
		utils.CreateTempEnv(t, validLocations[i])
		instance := checkExistence()
		assert.Equal(t, validLocations[i], instance)
		utils.RemoveTempEnv(t, validLocations[i])
	}
}
