package config

import (
	"testing"

	"github.com/Matt-Gleich/gh_fsync/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestRawRead(t *testing.T) {
	utils.ProjectRoot(t, 2)
	var instance Outline
	expected := Outline{GlobalReplace: map[string]string{"docker_username": "mattgleich", "github_username": "${{ file.USERNAME }}"}, Files: []FileOutline{{Name: "CONTRIBUTING.md", URL: "https://github.com/Matt-Gleich/go_template/blob/master/CONTRIBUTING.md", LocalReplace: map[string]string{"docker_username": "USERNAME"}}, {Name: "LICENSE.md", URL: "https://github.com/Matt-Gleich/go_template/blob/master/LICENSE.md", LocalReplace: map[string]string{"docker_username": "USERNAME"}}}}

	rawRead(&instance, "examples/config.yml")
	assert.Equal(t, expected, instance)
}

func TestCheckExistence(t *testing.T) {
	for i := range validLocations {
		utils.CreateTempEnv(t, validLocations[i])
		instance := checkExistence()
		assert.Equal(t, validLocations[i], instance)
		utils.RemoveTempEnv(t, validLocations[i])
	}
}
