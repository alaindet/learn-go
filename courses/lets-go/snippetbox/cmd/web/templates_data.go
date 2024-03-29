package main

import (
	"net/http"
	"time"

	"github.com/justinas/nosurf"
	"snippetbox.dev/internal/models"
)

type BreadcrumbLink struct {
	Url      string
	Label    string
	IsActive bool
}

type templateData struct {
	CurrentYear     int
	Breadcrumbs     []BreadcrumbLink
	Form            any
	Flash           string
	IsAuthenticated bool
	Snippet         *models.Snippet
	Snippets        []*models.Snippet
	CSRFToken       string
}

func (app *application) newTemplateData(r *http.Request) *templateData {

	isAuthenticated := app.isAuthenticated(r)
	breadcrumbs := []BreadcrumbLink{
		{"/", "Home", false},
	}

	if isAuthenticated {
		breadcrumbs = append(breadcrumbs, []BreadcrumbLink{
			{"/snippets/new", "Create snippet", false},
		}...)
	}

	return &templateData{
		CurrentYear:     time.Now().Year(),
		Flash:           app.sessionManager.PopString(r.Context(), sessionKeyFlash),
		IsAuthenticated: isAuthenticated,
		Breadcrumbs:     breadcrumbs,
		CSRFToken:       nosurf.Token(r),
	}
}

func (t *templateData) AddBreadcrumbs(breadcrumbs []BreadcrumbLink) {
	t.Breadcrumbs = append(t.Breadcrumbs, breadcrumbs...)
}
