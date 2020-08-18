package main

import (
	"fmt"

	"github.com/Matt-Gleich/gh_fsync/internal/config"
)

func main() {
	configuration := config.Read()
	fmt.Printf("%#v", configuration)
}
