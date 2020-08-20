package main

import (
	"github.com/Matt-Gleich/gh_fsync/internal/config"
	"github.com/Matt-Gleich/gh_fsync/internal/git"
	"github.com/Matt-Gleich/gh_fsync/internal/source"
	"github.com/Matt-Gleich/gh_fsync/internal/write"
)

func main() {
	configuration := config.Read()
	sourceFiles := source.GetFromSource(configuration)
	changes := source.GetChanges(sourceFiles)
	write.WriteChanges(changes)
	git.Commit()
	git.Push()
}
