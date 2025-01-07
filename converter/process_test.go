package converter

import (
	"github.com/PuerkitoBio/goquery"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestConvert(t *testing.T) {
	// Setup: Create a temporary input directory
	inputDir := t.TempDir()
	outputDir := t.TempDir()

	// Create dummy HTML files
	dummyHTML := `
		<html>
			<head><title>Test Post</title></head>
			<body>
				<h1 class="p-name">Test Title</h1>
				<div class="section-inner">
					<p>This is a test paragraph.</p>
				</div>
				<footer>
					<a class="p-author h-card">Test Author</a>
					<time class="dt-published" datetime="2025-01-01T00:00:00Z"></time>
				</footer>
			</body>
		</html>
	`
	dummyFilePath := filepath.Join(inputDir, "test-post.html")
	if err := os.WriteFile(dummyFilePath, []byte(dummyHTML), 0644); err != nil {
		t.Fatalf("Failed to create dummy HTML file: %v", err)
	}

	// Call the Convert function
	Convert(inputDir, outputDir, "posts")

	// Validate the output
	outputFilePath := filepath.Join(outputDir, "posts", "2025-01-01T00:00:00Z_test-title", "index.md")
	if _, err := os.Stat(outputFilePath); os.IsNotExist(err) {
		t.Errorf("Expected output file not created: %s", outputFilePath)
		return
	}

	// Read and verify the content
	content, err := os.ReadFile(outputFilePath)
	if err != nil {
		t.Fatalf("Failed to read output file: %v", err)
	}

	if !strings.Contains(string(content), "Test Title") {
		t.Errorf("Output file does not contain expected title: %s", content)
	}

	if !strings.Contains(string(content), "This is a test paragraph.") {
		t.Errorf("Output file does not contain expected body content: %s", content)
	}

	if !strings.Contains(string(content), "Test Author") {
		t.Errorf("Output file does not contain expected author: %s", content)
	}

	if !strings.Contains(string(content), "2025-01-01T00:00:00Z") {
		t.Errorf("Output file does not contain expected date: %s", content)
	}
}

func TestDocToMarkdown(t *testing.T) {
	html := `
		<div class="section-inner">
			<p>This is a paragraph.</p>
		</div>
	`
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(html))
	markdown := docToMarkdown(doc)

	expected := "This is a paragraph."
	if markdown != expected {
		t.Errorf("docToMarkdown() = %q; want %q", markdown, expected)
	}
}

func TestProcessPost(t *testing.T) {
	html := `
		<html>
			<head><title>Test Post</title></head>
			<body>
				<h1 class="p-name">Test Title</h1>
				<div class="section-inner">
					<p>This is a paragraph.</p>
				</div>
				<footer>
					<a class="p-author h-card">Test Author</a>
					<time class="dt-published" datetime="2025-01-01T00:00:00Z"></time>
				</footer>
			</body>
		</html>
	`
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(html))
	post, err := processPost(doc, nil, t.TempDir(), "posts")
	if err != nil {
		t.Fatalf("processPost() returned an error: %v", err)
	}

	if post.Title != "Test Title" {
		t.Errorf("post.Title = %q; want %q", post.Title, "Test Title")
	}
	if post.Author != "Test Author" {
		t.Errorf("post.Author = %q; want %q", post.Author, "Test Author")
	}
	if !strings.Contains(post.Body, "This is a paragraph.") {
		t.Errorf("post.Body = %q; want it to contain %q", post.Body, "This is a paragraph.")
	}
}
