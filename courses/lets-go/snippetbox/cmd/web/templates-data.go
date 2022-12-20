package main

import (
	"net/http"
	"time"

	"snippetbox.dev/internal/models"
)

type BreadcrumbLink struct {
	Url      string
	Label    string
	IsActive bool
}

type templateData struct {
	CurrentYear int
	Breadcrumbs []*BreadcrumbLink
	Form        any
	Snippet     *models.Snippet
	Snippets    []*models.Snippet
}

func (app *application) newTemplateData(r *http.Request) *templateData {
	return &templateData{
		CurrentYear: time.Now().Year(),
	}
}
