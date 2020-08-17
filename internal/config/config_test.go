package config

import (
	"testing"

	"github.com/Matt-Gleich/gh_fsync/pkg/utils"
	"github.com/stretchr/testify/assert"
)

const upLevels = 2

func TestCheckExistence(t *testing.T) {
	utils.ProjectRoot(t, upLevels)
	for i := range validLocations {
		utils.CreateTempEnv(t, validLocations[i])
		instance := checkExistence()
		assert.Equal(t, instance, validLocations[i])
		utils.RemoveTempEnv(t, validLocations[i])
	}
}

func TestRawRead(t *testing.T) {
	var instance Outline
	expected := &Outline{Variables: map[string]string{"docker_username": "mattgleich", "github_username": "${{ file.USERNAME }}"}, Files: []interface{}{map[string]interface{}{"CONTRIBUTING.md": "CONTRIBUTING.md", "variables": map[string]interface{}{"docker_username": "USERNAME"}}}}

	rawRead(&instance, "examples/config.yml")
	assert.Equal(t, instance, *expected)
}
