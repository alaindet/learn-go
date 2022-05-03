package blogrenderer

import (
	"html/template"
	"io"
)

type Post struct {
	Title       string
	Description string
	Body        string
	Tags        []string
}

// TODO: Move it to file
const (
	postTemplate = `<h1>{{.Title}}</h1><p>{{.Description}}</p>Tags: <ul>{{range .Tags}}<li>{{.}}</li>{{end}}</ul>`
)

func Render(w io.Writer, p Post) error {
	templ, err := template.New("blog").Parse(postTemplate)

	if err != nil {
		return err
	}

	err = templ.Execute(w, p)

	if err != nil {
		return err
	}

	return nil
}
