package utils

import (
	"fmt"
	"os/exec"
	"strings"
)

// Run a command and return the error
func RunCommand(command string, msg string) string {
	cmdChunks := strings.Split(command, " ")
	out, err := exec.Command(cmdChunks[0], cmdChunks[1:]...).Output()
	strOut := string(out)
	fmt.Print(strOut)
	CheckErr(msg, err)
	return strOut
}
