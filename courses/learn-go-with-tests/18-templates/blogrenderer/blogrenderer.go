package blogrenderer

import (
	"embed"
	"html/template"
	"io"
)

type Post struct {
	Title       string
	Description string
	Body        string
	Tags        []string
}

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

func Render(w io.Writer, p Post) error {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")

	if err != nil {
		return err
	}

	err = templ.Execute(w, p)

	if err != nil {
		return err
	}

	return nil
}
