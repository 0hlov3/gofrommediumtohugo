// models/post.go
package models

import (
	"time"
)

type Post struct {
	Title, Author, Body   string
	Date, Lastmod         string
	Subtitle, Description string
	Canonical, FullURL    string
	FeaturedImage         string
	Images                []string
	Tags                  []string
	HddFolder             string
	Draft                 bool
	IsComment             bool
}

func NewPost() *Post {
	return &Post{
		Lastmod: time.Now().Format(time.RFC3339),
	}
}

func (p *Post) SetDate(date string) {
	if date == "" {
		p.Date = time.Now().Format(time.RFC3339)
	} else {
		p.Date = date
	}
}

func (p *Post) AddTag(tag string) {
	p.Tags = append(p.Tags, tag)
}

func (p *Post) AddImage(image string) {
	p.Images = append(p.Images, image)
}

func (p *Post) SetFeaturedImage(image string) {
	p.FeaturedImage = image
}
