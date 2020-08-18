package source

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRawURL(t *testing.T) {
	instance1 := rawURL("https://github.com/Matt-Gleich/go_template/blob/master/CONTRIBUTING.md")
	assert.Equal(t, instance1, "https://raw.githubusercontent.com/Matt-Gleich/go_template/master/CONTRIBUTING.md")
	instance2 := rawURL("https://github.com/Matt-Gleich/Dot-Files/blob/master/LICENSE.md")
	assert.Equal(t, instance2, "https://raw.githubusercontent.com/Matt-Gleich/Dot-Files/master/LICENSE.md")
}
