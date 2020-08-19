package main

import (
	"fmt"

	"github.com/Matt-Gleich/gh_fsync/internal/config"
	"github.com/Matt-Gleich/gh_fsync/internal/source"
)

func main() {
	configuration := config.Read()
	sourceFiles := source.GetFromSource(configuration)
	fmt.Printf("%#v", sourceFiles)
}
