package models

import (
	"testing"
)

func TestNewPost(t *testing.T) {
	post := NewPost()

	// Check if Lastmod is correctly set to the current time
	if post.Lastmod == "" {
		t.Errorf("NewPost() did not set Lastmod")
	}

	// Verify that other fields are initialized correctly
	if post.Title != "" || post.Author != "" || post.Body != "" {
		t.Errorf("NewPost() initialized unexpected fields: %+v", post)
	}
}

func TestSetDate(t *testing.T) {
	post := NewPost()

	// Test with an empty date
	post.SetDate("")
	if post.Date == "" {
		t.Errorf("SetDate() failed to set current date when input was empty")
	}

	// Test with a specific date
	expectedDate := "2025-01-01T00:00:00Z"
	post.SetDate(expectedDate)
	if post.Date != expectedDate {
		t.Errorf("SetDate() = %q; want %q", post.Date, expectedDate)
	}
}

func TestAddTag(t *testing.T) {
	post := NewPost()

	// Add tags
	post.AddTag("tag1")
	post.AddTag("tag2")

	// Verify tags
	if len(post.Tags) != 2 || post.Tags[0] != "tag1" || post.Tags[1] != "tag2" {
		t.Errorf("AddTag() failed, got %+v", post.Tags)
	}
}

func TestAddImage(t *testing.T) {
	post := NewPost()

	// Add images
	post.AddImage("image1.jpg")
	post.AddImage("image2.png")

	// Verify images
	if len(post.Images) != 2 || post.Images[0] != "image1.jpg" || post.Images[1] != "image2.png" {
		t.Errorf("AddImage() failed, got %+v", post.Images)
	}
}

func TestSetFeaturedImage(t *testing.T) {
	post := NewPost()

	// Set a featured image
	featuredImage := "featured.jpg"
	post.SetFeaturedImage(featuredImage)

	// Verify the featured image
	if post.FeaturedImage != featuredImage {
		t.Errorf("SetFeaturedImage() = %q; want %q", post.FeaturedImage, featuredImage)
	}
}
