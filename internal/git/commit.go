package git

import (
	"os"

	"github.com/Matt-Gleich/gh_fsync/pkg/utils"
	"github.com/Matt-Gleich/logoru"
)

func Commit() {
	logoru.Info("Commiting changes")
	commitMsg := os.Getenv("INPUT_COMMIT_MESSAGE")
	utils.RunCommand("git", []string{"config", "--global", "user.email", "\"action@github.com\""}, "Failed to config user email")
	utils.RunCommand("git", []string{"config", "--global", "user.name", "\"Publishing Bot\""}, "Failed to config user name")
	utils.RunCommand("git", []string{"add", "."}, "Failed to stage changes")
	utils.RunCommand("git", []string{"status"}, "Failed to get git status")
	utils.RunCommand("git", []string{"commit", "-m", commitMsg}, "Failed to commit changes")
	utils.RunCommand("git", []string{"push"}, "Failed to push changes")
	logoru.Success("Committed changes!")
}
