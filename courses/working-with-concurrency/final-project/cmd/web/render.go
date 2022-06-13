package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"final_project/data"
)

var templatesPath = "./cmd/web/templates"

type TemplateData struct {
	StringMap     map[string]string
	IntMap        map[string]int
	FloatMap      map[string]float64
	Data          map[string]any
	Flash         string
	Warning       string
	Error         string
	Authenticated bool
	Now           time.Time
	User          *data.User
}

func (app *Config) render(
	w http.ResponseWriter,
	r *http.Request,
	t string,
	td *TemplateData,
) {
	templateSlice := []string{
		fmt.Sprintf("%s/%s", templatesPath, t),
		fmt.Sprintf("%s/base.layout.gohtml", templatesPath),
		fmt.Sprintf("%s/header.partial.gohtml", templatesPath),
		fmt.Sprintf("%s/navbar.partial.gohtml", templatesPath),
		fmt.Sprintf("%s/footer.partial.gohtml", templatesPath),
		fmt.Sprintf("%s/alerts.partial.gohtml", templatesPath),
	}

	if td == nil {
		td = &TemplateData{}
	}

	tmpl, err := template.ParseFiles(templateSlice...)

	if err != nil {
		app.ErrorLog.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, app.AddDefaultData(td, r)); err != nil {
		app.ErrorLog.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (app *Config) AddDefaultData(td *TemplateData, r *http.Request) *TemplateData {
	td.Flash = app.Session.PopString(r.Context(), "flash")
	td.Warning = app.Session.PopString(r.Context(), "warning")
	td.Error = app.Session.PopString(r.Context(), "error")
	td.Authenticated = app.IsAuthenticated(r)

	// Extract and deserialize from GOB format the user from Redis
	if td.Authenticated {
		user, ok := app.Session.Get(r.Context(), "user").(data.User)

		if !ok {
			app.ErrorLog.Println("cannot get user from session")
		} else {
			td.User = &user
		}
	}

	td.Now = time.Now()

	return td
}

func (app *Config) IsAuthenticated(r *http.Request) bool {
	return app.Session.Exists(r.Context(), "userID")
}
