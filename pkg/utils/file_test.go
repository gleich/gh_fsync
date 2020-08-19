package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSourceContent(t *testing.T) {
	ProjectRoot(t, 2)
	instance := SafeFileRead(".github/workflows/contributors.yml")
	assert.Equal(t, "name: contributors\n\non:\n  push:\n    branches:\n      - master\n\njobs:\n  contributor_list:\n    runs-on: ubuntu-latest\n    steps:\n      - uses: actions/checkout@master\n      - uses: cjdenio/contributor_list@master\n        with:\n          commit_message: 📝 Update contributors list\n          max_contributors: 10\n", instance)
}
