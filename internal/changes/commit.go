package changes

import (
	"os"

	"github.com/Matt-Gleich/logoru"
	"github.com/go-git/go-git/v5"
)

func Commit() git.Repository {
	logoru.Info("Commiting changes")
	repo, err := git.PlainOpen(".")
	if err != nil {
		logoru.Error("Failed to open git repo", err)
		os.Exit(1)
	}
	tree, err := repo.Worktree()
	if err != nil {
		logoru.Error("Failed to get working tree", err)
		os.Exit(1)
	}

	_, err = tree.Commit("", &git.CommitOptions{
		All: true,
	})
	if err != nil {
		logoru.Error("Failed to commit changes", err)
		os.Exit(1)
	}
	logoru.Success("Committed changes!")
	return *repo
}
