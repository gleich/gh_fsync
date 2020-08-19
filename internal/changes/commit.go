package changes

import (
	"github.com/Matt-Gleich/logoru"
	"github.com/go-git/go-git/v5"
)

func Commit() {
	logoru.Info("Commiting changes")
	repo, err := git.PlainOpen(".")
	if err != nil {
		logoru.Error("Failed to open git repo", err)
	}
	tree, err := repo.Worktree()
	if err != nil {
		logoru.Error("Failed to get working tree", err)
	}

	_, err = tree.Commit("", &git.CommitOptions{
		All: true,
	})
	if err != nil {
		logoru.Error("Failed to commit changes")
	}
	logoru.Success("Committed changes!")
}
