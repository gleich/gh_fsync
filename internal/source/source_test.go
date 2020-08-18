package source

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetContent(t *testing.T) {
	instance1 := getContent("https://raw.githubusercontent.com/Matt-Gleich/jsx/master/public/robots.txt")
	assert.Equal(t, instance1, "# https://www.robotstxt.org/robotstxt.html\nUser-agent: *\nDisallow:\n")
	instance2 := getContent("https://raw.githubusercontent.com/Matt-Gleich/CongressPresenation/master/.metadata")
	assert.Equal(t, instance2, "# This file tracks properties of this Flutter project.\n# Used by Flutter tool to assess capabilities and perform upgrades etc.\n#\n# This file should be version controlled and should not be manually edited.\n\nversion:\n  revision: 8735ab1e35346ae20b6c347d259b07b1589756a5\n  channel: master\n\nproject_type: app\n")
}

func TestRawURL(t *testing.T) {
	instance1 := rawURL("https://github.com/Matt-Gleich/go_template/blob/master/CONTRIBUTING.md")
	assert.Equal(t, instance1, "https://raw.githubusercontent.com/Matt-Gleich/go_template/master/CONTRIBUTING.md")
	instance2 := rawURL("https://github.com/Matt-Gleich/Dot-Files/blob/master/LICENSE.md")
	assert.Equal(t, instance2, "https://raw.githubusercontent.com/Matt-Gleich/Dot-Files/master/LICENSE.md")
}
