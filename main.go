package main

import (
	"github.com/Matt-Gleich/gh_fsync/pkg/config"
	"github.com/Matt-Gleich/gh_fsync/pkg/git"
	"github.com/Matt-Gleich/gh_fsync/pkg/source"
	"github.com/Matt-Gleich/gh_fsync/pkg/write"
)

func main() {
	configuration := config.Read()
	sourceFiles := source.GetFromSource(configuration)
	changes := source.GetChanges(sourceFiles)
	write.WriteChanges(changes)
	git.Commit(configuration)
	git.Push()
}
