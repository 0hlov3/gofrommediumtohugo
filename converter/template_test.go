package converter

import (
	"os"
	"strings"
	"testing"

	"github.com/0hlov3/GoFromMediumToHugo/models"
)

func TestWrite(t *testing.T) {
	// Create a temporary file
	tempFile, err := os.CreateTemp("", "test_post_*.md")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tempFile.Name()) // Clean up after test

	// Create a mock Post
	post := &models.Post{
		Title:         "Test Title",
		Author:        "Test Author",
		Date:          "2025-01-01T00:00:00Z",
		Lastmod:       "2025-01-02T00:00:00Z",
		Subtitle:      "Test Subtitle",
		Description:   "Test Description",
		Canonical:     "test-canonical",
		Body:          "This is the body of the test post.",
		FeaturedImage: "featured.jpg",
		Tags:          []string{"tag1", "tag2"},
		Images:        []string{"image1.jpg", "image2.jpg"},
		Draft:         false,
	}

	// Call Write function
	err = Write(post, tempFile.Name())
	if err != nil {
		t.Fatalf("Write() returned an error: %v", err)
	}

	// Read the file and validate content
	content, err := os.ReadFile(tempFile.Name())
	if err != nil {
		t.Fatalf("Failed to read temporary file: %v", err)
	}

	output := string(content)

	// Verify the content contains expected values
	if !strings.Contains(output, "title: \"Test Title\"") {
		t.Errorf("Output missing title: %s", output)
	}
	if !strings.Contains(output, "author: \"Test Author\"") {
		t.Errorf("Output missing author: %s", output)
	}
	if !strings.Contains(output, "date: 2025-01-01T00:00:00Z") {
		t.Errorf("Output missing date: %s", output)
	}
	if !strings.Contains(output, "lastmod: 2025-01-02T00:00:00Z") {
		t.Errorf("Output missing lastmod: %s", output)
	}
	if !strings.Contains(output, "subtitle: \"Test Subtitle\"") {
		t.Errorf("Output missing subtitle: %s", output)
	}
	if !strings.Contains(output, "description: \"Test Description\"") {
		t.Errorf("Output missing description: %s", output)
	}
	if !strings.Contains(output, "This is the body of the test post.") {
		t.Errorf("Output missing body: %s", output)
	}
	if !strings.Contains(output, "aliases:\n  - \"/test-canonical\"") {
		t.Errorf("Output missing canonical alias: %s", output)
	}
	if !strings.Contains(output, "image: \"featured.jpg\"") {
		t.Errorf("Output missing featured image: %s", output)
	}
	if !strings.Contains(output, "- tag1") || !strings.Contains(output, "- tag2") {
		t.Errorf("Output missing tags: %s", output)
	}
	if !strings.Contains(output, "- \"image1.jpg\"") || !strings.Contains(output, "- \"image2.jpg\"") {
		t.Errorf("Output missing images: %s", output)
	}
}
