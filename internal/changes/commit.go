package changes

import (
	"fmt"
	"os"

	"github.com/Matt-Gleich/gh_fsync/pkg/utils"
	"github.com/Matt-Gleich/logoru"
)

func Commit() {
	logoru.Info("Commiting changes")
	commitMsg := os.Getenv("INPUT_COMMIT_MESSAGE")
	utils.RunCommand("git config --global user.email \"action@github.com\"", "Failed to config user email")
	utils.RunCommand("git config --global user.name \"Publishing Bot\"", "Failed to config user name")
	utils.RunCommand("git add .", "Failed to stage changes")
	utils.RunCommand(fmt.Sprintf("git commit -m \"%v\"", commitMsg), "Failed to commit changes")
	utils.RunCommand("git push", "Failed to push changes")
	logoru.Success("Committed changes!")
}
