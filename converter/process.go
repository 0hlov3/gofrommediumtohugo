// converter/process.go
package converter

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/0hlov3/GoFromMediumToHugo/models"
	"github.com/0hlov3/GoFromMediumToHugo/utils"
	"github.com/PuerkitoBio/goquery"
)

// HTML Selectors
const (
	TitleSelector         = "h1.p-name, h3.graf--h3.graf--leading.graf--title"
	SubtitleSelector      = "h3.graf--h3.graf-after--h3"
	BodySelector          = "div.section-inner"
	FeaturedImageSelector = "img[data-is-featured='true']"
	AuthorSelector        = "footer .p-author.h-card"
	DateSelector          = "footer time.dt-published"
	CanonicalSelector     = "footer .p-canonical"
)

func Convert(postsHTMLFolder, hugoContentFolder, contentType string) {
	files, err := os.ReadDir(postsHTMLFolder)
	if err != nil {
		fmt.Printf("Error reading directory: %s\n", err)
		return
	}

	err = os.MkdirAll(hugoContentFolder, os.ModePerm)
	if err != nil {
		fmt.Printf("Error creating directory: %s\n", err)
		return
	}

	fmt.Printf("Found %d articles.\n", len(files))

	for i, f := range files {
		if !strings.HasSuffix(f.Name(), ".html") || f.IsDir() {
			logSkip(f)
			continue
		}

		fmt.Printf("[%d/%d] Processing: %s... \n", i+1, len(files), f.Name())
		outpath, err := processFile(f, postsHTMLFolder, hugoContentFolder, contentType)

		if err == nil {
			fmt.Printf("\033[32m✔ Wrote article to %s\033[0m\n", outpath)
		} else {
			fmt.Printf("\033[31m✘ Failed to write article: %v\033[0m\n", err)
		}
	}
}

func processFile(file os.DirEntry, postsHTMLFolder, hugoContentFolder, contentType string) (string, error) {
	inpath := filepath.Join(postsHTMLFolder, file.Name())
	doc, err := read(inpath)
	if err != nil {
		return "", fmt.Errorf("error reading HTML file %s: %w", file.Name(), err)
	}

	post, err := processPost(doc, file, hugoContentFolder, contentType)
	if err != nil {
		return "", fmt.Errorf("error processing post %s: %w", file.Name(), err)
	}

	outpath := filepath.Join(post.HddFolder, "index.md")
	if err := Write(post, outpath); err != nil {
		return "", fmt.Errorf("error writing post %s: %w", file.Name(), err)
	}
	return outpath, nil
}

func logSkip(file os.DirEntry) {
	if file.IsDir() {
		fmt.Printf("Skipping directory: %s\n", file.Name())
	} else {
		fmt.Printf("Skipping non-HTML file: %s\n", file.Name())
	}
}

func processPost(doc *goquery.Document, file os.DirEntry, contentFolder, contentType string) (*models.Post, error) {
	post := models.NewPost()

	// Extract metadata
	post.Title = strings.TrimSpace(doc.Find(TitleSelector).First().Text())
	post.Subtitle = strings.TrimSpace(doc.Find(SubtitleSelector).Text())
	post.Author = strings.TrimSpace(doc.Find(AuthorSelector).Text())
	post.SetDate(doc.Find(DateSelector).AttrOr("datetime", ""))
	post.Canonical = strings.TrimSpace(doc.Find(CanonicalSelector).AttrOr("href", ""))

	// Generate slug and prepare directories
	slug := utils.GenerateSlug(post.Title)
	if slug == "" {
		slug = "noname"
	}
	pageBundle := fmt.Sprintf("%s_%s", post.Date, slug)
	post.HddFolder = filepath.Join(contentFolder, contentType, pageBundle)

	if err := os.RemoveAll(post.HddFolder); err != nil {
		return nil, fmt.Errorf("error cleaning folder: %w", err)
	}
	if err := os.MkdirAll(post.HddFolder, os.ModePerm); err != nil {
		return nil, fmt.Errorf("error creating folder: %w", err)
	}

	// Fetch and replace images
	images, featuredImage, err := FetchAndReplaceImages(doc, post.HddFolder)
	if err != nil {
		return nil, fmt.Errorf("error fetching images: %w", err)
	}
	post.Images = images
	post.FeaturedImage = featuredImage

	// Extract the updated body after images are replaced
	post.Body = docToMarkdown(doc) // Use the updated DOM here

	return post, nil
}

func docToMarkdown(doc *goquery.Document) string {
	var body strings.Builder
	doc.Find(BodySelector).Each(func(i int, s *goquery.Selection) {
		html, _ := s.Html()
		body.WriteString(utils.ConvertHTMLToMarkdown(html))
		body.WriteString("\n\n")
	})
	return strings.TrimSpace(body.String())
}

func read(path string) (*goquery.Document, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return goquery.NewDocumentFromReader(file)
}
