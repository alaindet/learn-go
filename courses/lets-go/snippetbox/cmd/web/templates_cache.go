package main

import (
	"html/template"
	"io/fs"
	"path/filepath"

	"snippetbox.dev/ui"
)

func newTemplateCache() (map[string]*template.Template, error) {

	cache := map[string]*template.Template{}

	pages, err := fs.Glob(ui.Files, "html/pages/*.html")

	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		fileName := filepath.Base(page)

		patterns := []string{
			"html/base.html",
			"html/partials/*.html",
			page,
		}

		ts := template.New(fileName).Funcs(templateFunctions)
		ts, err := ts.ParseFS(ui.Files, patterns...)

		if err != nil {
			return nil, err
		}

		cache[fileName] = ts
	}

	return cache, nil
}
