// converter/images.go
package converter

import (
	"fmt"
	"github.com/0hlov3/GoFromMediumToHugo/utils"
	"github.com/PuerkitoBio/goquery"
	"path/filepath"
)

func FetchAndReplaceImages(doc *goquery.Document, folder string) ([]string, string, error) {
	images := doc.Find("img")
	if images.Length() == 0 {
		return nil, "", nil
	}

	var index int
	var featuredImage string
	var result []string

	images.Each(func(i int, imgDomElement *goquery.Selection) {
		index++
		original, has := imgDomElement.Attr("src")
		if !has {
			fmt.Print("warning: image missing src attribute\n")
			return
		}

		ext := filepath.Ext(original)
		if len(ext) < 2 {
			ext = ".jpg"
		}
		filename := fmt.Sprintf("%d%s", index, ext)
		imagePath := filepath.Join(folder, filename)

		// Use the DownloadFile utility
		err := utils.DownloadFile(original, imagePath)
		if err != nil {
			fmt.Printf("error downloading image: %s\n", err)
			return
		}

		// Update the <img> tag's src attribute to the local filename
		imgDomElement.SetAttr("src", filename)
		result = append(result, filename)

		// Check if the image is the featured image
		if _, isFeatured := imgDomElement.Attr("data-is-featured"); isFeatured {
			featuredImage = filename
		}
	})

	return result, featuredImage, nil
}
