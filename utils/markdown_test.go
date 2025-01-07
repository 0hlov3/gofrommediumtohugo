package utils

import (
	"testing"
)

func TestConvertHTMLToMarkdown(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		// Test a basic paragraph
		{"<p>Hello, World!</p>", "Hello, World!"},

		// Test strong and emphasis
		{"<strong>Bold</strong> and <em>italic</em>", "**Bold** and *italic*"},

		// Test a link
		{"<a href=\"https://example.com\">Example</a>", "[Example](https://example.com)"},

		// Test a list
		{"<ul><li>Item 1</li><li>Item 2</li></ul>", "- Item 1\n- Item 2"},

		// Test an image
		{"<img src=\"https://example.com/image.jpg\" alt=\"Example Image\">", "![Example Image](https://example.com/image.jpg)"},

		// Test mixed HTML
		{"<h1>Heading</h1><p>Paragraph</p>", "# Heading\n\nParagraph"},

		// Test trimming whitespace
		{"   <p>Trimmed</p>   ", "Trimmed"},

		// Test code blocks
		{"<pre><code class=\"markup--code markup--pre-code\">❯ brew upgrade<br>brew upgrade</code></pre>", "```\n❯ brew upgrade\nbrew upgrade\n```"},
		{"<code class=\"markup--code markup--pre-code\">❯ brew upgrade<br>brew upgrade</code>", "`❯ brew upgrade brew upgrade`"},
	}

	for _, test := range tests {
		output := ConvertHTMLToMarkdown(test.input)
		if output != test.expected {
			t.Errorf("ConvertHTMLToMarkdown(%q) = %q; want %q", test.input, output, test.expected)
		}
	}
}
