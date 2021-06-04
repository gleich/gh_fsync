package main

import (
	"github.com/gleich/gh_fsync/pkg/config"
	"github.com/gleich/gh_fsync/pkg/git"
	"github.com/gleich/gh_fsync/pkg/source"
	"github.com/gleich/gh_fsync/pkg/write"
)

func main() {
	configuration := config.Read()
	sourceFiles := source.GetFromSource(configuration)
	changes := source.GetChanges(sourceFiles)
	write.WriteChanges(changes)
	git.Commit(configuration)
	git.Push()
}
