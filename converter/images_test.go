package converter

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestFetchAndReplaceImages(t *testing.T) {
	// Create a mock HTTP server to serve image files
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("mock image content"))
	}))
	defer mockServer.Close()

	// Create a temporary folder for the article
	tempFolder := t.TempDir()

	// Sample HTML document
	html := `
		<html>
			<body>
				<img src="` + mockServer.URL + `/image1.jpg" data-is-featured="true"/>
				<img src="` + mockServer.URL + `/image2.jpg"/>
			</body>
		</html>
	`
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader([]byte(html)))
	if err != nil {
		t.Fatalf("Failed to parse HTML: %v", err)
	}

	// Call FetchAndReplaceImages
	images, featuredImage, err := FetchAndReplaceImages(doc, tempFolder)
	if err != nil {
		t.Fatalf("FetchAndReplaceImages returned an error: %v", err)
	}

	// Verify results
	expectedImages := []string{
		"1.jpg",
		"2.jpg",
	}
	if len(images) != len(expectedImages) {
		t.Errorf("Expected %d images, got %d", len(expectedImages), len(images))
	}
	for i, img := range images {
		if img != expectedImages[i] {
			t.Errorf("Expected image %q, got %q", expectedImages[i], img)
		}
	}

	// Verify the featured image
	expectedFeaturedImage := "1.jpg"
	if featuredImage != expectedFeaturedImage {
		t.Errorf("Expected featured image %q, got %q", expectedFeaturedImage, featuredImage)
	}

	// Verify files were created
	for _, img := range images {
		localPath := filepath.Join(tempFolder, img)
		if _, err := os.Stat(localPath); os.IsNotExist(err) {
			t.Errorf("Image file %q was not created", localPath)
		}
	}
}
