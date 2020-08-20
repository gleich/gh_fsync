package changes

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/Matt-Gleich/logoru"
)

func Commit() {
	logoru.Info("Commiting changes")
	err := exec.Command("git", "add", ".").Run()
	if err != nil {
		logoru.Error("Failed to stage changes", err)
		os.Exit(1)
	}
	commitMsg := os.Getenv("COMMIT_MESSAGE")
	err = exec.Command("git", "commit", "-m", fmt.Sprintf(`"%v"`, commitMsg)).Run()
	if err != nil {
		logoru.Error("Failed to commit changes", err)
		os.Exit(1)
	}
	logoru.Success("Committed changes!")
}
