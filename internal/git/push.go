package git

import (
	"os"
	"os/exec"

	"github.com/Matt-Gleich/logoru"
)

// Push changes
func Push() {
	logoru.Info("Pushing changes")
	err := exec.Command("git", "push").Run()
	if err != nil {
		logoru.Error("Failed to push changes", err)
		os.Exit(1)
	}
	logoru.Success("Pushed changes to master! Have a good day 👋")
}
