package source

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/Matt-Gleich/logoru"
)

// Get the content of a file from the raw url
// Returns the contents of that file
func getContent(url string) string {
	// Making get request
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		logoru.Error("Failed to get contents of", url, ";", err)
	}
	defer resp.Body.Close()
	// Reading the actual response
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logoru.Error("Failed to parse get request for file;", url, ";", err)
	}
	return string(data)
}

// Replace the github view of a url to a raw text view
// From: https://github.com/Matt-Gleich/go_template/blob/master/CONTRIBUTING.md
// To: https://raw.githubusercontent.com/Matt-Gleich/go_template/master/CONTRIBUTING.md
func rawURL(url string) string {
	chunks := strings.Split(url, "/")
	blobRemoved := strings.Join(append(chunks[:5], chunks[6:]...), "/")
	return strings.Replace(blobRemoved, "https://github.com", "https://raw.githubusercontent.com", 1)
}
