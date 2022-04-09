package cyoaweb

import (
	"flag"
	"html/template"
	"net/http"
	"os"

	"gophercises-cyoa/cyoa"
)

func LoadStory() cyoa.Story {
	filename := flag.String(
		"file",
		"stories/gopher.json",
		"the JSON file with CYOA story",
	)
	flag.Parse()

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	story, err := cyoa.JsonStory(f)
	if err != nil {
		panic(err)
	}

	return story
}

type handler struct {
	s Story
}

func NewHttpHandler(s cyoa.Story) http.Handler {
	return handler{}
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("").Parse(defaultHtmlTemplate))
	_ = tmpl // TODO
}
