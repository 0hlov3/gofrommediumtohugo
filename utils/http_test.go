package utils

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func TestDownloadFile(t *testing.T) {
	// Setup a mock HTTP server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Serve a mock response
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("mock file content"))
	}))
	defer mockServer.Close()

	// Create a temporary file for download
	tempFile, err := os.CreateTemp("", "downloaded_file_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tempFile.Name()) // Clean up after test

	// Call DownloadFile
	err = DownloadFile(mockServer.URL, tempFile.Name())
	if err != nil {
		t.Fatalf("DownloadFile() returned an error: %v", err)
	}

	// Read the file content
	content, err := ioutil.ReadFile(tempFile.Name())
	if err != nil {
		t.Fatalf("Failed to read downloaded file: %v", err)
	}

	// Verify the content matches the mock response
	expectedContent := "mock file content"
	if string(content) != expectedContent {
		t.Errorf("Downloaded file content = %q; want %q", string(content), expectedContent)
	}
}

func TestCreateHTTPClient(t *testing.T) {
	// Create an HTTP client with a 2-second timeout
	client := CreateHTTPClient(2)

	// Verify the timeout is set correctly
	expectedTimeout := 2
	if client.Timeout.Seconds() != float64(expectedTimeout) {
		t.Errorf("HTTP client timeout = %v; want %v", client.Timeout.Seconds(), expectedTimeout)
	}

	// Test the client with a slow server
	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Delay response beyond the timeout
		time.Sleep(3 * time.Second)
		w.WriteHeader(http.StatusOK)
	}))
	defer slowServer.Close()

	// Perform a request with the client
	resp, err := client.Get(slowServer.URL)
	if err == nil {
		resp.Body.Close()
		t.Error("Expected timeout error, but got none")
	}
}
