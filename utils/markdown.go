// utils/markdown.go
package utils

import (
	"github.com/JohannesKaufmann/dom"
	//htmltomarkdown "github.com/JohannesKaufmann/html-to-markdown/v2"
	"github.com/JohannesKaufmann/html-to-markdown/v2/converter"
	"github.com/JohannesKaufmann/html-to-markdown/v2/plugin/base"
	"github.com/JohannesKaufmann/html-to-markdown/v2/plugin/commonmark"
	"golang.org/x/net/html"
	"strings"
)

// ConvertHTMLToMarkdown converts an HTML string to a Markdown string.
func ConvertHTMLToMarkdown(html string) string {
	conv := converter.NewConverter(
		converter.WithPlugins(
			base.NewBasePlugin(),
			commonmark.NewCommonmarkPlugin(),
		),
	)

	conv.Register.RendererFor("pre", converter.TagTypeBlock, renderPreTag, converter.PriorityEarly)

	// Perform the conversion
	// markdown, err := htmltomarkdown.ConvertString(html)
	markdown, err := conv.ConvertString(html)
	if err != nil {
		// Handle errors gracefully
		return ""
	}

	// Trim whitespace from the final output
	return strings.TrimSpace(markdown)
}

func renderPreTag(_ converter.Context, w converter.Writer, node *html.Node) converter.RenderStatus {
	// Extract the content of the <pre> tag
	codeContent, _ := getCodeWithoutTags(node)

	// Get the language from the `data-code-block-lang` attribute
	lang := dom.GetAttributeOr(node, "data-code-block-lang", "")
	if lang == "" {
		lang = "plaintext" // Default to plaintext if no language is specified
	}
	// Write the Markdown code block with the language and content
	w.WriteString("```" + lang + "\n")
	w.Write([]byte(codeContent))
	w.WriteString("\n```")

	return converter.RenderSuccess
}

func getCodeWithoutTags(startNode *html.Node) (string, string) {
	var buf strings.Builder

	var f func(*html.Node)
	f = func(n *html.Node) {
		// Ignore non-relevant nodes
		if n.Type == html.ElementNode && (n.Data == "style" || n.Data == "script" || n.Data == "textarea") {
			return
		}
		if n.Type == html.ElementNode && (n.Data == "br" || n.Data == "div") {
			buf.WriteString("\n")
		}

		// Add text nodes
		if n.Type == html.TextNode {
			buf.WriteString(n.Data)
		}

		// Process child nodes recursively
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(startNode)

	// Return the content as a string
	return buf.String(), ""
}
