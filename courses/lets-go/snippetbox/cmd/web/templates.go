package main

import (
	"html/template"
	"net/http"
	"path/filepath"
	"time"

	"snippetbox.dev/internal/models"
)

var templateFunctions = template.FuncMap{
	"friendlyDate": friendlyDate,
}

type templateData struct {
	CurrentYear int
	Snippet     *models.Snippet
	Snippets    []*models.Snippet
}

func friendlyDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

func newTemplateCache(basePath string) (map[string]*template.Template, error) {

	cache := map[string]*template.Template{}

	// globalTemplates := []string{
	// 	basePath + "/base.html",
	// 	basePath + "/partials/nav.html",
	// }

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

func (app *application) newTemplateData(r *http.Request) *templateData {
	return &templateData{
		CurrentYear: time.Now().Year(),
	}
}
