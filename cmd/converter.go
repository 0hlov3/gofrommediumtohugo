// cmd/converter.go
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"

	"github.com/0hlov3/GoFromMediumToHugo/converter"
)

var (
	postsHTMLFolder   string
	hugoContentFolder string
	contentType       string
)

var conv converter.Converter = converter.NewDefaultConverter()

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert Medium HTML files to Hugo markdown",
	Run: func(cmd *cobra.Command, args []string) {
		if postsHTMLFolder == "" || hugoContentFolder == "" || contentType == "" {
			fmt.Println("All flags are required.")
			cmd.Help()
			os.Exit(1)
		}

		fmt.Printf("Converting Medium posts from %s to %s with content type %s...\n",
			postsHTMLFolder, hugoContentFolder, contentType)

		// Call the conversion logic
		conv.Convert(postsHTMLFolder, hugoContentFolder, contentType)
	},
}

func init() {
	convertCmd.Flags().StringVarP(&postsHTMLFolder, "posts", "p", "", "Path to Medium HTML posts")
	convertCmd.Flags().StringVarP(&hugoContentFolder, "output", "o", "", "Output folder for Hugo content")
	convertCmd.Flags().StringVarP(&contentType, "type", "t", "posts", "Hugo content type")

	rootCmd.AddCommand(convertCmd)
}
