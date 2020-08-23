package git

import "os"

// Change the GITHUB_TOKEN to the GITHUB_PAT variable if it exists
func Authenticate() {
	pat := os.Getenv("INPUT_PERSONAL_ACCESS_TOKEN")
	if pat != "" {
		os.Setenv("GITHUB_TOKEN", pat)
	}
}
