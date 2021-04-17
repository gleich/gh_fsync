package git

import (
	"github.com/Matt-Gleich/gh_fsync/pkg/config"
	"github.com/Matt-Gleich/gh_fsync/pkg/utils"
	"github.com/Matt-Gleich/logoru"
)

// Commit the changes
func Commit(config config.Outline) {
	logoru.Info("🤝 Commiting changes")
	commitMsg := config.CommitMessage
	if commitMsg == "" {
		commitMsg = "Update via sync"
	}
	utils.RunCommand(
		"git",
		[]string{"config", "--global", "user.email", "\"action@github.com\""},
		"Failed to config user email",
	)
	utils.RunCommand(
		"git",
		[]string{"config", "--global", "user.name", "\"Publishing Bot\""},
		"Failed to config user name",
	)
	utils.RunCommand("git", []string{"add", "."}, "Failed to stage changes")
	utils.RunCommand("git", []string{"status"}, "Failed to get git status")
	utils.RunCommand("git", []string{"commit", "-m", commitMsg}, "Failed to commit changes")
	logoru.Success("✅ Committed changes")
}
