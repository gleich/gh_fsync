package utils

import "testing"

func TestRunCommand(t *testing.T) {
	const failMsg = "Failed :("
	RunCommand("ls -la", failMsg)
	RunCommand("ls", failMsg)
}
