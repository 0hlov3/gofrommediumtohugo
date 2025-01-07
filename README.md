# GoFromMediumToHugo

GoFromMediumToHugo is a CLI tool written in Go to convert Medium export HTML files into Markdown files compatible with Hugo, a popular static site generator. It handles metadata extraction, image downloading, and content formatting.

## Features

- Converts Medium HTML files into Hugo-compatible Markdown files.
- Extracts metadata such as title, subtitle, author, and publication date.
- Downloads images locally and updates references in the Markdown files.
- Automatically generates slugs and organizes content into page bundles.

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/0hlov3/GoFromMediumToHugo.git
   cd GoFromMediumToHugo
   ```

2. Build the binary:
   ```bash
   go build -o mediumtohugo main.go
   ```

3. (Optional) Install the binary globally:
   ```bash
   mv mediumtohugo /usr/local/bin/
   ```

## Dependencies

This project uses the following Go packages:

- [github.com/PuerkitoBio/goquery](https://github.com/PuerkitoBio/goquery): For parsing and manipulating HTML documents.
- [github.com/spf13/cobra](https://github.com/spf13/cobra): For building the CLI.
- [github.com/spf13/viper](https://github.com/spf13/viper): For configuration management.
- [github.com/JohannesKaufmann/html-to-markdown](https://github.com/JohannesKaufmann/html-to-markdown): For converting HTML to Markdown.

Install dependencies by running:
```bash
go mod tidy
```

## Usage

Run the tool with the following command:
```bash
mediumtohugo convert --posts <medium_posts_folder> --output <output_folder> --type <content_type>
```

### Parameters:

- `--posts, -p`: Path to the folder containing Medium export HTML files.
- `--output, -o`: Path to the folder where Hugo-compatible content will be saved.
- `--type, -t`: Hugo content type (e.g., `posts`). Default: `posts`.

### Example:
```bash
mediumtohugo convert --posts ~/Downloads/medium-export/posts/ --output ~/my-hugo-site/content/posts/ --type posts
```

### Output Structure

The tool organizes files into Hugo page bundles:
```
<output_folder>/
  <content_type>/
    <date>_<slug>/
      index.md
      <image_files>
```

### Configuration File (Optional)

You can create a `config.yaml` file for default parameters:
```yaml
postsHTMLFolder: "~/Downloads/medium-export/posts/"
hugoContentFolder: "~/my-hugo-site/content/posts/"
contentType: "posts"
```
Run the tool without flags to use the config file:
```bash
mediumtohugo convert
```

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests for improvements or new features.

## License

This project is licensed under the [MIT License](LICENSE).
