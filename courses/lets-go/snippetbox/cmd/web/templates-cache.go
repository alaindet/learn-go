package main

import (
	"html/template"
	"path/filepath"
)

func newTemplateCache(basePath string) (map[string]*template.Template, error) {

	cache := map[string]*template.Template{}

	baseTmpl := basePath + "/base.html"
	pages, err := filepath.Glob(basePath + "/pages/*.html")

	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		fileName := filepath.Base(page)

		// Create a template with given name, HTML from base template and common
		// template functions
		ts, err := template.New(fileName).Funcs(templateFunctions).ParseFiles(baseTmpl)
		if err != nil {
			return nil, err
		}

		// Add all partials to this template set
		ts, err = ts.ParseGlob(basePath + "/partials/*.html")
		if err != nil {
			return nil, err
		}

		// Add the current page to this template set
		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		cache[fileName] = ts
	}

	return cache, nil
}
