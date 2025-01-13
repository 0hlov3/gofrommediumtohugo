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
		{"<pre><code class=\"markup--code markup--pre-code\">❯ brew upgrade<br>brew upgrade</code></pre>", "```plaintext\n❯ brew upgrade\nbrew upgrade\n```"},
		{"<code class=\"markup--code markup--pre-code\">❯ brew upgrade<br>brew upgrade</code>", "`❯ brew upgrade brew upgrade`"},
		{"<pre data-code-block-mode=\"2\" spellcheck=\"false\" data-code-block-lang=\"yaml\" name=\"0568\" id=\"0568\" class=\"graf graf--pre graf-after--blockquote graf--preV2\">❯ brew upgrade<br>brew upgrade</pre>", "```yaml\n❯ brew upgrade\nbrew upgrade\n```"},
		{"<pre data-code-block-lang=\"bash\">sudo pacman -S --needed gnupg</pre>", "```bash\nsudo pacman -S --needed gnupg\n```"},
		{"<pre data-code-block-lang=\"yaml\">key: value</pre>", "```yaml\nkey: value\n```"},
		{"<pre>sudo pacman -S --needed gnupg</pre>", "```plaintext\nsudo pacman -S --needed gnupg\n```"},
	}

	for _, test := range tests {
		output := ConvertHTMLToMarkdown(test.input)
		if output != test.expected {
			t.Errorf("ConvertHTMLToMarkdown(%q) = %q; want %q", test.input, output, test.expected)
		}
	}
}
