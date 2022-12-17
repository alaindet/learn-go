package main

import (
	"html/template"
	"path/filepath"

	"snippetbox.dev/internal/models"
)

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}

func newTemplateCache(basePath string) (map[string]*template.Template, error) {

	cache := map[string]*template.Template{}

	globalTemplates := []string{
		basePath + "/base.html",
		basePath + "/partials/nav.html",
	}

	pages, err := filepath.Glob(basePath + "/pages/*.html")

	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		fileName := filepath.Base(page)
		templateFiles := append(globalTemplates, page)
		tmpl, err := template.ParseFiles(templateFiles...)

		if err != nil {
			return nil, err
		}

		cache[fileName] = tmpl
	}

	return cache, nil
}
