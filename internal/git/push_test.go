package git

import (
	"fmt"
	"testing"

	"github.com/Matt-Gleich/gh_fsync/pkg/utils"
	"github.com/tj/assert"
)

func TestGetOriginURL(t *testing.T) {
	instance1 := getOriginURL()
	assert.Equal(t, "https://github.com/Matt-Gleich/gh_fsync.git", instance1)

	// Setting up environment
	const (
		patEnvName = "INPUT_PERSONAL_ACCESS_TOKEN"
		ghEnvName  = "INPUT_GITHUB_USERNAME"
		fakePAT    = "g7f8d09ds7g0897g0f8d"
		username   = "Matt-Gleich"
	)
	utils.CreateTempEnv(t, patEnvName, fakePAT) // Fake PAT
	utils.CreateTempEnv(t, ghEnvName, username)
	defer utils.RemoveTempEnv(t, patEnvName)
	defer utils.RemoveTempEnv(t, ghEnvName)

	instance2 := getOriginURL()
	assert.Equal(t, fmt.Sprintf("https://%v:%v@%v",
		username,
		fakePAT,
		"github.com/Matt-Gleich/gh_fsync.git",
	), instance2)
}
