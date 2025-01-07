package converter

import (
	"testing"
)

func TestDefaultConverter(t *testing.T) {
	// Mock the Convert function
	mockCalled := false
	mockPostsHTMLFolder := ""
	mockHugoContentFolder := ""
	mockContentType := ""

	mockConvert := func(postsHTMLFolder, hugoContentFolder, contentType string) {
		mockCalled = true
		mockPostsHTMLFolder = postsHTMLFolder
		mockHugoContentFolder = hugoContentFolder
		mockContentType = contentType
	}

	// Create an instance of DefaultConverter with the mock
	converter := &DefaultConverter{
		ConvertFunc: mockConvert,
	}

	// Call the Convert method
	testPosts := "/path/to/posts"
	testHugoContent := "/path/to/output"
	testType := "posts"
	converter.Convert(testPosts, testHugoContent, testType)

	// Assertions
	if !mockCalled {
		t.Error("Convert was not called")
	}
	if mockPostsHTMLFolder != testPosts {
		t.Errorf("Convert called with incorrect postsHTMLFolder: got %q, want %q", mockPostsHTMLFolder, testPosts)
	}
	if mockHugoContentFolder != testHugoContent {
		t.Errorf("Convert called with incorrect hugoContentFolder: got %q, want %q", mockHugoContentFolder, testHugoContent)
	}
	if mockContentType != testType {
		t.Errorf("Convert called with incorrect contentType: got %q, want %q", mockContentType, testType)
	}
}
