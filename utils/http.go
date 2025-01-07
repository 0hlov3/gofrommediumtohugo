// utils/http.go
package utils

import (
	"io"
	"net/http"
	"os"
	"time"
)

// DownloadFile downloads a file from the given URL and saves it to the specified filepath.
func DownloadFile(url, filepath string) error {
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

// CreateHTTPClient creates an HTTP client with a custom timeout.
func CreateHTTPClient(timeoutSeconds int) *http.Client {
	return &http.Client{
		Timeout: time.Duration(timeoutSeconds) * time.Second,
	}
}
