package main

import (
	"github.com/Matt-Gleich/gh_fsync/internal/changes"
	"github.com/Matt-Gleich/gh_fsync/internal/config"
	"github.com/Matt-Gleich/gh_fsync/internal/source"
	"github.com/Matt-Gleich/gh_fsync/internal/write"
)

func main() {
	configuration := config.Read()
	sourceFiles := source.GetFromSource(configuration)
	write.WriteChanges(sourceFiles)
	repo := changes.Commit()
	changes.Push(repo)
}
