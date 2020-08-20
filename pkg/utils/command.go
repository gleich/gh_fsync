package utils

import (
	"bytes"
	"os/exec"
)

// Run a command and return the error
func RunCommand(name string, args []string, msg string) string {
	cmd := exec.Command(name, args...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	CheckErr(msg+":"+stderr.String(), err)
	return out.String()
}
