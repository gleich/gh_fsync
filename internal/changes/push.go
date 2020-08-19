package changes

import (
	"github.com/Matt-Gleich/logoru"
	"github.com/go-git/go-git/v5"
)

func Push(repo git.Repository) {
	logoru.Info("Pushing changes")
	err := repo.Push(&git.PushOptions{})
	if err != nil {
		logoru.Error("Failed to push changes", err)
	}
	logoru.Success("Pushed changes to master! Have a good day 👋")
}
