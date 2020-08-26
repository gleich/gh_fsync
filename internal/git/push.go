package git

import (
	"fmt"
	"os"
	"strings"

	"github.com/Matt-Gleich/gh_fsync/pkg/utils"
	"github.com/Matt-Gleich/logoru"
)

// Push changes
func Push() {
	logoru.Info("🚀 Pushing changes to master")
	utils.RunCommand("git", []string{
		"push",
		getOriginURL(),
	},
		"Failed to push changes",
	)
	logoru.Success("✅ Pushed changes to master! Have a good day 👋")
}

// Get origin url
func getOriginURL() string {
	pat := os.Getenv("INPUT_PERSONAL_ACCESS_TOKEN")
	url := strings.TrimSuffix(utils.RunCommand("git", []string{"remote", "get-url", "origin"}, "Failed to get remote url"), "\n")
	if pat != "" {
		username := os.Getenv("INPUT_GITHUB_USERNAME")
		if username == "" {
			logoru.Error("Failed to get GitHub username")
			os.Exit(1)
		}
		url = fmt.Sprintf("https://%v:%v@%v",
			username,
			pat,
			url[8:],
		)
	}
	if url[len(url)-4:] != ".git" {
		url = url + ".git"
	}
	return url
}
