package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSourceContent(t *testing.T) {
	ProjectRoot(t, 2)
	instance := SafeFileRead("action.yml")
	assert.Equal(
		t,
		"name: 'gh_fsync'\ndescription: '🔄 GitHub action to sync files across repos in GitHub'\nbranding:\n  icon: 'folder-plus'\n  color: 'green'\nruns:\n  using: 'docker'\n  image: 'Dockerfile'\n",
		instance,
	)
}
