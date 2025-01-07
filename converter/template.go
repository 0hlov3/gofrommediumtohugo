// converter/template.go
package converter

import (
	"os"
	"text/template"

	"github.com/0hlov3/GoFromMediumToHugo/models"
)

var tmpl = template.Must(template.New("post").Parse(`---
title: "{{ .Title }}"
author: "{{ .Author }}"
date: {{ .Date }}
lastmod: {{ .Lastmod }}
{{ if .Draft }}draft: {{ .Draft }}{{ end }}
description: "{{ .Description }}"
subtitle: "{{ .Subtitle }}"
{{ if .Tags }}tags:
{{ range .Tags }} - {{ . }}
{{ end }}{{ end }}
{{ if .FeaturedImage }}image: "{{ .FeaturedImage }}" {{ end }}
{{ if .Images }}images:
{{ range .Images }} - "{{ . }}"
{{ end }}{{ end }}
{{ if .Canonical }}aliases:
  - "/{{ .Canonical }}"
{{ end }}
---

{{ .Body }}
`))

func Write(post *models.Post, path string) error {
	os.Remove(path)
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return tmpl.Execute(f, post)
}
