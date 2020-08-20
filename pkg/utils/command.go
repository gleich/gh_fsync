package utils

import (
	"os"
	"os/exec"
	"strings"

	"github.com/Matt-Gleich/logoru"
)

// Run a command and return the error
func RunCommand(command string, msg string) {
	cmdChunks := strings.Split(command, " ")
	cmd := exec.Command(cmdChunks[0], cmdChunks[1:]...)
	err := cmd.Run()
	if err != nil {
		logoru.Error(msg, cmd.Stderr)
		os.Exit(1)
	}
}
