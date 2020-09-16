package git

import (
	"github.com/Matt-Gleich/gh_fsync/pkg/utils"
	"github.com/Matt-Gleich/logoru"
)

// Push changes
func Push() {
	logoru.Info("🚀 Pushing changes to master")
	utils.RunCommand("git", []string{"push"},
		"Failed to push changes",
	)
	logoru.Success("✅ Pushed changes to master! Have a good day 👋")
}
