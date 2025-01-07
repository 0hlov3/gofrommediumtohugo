// utils/markdown.go
package utils

import (
	htmltomarkdown "github.com/JohannesKaufmann/html-to-markdown/v2"
	"strings"
)

// ConvertHTMLToMarkdown converts an HTML string to a Markdown string.
func ConvertHTMLToMarkdown(html string) string {
	// Perform the conversion
	markdown, err := htmltomarkdown.ConvertString(html)
	if err != nil {
		// Handle errors gracefully
		return ""
	}

	// Trim whitespace from the final output
	return strings.TrimSpace(markdown)
}
