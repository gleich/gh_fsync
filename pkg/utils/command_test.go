package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunCommand(t *testing.T) {
	const failMsg = "Failed :("
	RunCommand("ls", []string{}, failMsg)
	assert.Equal(t, "Hello\n", RunCommand("echo", []string{"Hello"}, failMsg))
}
