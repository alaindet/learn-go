package main

import "net/http"

func (app *application) customNotFound(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Breadcrumbs = []BreadcrumbLink{
		{"/", "Home", true},
		{r.URL.Path, "Page not found", true},
	}
	app.render(w, http.StatusNotFound, "error-not-found.html", data)

	// app.notFound(w)
}
