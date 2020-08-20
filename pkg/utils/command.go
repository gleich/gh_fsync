package utils

import (
	"os/exec"
	"strings"
)

// Run a command and return the error
func RunCommand(command string, msg string) {
	cmdChunks := strings.Split(command, " ")
	cmd := exec.Command(cmdChunks[0], cmdChunks[1:]...)
	err := cmd.Run()
	CheckErr(msg, err)
}
