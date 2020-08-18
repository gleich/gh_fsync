package source

import "strings"

// Replace the github view of a url to a raw text view
// From: https://github.com/Matt-Gleich/go_template/blob/master/CONTRIBUTING.md
// To: https://raw.githubusercontent.com/Matt-Gleich/go_template/master/CONTRIBUTING.md
func rawURL(url string) string {
	chunks := strings.Split(url, "/")
	blobRemoved := strings.Join(append(chunks[:5], chunks[6:]...), "/")
	return strings.Replace(blobRemoved, "https://github.com", "https://raw.githubusercontent.com", 1)
}
